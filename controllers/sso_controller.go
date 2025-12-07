package controllers

import (
	"diflow/api/database"
	"diflow/api/models"
	"diflow/api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SSOController struct {
	BaseController
}

func (sc *SSOController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	if !user.CheckPassword(req.Password) {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Generate JWT token
	accessToken, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// Generate refresh token
	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Username)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate refresh token")
		return
	}

	// Save tokens to database
	expiresAt := time.Now().Add(30 * 24 * time.Hour)
	refreshTokenModel := models.Token{
		UserID:    user.ID,
		TypeID:    int(models.TokenTypeOAuthRefresh),
		Value:     refreshToken,
		ExpiresAt: &expiresAt,
	}
	database.DB.Create(&refreshTokenModel)

	response := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}

	utils.SuccessResponse(c, http.StatusOK, response)
}

func (sc *SSOController) JWT(c *gin.Context) {
	// Similar to login but returns only JWT token
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	if !user.CheckPassword(req.Password) {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"token": token})
}

func (sc *SSOController) QueryParam(c *gin.Context) {
	// Generate query param token
	user := sc.GetUser(c)
	if user == nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// Generate a simple token
	tokenValue := utils.GenerateRandomString(32)
	expiresAt := time.Now().Add(24 * time.Hour)

	token := models.Token{
		UserID:    user.ID,
		TypeID:    int(models.TokenTypeQueryParam),
		Value:     tokenValue,
		ExpiresAt: &expiresAt,
	}

	if err := database.DB.Create(&token).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"access_token": tokenValue})
}

func (sc *SSOController) OAuth(c *gin.Context) {
	// OAuth flow - simplified implementation
	sc.Login(c)
}

