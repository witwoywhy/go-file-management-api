package list

import (
	"file-management-api/ports/listfiles"
	"file-management-api/utils/errs"
	"log"
	"net/http"
)

type service struct {
	listFiles listfiles.Port
}

func New(listFiles listfiles.Port) Service {
	return &service{
		listFiles: listFiles,
	}
}

func (s *service) Execute(request Request) (*Response, errs.AppError) {
	if err := request.Validate(); err != nil {
		log.Printf("Failed to validate request: %v", err)
		return nil, errs.New(http.StatusBadRequest, errs.E001, "")
	}

	list, err := s.listFiles.Execute(request.ToListFilesRequest())
	if err != nil {
		log.Printf("Failed to get list files: %v", err)
		return nil, errs.New(http.StatusInternalServerError, errs.T001, "get list failed")
	}

	return buildResponse(list), nil
}
