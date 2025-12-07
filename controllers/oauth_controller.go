package controllers

import (
	"diflow/api/database"
	"diflow/api/models"
	"diflow/api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OAuthController struct {
	BaseController
}

func (oc *OAuthController) Login(c *gin.Context) {
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

	accessToken, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Username)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate refresh token")
		return
	}

	expiresAt := time.Now().Add(30 * 24 * time.Hour)
	refreshTokenModel := models.Token{
		UserID:    user.ID,
		TypeID:    int(models.TokenTypeOAuthRefresh),
		Value:     refreshToken,
		ExpiresAt: &expiresAt,
	}
	database.DB.Create(&refreshTokenModel)

	utils.SuccessResponse(c, http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user":          user,
	})
}

func (oc *OAuthController) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Name:     req.Name,
	}

	if err := user.SetPassword(req.Password); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to set password")
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, user)
}

func (oc *OAuthController) Refresh(c *gin.Context) {
	var req RefreshTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var token models.Token
	if err := database.DB.Where("value = ? AND type_id = ?", req.RefreshToken, models.TokenTypeOAuthRefresh).First(&token).Error; err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid refresh token")
		return
	}

	if token.IsExpired() {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Refresh token expired")
		return
	}

	var user models.User
	if err := database.DB.First(&user, token.UserID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User not found")
		return
	}

	accessToken, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}

func (oc *OAuthController) Logout(c *gin.Context) {
	// Invalidate refresh token
	user := oc.GetUser(c)
	if user != nil {
		database.DB.Where("user_id = ? AND type_id = ?", user.ID, models.TokenTypeOAuthRefresh).Delete(&models.Token{})
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func (oc *OAuthController) Access(c *gin.Context) {
	user := oc.GetUser(c)
	if user == nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, user)
}

func (oc *OAuthController) JWT(c *gin.Context) {
	oc.Login(c)
}

