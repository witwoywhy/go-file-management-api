package handler

import (
	"file-management-api/services/remove"
	"net/http"

	"github.com/gin-gonic/gin"
)

type removeHandler struct {
	service remove.Service
}

func NewRemoveHandler(service remove.Service) *removeHandler {
	return &removeHandler{
		service: service,
	}
}

func (h *removeHandler) Handle(ctx *gin.Context) {
	var request remove.Request
	request.UserID = ctx.Param("userId")
	request.FileName = ctx.Param("fileName")

	response, svcErr := h.service.Execute(request)
	if svcErr != nil {
		ctx.JSON(svcErr.GetHttpStatus(), svcErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
