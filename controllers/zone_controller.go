package controllers

import (
	"diflow/api/database"
	"diflow/api/models"
	"diflow/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ZoneController struct {
	BaseController
}

func (zc *ZoneController) Index(c *gin.Context) {
	user := zc.GetUser(c)
	query := database.DB.Model(&models.Zone{}).Where("deleted_at IS NULL")
	query = zc.ByCredentials(query, user)

	var zones []models.Zone
	if err := query.Preload("Floor").Preload("Building").Find(&zones).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, zones)
}

func (zc *ZoneController) View(c *gin.Context) {
	id := c.Param("id")
	user := zc.GetUser(c)

	var zone models.Zone
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = zc.ByCredentials(query, user)

	if err := query.Preload("Floor").Preload("Building").Preload("Equipments").First(&zone).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Zone not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, zone)
}

func (zc *ZoneController) Create(c *gin.Context) {
	var zone models.Zone
	if err := c.ShouldBindJSON(&zone); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Create(&zone).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, zone)
}

func (zc *ZoneController) Update(c *gin.Context) {
	id := c.Param("id")
	user := zc.GetUser(c)

	var zone models.Zone
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = zc.ByCredentials(query, user)

	if err := query.First(&zone).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Zone not found")
		return
	}

	if err := c.ShouldBindJSON(&zone); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Save(&zone).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, zone)
}

