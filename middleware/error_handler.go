package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			log.Printf("Error: %v", err)

			status := http.StatusInternalServerError
			if err.Meta != nil {
				if s, ok := err.Meta.(int); ok {
					status = s
				}
			}

			c.JSON(status, gin.H{
				"error": err.Error(),
			})
		}
	}
}

