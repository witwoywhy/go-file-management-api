package remove

import (
	"file-management-api/ports/removefile"
	"file-management-api/utils/errs"
	"log"
	"net/http"
)

type service struct {
	removeFile removefile.Port
}

func New(removeFile removefile.Port) Service {
	return &service{
		removeFile: removeFile,
	}
}

func (s *service) Execute(request Request) (*Response, errs.AppError) {
	if err := request.Validate(); err != nil {
		log.Printf("Failed to validate request: %v", err)
		return nil, errs.New(http.StatusBadRequest, errs.E001, "")
	}

	err := s.removeFile.Execute(request.ToRemoveFileRequest())
	if err != nil {
		log.Printf("Failed to remove file: %v", err)
		return nil, errs.New(http.StatusInternalServerError, errs.T001, "remove failed")
	}

	return &Response{}, nil
}
