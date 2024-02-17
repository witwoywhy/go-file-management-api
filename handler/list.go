package handler

import (
	"file-management-api/services/list"
	"net/http"

	"github.com/gin-gonic/gin"
)

type listHandler struct {
	service list.Service
}

func NewListHandler(service list.Service) *listHandler {
	return &listHandler{
		service: service,
	}
}

func (h *listHandler) Handle(ctx *gin.Context) {
	var request list.Request
	request.UserID = ctx.Param("userId")

	response, svcErr := h.service.Execute(request)
	if svcErr != nil {
		ctx.JSON(svcErr.GetHttpStatus(), svcErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
