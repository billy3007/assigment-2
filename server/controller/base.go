package controller

import (
	"assigment-2/server/controller/views"

	"github.com/gin-gonic/gin"
)

func WriteJsonResponse(ctx *gin.Context, response *views.Response) {
	ctx.JSON(response.Status, response)
}
