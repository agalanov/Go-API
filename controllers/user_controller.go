package controllers

import (
	"diflow/api/database"
	"diflow/api/models"
	"diflow/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (uc *UserController) View(c *gin.Context) {
	username := c.Param("id")
	user := uc.GetUser(c)

	var result models.User
	query := database.DB.Where("username = ?", username)
	query = uc.ByCredentials(query, user)

	if err := query.First(&result).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, result)
}

