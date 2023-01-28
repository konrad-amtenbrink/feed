package storage

type (
	MockStorage interface {
		Upload(filename string) error
		Download(filename string) ([]byte, error)
	}

	mockBucket struct {
	}
)

func NewMockStorage() MockStorage {
	return mockBucket{}
}

func (bucket mockBucket) Upload(filepath string) error {
	return nil
}

func (bucket mockBucket) Download(filename string) ([]byte, error) {
	return []byte{}, nil
}
