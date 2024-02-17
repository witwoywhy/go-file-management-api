package uploadfile

import (
	"context"

	"github.com/minio/minio-go/v7"
)

type adaptorMinio struct {
	client *minio.Client
	ctx    context.Context
}

func NewAdaptorMinio(client *minio.Client) Port {
	return &adaptorMinio{
		client: client,
		ctx:    context.Background(),
	}
}

func (a *adaptorMinio) Execute(request Request) error {
	b, err := request.File.Open()
	if err != nil {
		return err
	}
	defer b.Close()

	_, err = a.client.PutObject(
		a.ctx,
		request.BucketName,
		request.File.Filename,
		b,
		request.File.Size,
		minio.PutObjectOptions{
			ContentType: request.Header,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
