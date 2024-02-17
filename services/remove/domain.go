package remove

import "file-management-api/utils/errs"

type Service interface {
	Execute(request Request) (*Response, errs.AppError)
}

type Request struct {
	UserID   string `json:"userId" validate:"uuid4"`
	FileName string `json:"fileName" validate:"required"`
}

type Response struct{}
