package controllers

import (
	"diflow/api/database"
	"diflow/api/models"
	"diflow/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SettingController struct {
	BaseController
}

func (sc *SettingController) Index(c *gin.Context) {
	user := sc.GetUser(c)
	query := database.DB.Model(&models.Setting{}).Where("deleted_at IS NULL")
	query = sc.ByCredentials(query, user)

	var settings []models.Setting
	if err := query.Find(&settings).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, settings)
}

func (sc *SettingController) View(c *gin.Context) {
	id := c.Param("id")
	user := sc.GetUser(c)

	var setting models.Setting
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = sc.ByCredentials(query, user)

	if err := query.First(&setting).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Setting not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, setting)
}

func (sc *SettingController) Create(c *gin.Context) {
	var setting models.Setting
	if err := c.ShouldBindJSON(&setting); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Create(&setting).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, setting)
}

func (sc *SettingController) Update(c *gin.Context) {
	id := c.Param("id")
	user := sc.GetUser(c)

	var setting models.Setting
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = sc.ByCredentials(query, user)

	if err := query.First(&setting).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Setting not found")
		return
	}

	if err := c.ShouldBindJSON(&setting); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Save(&setting).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, setting)
}

