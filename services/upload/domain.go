package upload

import (
	"file-management-api/utils/errs"
	"mime/multipart"
)

type Service interface {
	Execute(request Request) (*Response, errs.AppError)
}

type Request struct {
	UserID string `json:"userId" validate:"uuid4"`

	File *multipart.FileHeader
}

type Response struct{}
