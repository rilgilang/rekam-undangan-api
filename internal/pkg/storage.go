package pkg

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/rilgilang/rekam-undangan-api/config/dotenv"
	"io/ioutil"
)

type Storage interface {
	GetFile(ctx context.Context, path string) (fileBytes []byte, fileName, contentType *string, err error)
}

type storagePkg struct {
	client *minio.Client
	config *dotenv.Config
}

func NewStorage(client *minio.Client, config *dotenv.Config) Storage {
	return &storagePkg{
		client: client,
		config: config,
	}
}

func (c *storagePkg) GetFile(ctx context.Context, path string) (fileBytes []byte, fileName, contentType *string, err error) {

	object, err := c.client.GetObject(ctx, c.config.MinioBucket, path, minio.GetObjectOptions{})

	if err != nil {
		return nil, nil, nil, err
	}

	defer object.Close()

	// Read the object content
	fileBytes, err = ioutil.ReadAll(object)
	if err != nil {
		return nil, nil, nil, err
	}

	// Object stats
	objectStats, err := object.Stat()

	if err != nil {
		return nil, nil, nil, err
	}

	//_, err = object.Read(fileBytes)
	//if err != nil {
	//	//m.logger.Error("minio_get_error:", err.Error())
	//	return nil, err
	//}

	return fileBytes, &objectStats.Key, &objectStats.ContentType, nil
}
