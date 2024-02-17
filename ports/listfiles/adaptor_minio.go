package listfiles

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
	ch := a.client.ListObjects(a.ctx, request.BucketName, minio.ListObjectsOptions{})

	var response Response
	for object := range ch {
		response = append(response, List{
			FileName:     object.Key,
			UploadDateAt: object.LastModified,
		})
	}
	return &response, nil
}
