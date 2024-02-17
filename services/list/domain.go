package list

import (
	"file-management-api/utils/errs"
)

type Service interface {
	Execute(request Request) (*Response, errs.AppError)
}

type Request struct {
	UserID string `json:"userId" validate:"uuid4"`
}

type Response = []List

type List struct {
	FileName     string `json:"fileName"`
	UploadDateAt string `json:"uploadDateAt"`
}
