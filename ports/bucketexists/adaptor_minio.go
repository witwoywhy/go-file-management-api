package bucketexists

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

func (a *adaptorMinio) Execute(request Request) (*Response, error) {
	exists, err := a.client.BucketExists(a.ctx, request.BucketName)
	if err != nil {
		return nil, err
	}

	return &Response{
		Exists: exists,
	}, nil
}
