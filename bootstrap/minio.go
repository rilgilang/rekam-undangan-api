package bootstrap

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rilgilang/sticker-collection-api/config/dotenv"
)

func NewMinio(config *dotenv.Config) (*minio.Client, error) {
	endpoint := config.MinioEndpoint

	accessKey := config.MinioAccessKey
	secretAccessKey := config.MinioSecretAccessKey

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretAccessKey, ""),
		Secure: true,
	})

	if err != nil {
		return nil, err
	}

	return minioClient, nil
}
