package list

import (
	"file-management-api/ports/listfiles"
	"file-management-api/utils/validate"
	"time"
)

func (request *Request) Validate() error {
	if err := validate.Validate.Struct(request); err != nil {
		return err
	}

	return nil
}

func (request *Request) ToListFilesRequest() listfiles.Request {
	return listfiles.Request{
		BucketName: request.UserID,
	}
}

func buildResponse(list *listfiles.Response) *Response {
	response := make(Response, len(*list))

	for i, file := range *list {
		response[i] = List{
			FileName:     file.FileName,
			UploadDateAt: file.UploadDateAt.Local().Format(time.RFC3339),
		}
	}

	return &response
}
