package httpserv

import (
	"file-management-api/handler"
	"file-management-api/infrastructure"
	"file-management-api/ports/bucketexists"
	"file-management-api/ports/createbucket"
	"file-management-api/ports/listfiles"
	"file-management-api/ports/removefile"
	"file-management-api/ports/uploadfile"
	"file-management-api/services/list"
	"file-management-api/services/remove"
	"file-management-api/services/upload"

	"github.com/gin-gonic/gin"
)

func bindUploadRoute(app *gin.Engine) {
	bucketExists := bucketexists.NewAdaptorMinio(infrastructure.MinioClient)
	listFiles := listfiles.NewAdaptorMinio(infrastructure.MinioClient)
	uploadFile := uploadfile.NewAdaptorMinio(infrastructure.MinioClient)
	createBucket := createbucket.NewAdaptorMinio(infrastructure.MinioClient)

	service := upload.New(bucketExists, listFiles, createBucket, uploadFile)
	handle := handler.NewUploadHandler(service)

	app.PUT("/v1/upload/:userId", handle.Handle)
}

func bindListRoute(app *gin.Engine) {
	listFiles := listfiles.NewAdaptorMinio(infrastructure.MinioClient)

	service := list.New(listFiles)
	handle := handler.NewListHandler(service)

	app.GET("/v1/list/:userId", handle.Handle)
}

func bindRemoveRoute(app *gin.Engine) {
	removeFile := removefile.NewAdaptorMinio(infrastructure.MinioClient)

	service := remove.New(removeFile)
	handle := handler.NewRemoveHandler(service)

	app.DELETE("/v1/file/:userId/:fileName", handle.Handle)
}
