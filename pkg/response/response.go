package response

import (
	"github.com/gin-gonic/gin"
)

type (
	MsgResponse struct {
		Status  string      `json:"status"`
		Message string      `json:"message,omitempty"`
		Data    interface{} `json:"data,omitempty"`
	}
)

func ErrResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, &MsgResponse{
		Status:  "error",
		Message: message,
	})
}

func SuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, &MsgResponse{
		Status: "success",
		Data:   data,
	})
}
