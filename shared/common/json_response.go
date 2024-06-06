package common

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/shared/model"
)

// Common Response Send Create Response
func SendCreateResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, &model.SingleResponse{
		Status: model.Status{
			Code:    http.StatusCreated,
			Message: "Created",
		},
		Data: data,
	})
}

// Common Response Send Single Response
func SendSingleResponse(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, &model.SingleResponse{
		Status: model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data: data,
	})
}

// Common Response Send Paged Response
func SendPagedResponse(ctx *gin.Context, data []interface{}, paging model.Paging, message string) {
	ctx.JSON(http.StatusOK, &model.PagedResponse{
		Status: model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data:   data,
		Paging: paging,
	})
}

// Common Response Send Error Response
func SendErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, &model.Status{
		Code:    code,
		Message: message,
	})
}

// Common Response Send Success Response
func SendSuccessResponse(ctx *gin.Context, code int, message any) {
	ctx.AbortWithStatusJSON(code, &model.Status{
		Code:    code,
		Message: message,
	})
}

func ToJSON(payload interface{}) (io.Reader, error) {
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(jsonBytes), nil
}
