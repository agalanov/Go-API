package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Items interface{} `json:"items,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func SuccessResponse(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func ErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func CollectionResponse(c *gin.Context, status int, items interface{}) {
	c.JSON(status, gin.H{"items": items})
}

