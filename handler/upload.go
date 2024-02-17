package handler

import (
	"file-management-api/services/upload"
	"net/http"

	"github.com/gin-gonic/gin"
)

type uploadHandler struct {
	service upload.Service
}

func NewUploadHandler(service upload.Service) *uploadHandler {
	return &uploadHandler{
		service: service,
	}
}

func (h *uploadHandler) Handle(ctx *gin.Context) {
	var request upload.Request
	request.UserID = ctx.Param("userId")

	if file, err := ctx.FormFile("file"); err != nil {
		ctx.Status(http.StatusBadGateway)
	} else {
		request.File = file
	}

	response, svcErr := h.service.Execute(request)
	if svcErr != nil {
		ctx.JSON(svcErr.GetHttpStatus(), svcErr)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
