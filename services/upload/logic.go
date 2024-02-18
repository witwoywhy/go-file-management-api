package upload

import (
	"errors"
	"file-management-api/ports/bucketexists"
	"file-management-api/ports/config"
	"file-management-api/ports/createbucket"
	"file-management-api/ports/listfiles"
	"file-management-api/ports/uploadfile"
	"file-management-api/utils/validate"
	"fmt"
	"strings"
)

func (request *Request) Validate() error {
	if err := validate.Validate.Struct(request); err != nil {
		return err
	}

	if request.File.Size > config.Config.MaxSizeFile {
		return errors.New("over max size")
	}

	split := strings.Split(request.File.Filename, ".")
	if len(split) < 2 {
		return errors.New("not found file extension")
	}

	last := split[len(split)-1]
	if _, ok := config.Config.AllowFileExtensions[last]; !ok {
		return fmt.Errorf("file extension not allow: .%v", last)
	}

	return nil
}

func (request *Request) BuildUploadFileRequest() uploadfile.Request {
	return uploadfile.Request{
		BucketName: request.UserID,
		Header:     request.File.Header["Content-Type"][0],
		File:       request.File,
	}
}

func (request *Request) ToBucketExistsRequest() bucketexists.Request {
	return bucketexists.Request{
		BucketName: request.UserID,
	}
}

func (request *Request) ToCreateBucketRequest() createbucket.Request {
	return createbucket.Request{
		BucketName: request.UserID,
	}
}

func (request *Request) ToListFilesRequest() listfiles.Request {
	return listfiles.Request{
		BucketName: request.UserID,
	}
}

func (request *Request) IsExistingFileInBucket(list *listfiles.Response) bool {
	for _, file := range *list {
		if request.File.Filename == file.FileName {
			return true
		}
	}
	return false
}
