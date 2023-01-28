package storage

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/konrad-amtenbrink/feed/config"
)

type (
	Storage interface {
		Upload(fileId string, reader io.Reader) error
		Download(filename string) ([]byte, error)
	}

	CloseFunc func() error

	bucket struct {
		session    *session.Session
		bucketName string
		region     string
	}
)

func NewStorage(ctx context.Context, cfg config.AWSConfig) (Storage, CloseFunc, error) {
	session, err := connectAws(cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("initiliazing storage: %v", err)
	}

	closeFunc := func() error {
		// TODO handle session close
		return nil
	}
	return bucket{session, cfg.BucketName, cfg.Region}, closeFunc, nil
}

func connectAws(cfg config.AWSConfig) (*session.Session, error) {
	session, err := session.NewSession(
		&aws.Config{
			Region: aws.String(cfg.Region),
			Credentials: credentials.NewStaticCredentials(
				cfg.AccessKey,
				cfg.SecretAccessKey,
				"",
			),
		})
	if err != nil {
		return nil, fmt.Errorf("connecting to aws: %v", err)
	}
	return session, nil
}

func (bucket bucket) Upload(fileId string, reader io.Reader) error {
	session := bucket.session
	uploader := s3manager.NewUploader(session)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket.bucketName),
		ACL:    aws.String("public-read"),
		Key:    aws.String(fileId),
		Body:   reader,
	})
	if err != nil {
		return fmt.Errorf("uploading to aws: %v", err)
	}

	log.Default().Println("File uploaded to", bucket.bucketName, "/", fileId)

	return nil
}

func (bucket bucket) Download(filename string) ([]byte, error) {
	session := bucket.session
	downloader := s3manager.NewDownloader(session)

	buffer := &aws.WriteAtBuffer{}
	_, err := downloader.Download(buffer,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket.bucketName),
			Key:    aws.String(filename + ".md"),
		})

	if err != nil {
		return nil, fmt.Errorf("downloading from aws: %v", err)
	}

	log.Default().Println("File downloaded: ", filename)

	return buffer.Bytes(), nil
}
