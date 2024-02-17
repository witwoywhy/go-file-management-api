package upload

import (
	"file-management-api/ports/bucketexists"
	"file-management-api/ports/createbucket"
	"file-management-api/ports/listfiles"
	"file-management-api/ports/uploadfile"
	"file-management-api/utils/errs"
	"log"
	"net/http"
)

type service struct {
	bucketExists bucketexists.Port
	listFiles    listfiles.Port
	createBucket createbucket.Port
	uploadFile   uploadfile.Port
}

func New(
	bucketExists bucketexists.Port,
	listFiles listfiles.Port,
	createBucket createbucket.Port,
	uploadFile uploadfile.Port,
) Service {
	return &service{
		bucketExists: bucketExists,
		listFiles:    listFiles,
		createBucket: createBucket,
		uploadFile:   uploadFile,
	}
}

func (s *service) Execute(request Request) (*Response, errs.AppError) {
	if err := request.Validate(); err != nil {
		log.Printf("Failed to validate request: %v", err)
		return nil, errs.New(http.StatusBadRequest, errs.E001, "")
	}

	exists, err := s.bucketExists.Execute(request.ToBucketExistsRequest())
	if err != nil {
		log.Printf("Failed to check exists bucket: %v", err)
		return nil, errs.New(http.StatusInternalServerError, errs.T001, "check existing bucket failed")
	}

	if !exists.Exists {
		err := s.createBucket.Execute(request.ToCreateBucketRequest())
		if err != nil {
			log.Printf("Failed to create bucket: %v", err)
			return nil, errs.New(http.StatusInternalServerError, errs.T001, "create bucket failed")
		}
	}

	listFiles, err := s.listFiles.Execute(request.ToListFilesRequest())
	if err != nil {
		log.Printf("Failed to get list files: %v", err)
		return nil, errs.New(http.StatusInternalServerError, errs.T001, "get list failed")
	}

	if request.IsExistingFileInBucket(listFiles) {
		return nil, errs.New(http.StatusBadRequest, errs.E001, "already exists")
	}

	err = s.uploadFile.Execute(request.BuildUploadFileRequest())
	if err != nil {
		log.Printf("Failed to upload file: %v", err)
		return nil, errs.New(http.StatusInternalServerError, errs.T001, "upload file failed")
	}

	return &Response{}, nil
}
