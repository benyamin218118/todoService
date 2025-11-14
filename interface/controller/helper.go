package controller

import (
	"net/http"

	"github.com/benyamin218118/todoService/domain"
	"github.com/gin-gonic/gin"
)

func ResponseBadRequest(ctx *gin.Context) {
	Response(ctx, http.StatusBadRequest, map[string]any{"message": "bad request"})
}

func Response(ctx *gin.Context, status int, data any) {
	if strData, isString := data.(string); isString {
		ctx.String(status, strData)
	} else {
		ctx.JSON(status, data)
	}
}

func HandleIfError(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	status := http.StatusInternalServerError
	switch err.(type) {
	case *domain.BadRequestError:
		status = http.StatusBadRequest
	case *domain.ForbiddenError:
		status = http.StatusForbidden
	default:
		status = http.StatusInternalServerError
	}
	ctx.JSON(status, map[string]any{"message": err.Error()})

	return true
}
