package remove

import (
	"file-management-api/ports/removefile"
	"file-management-api/utils/validate"
)

func (request *Request) Validate() error {
	if err := validate.Validate.Struct(request); err != nil {
		return err
	}

	return nil
}

func (request *Request) ToRemoveFileRequest() removefile.Request {
	return removefile.Request{
		BucketName: request.UserID,
		FileName:   request.FileName,
	}
}
