package uploadfile

import "mime/multipart"

type Port interface {
	Execute(request Request) error
}

type Request struct {
	BucketName string
	Header     string
	File       *multipart.FileHeader
}
