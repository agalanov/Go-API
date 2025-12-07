package middleware

import (
	"net/http"
	"strings"

	"diflow/api/database"
	"diflow/api/models"
	"diflow/api/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Try to get token from Authorization header
		authHeader := c.GetHeader("Authorization")
		var tokenString string

		if authHeader != "" {
			// Bearer token
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				tokenString = parts[1]
			}
		}

		// Try to get token from query parameter
		if tokenString == "" {
			tokenString = c.Query("access_token")
		}

		// Try JWT token
		if tokenString != "" {
			// Validate JWT token
			jwtToken, err := utils.ValidateJWT(tokenString)
			if err == nil && jwtToken.Valid {
				if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok {
					if userIDFloat, ok := claims["user_id"].(float64); ok {
						userID := uint(userIDFloat)

						var user models.User
						if err := database.DB.First(&user, userID).Error; err == nil {
							c.Set("user", &user)
							c.Set("user_id", userID)
							c.Next()
							return
						}
					}
				}
			}

			// Try database token
			var dbToken models.Token
			if err := database.DB.Where("value = ? AND expires_at > NOW()", tokenString).First(&dbToken).Error; err == nil {
				var user models.User
				if err := database.DB.First(&user, dbToken.UserID).Error; err == nil {
					c.Set("user", &user)
					c.Set("user_id", user.ID)
					c.Next()
					return
				}
			}
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
	}
}

