package integrationtest

import (
	"context"
	"file-management-api/infrastructure"
	"file-management-api/ports/config"
	"net/http"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	testcontainersgo "github.com/testcontainers/testcontainers-go"
	testcontainers "github.com/testcontainers/testcontainers-go/modules/minio"
)

const (
	baseUrl  = "http://localhost:8080"
	userID   = "0f5384b4-cfe2-4e3e-95d4-3ba6b6d639ec"
	fileName = "cat.jpg"
)

var (
	client = &http.Client{}
)

func must(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func runContainer(t *testing.T, ctx context.Context) *testcontainers.MinioContainer {
	c, err := testcontainers.RunContainer(
		ctx,
		testcontainersgo.WithImage("quay.io/minio/minio"),
	)
	must(t, err)

	return c
}

func initApp(t *testing.T, container *testcontainers.MinioContainer, ctx context.Context) {
	endPoint, err := container.Endpoint(ctx, "")
	must(t, err)

	client, err := minio.New(endPoint, &minio.Options{
		Creds: credentials.NewStaticV4(container.Username, container.Password, ""),
	})
	must(t, err)

	infrastructure.MinioClient = client

	viper.Set("app.port", "8080")

	config.Config.AllowFileExtensions = map[string]bool{
		"jpg": true,
	}
	config.Config.MaxSizeFile = 5000000
}
