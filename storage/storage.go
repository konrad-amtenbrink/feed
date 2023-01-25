package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/konrad-amtenbrink/feed/config"
)

type (
	Storage interface {
		Upload(filename string) error
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

func (bucket bucket) Upload(filename string) error {
	session := bucket.session
	uploader := s3manager.NewUploader(session)

	content, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("uploading to aws: %v", err)
	}

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket.bucketName),
		ACL:    aws.String("public-read"),
		Key:    aws.String(filename),
		Body:   content,
	})

	if err != nil {
		return fmt.Errorf("uploading to aws: %v", err)
	}

	filepath := "https://" + bucket.bucketName + "." + "s3-" + bucket.region + ".amazonaws.com/" + filename
	log.Default().Println("File uploaded to", filepath)

	return nil
}
