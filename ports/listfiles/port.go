package listfiles

import (
	"time"
)

type Port interface {
	Execute(request Request) (*Response, error)
}

type Request struct {
	BucketName string
}

type Response = []List

type List struct {
	FileName     string
	UploadDateAt time.Time
}
