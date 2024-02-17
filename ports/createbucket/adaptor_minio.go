package createbucket

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
	return a.client.MakeBucket(a.ctx, request.BucketName, minio.MakeBucketOptions{})
}
