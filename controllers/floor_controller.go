package controllers

import (
	"diflow/api/database"
	"diflow/api/models"
	"diflow/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FloorController struct {
	BaseController
}

func (fc *FloorController) Index(c *gin.Context) {
	user := fc.GetUser(c)
	query := database.DB.Model(&models.Floor{}).Where("deleted_at IS NULL")
	query = fc.ByCredentials(query, user)

	var floors []models.Floor
	if err := query.Preload("Building").Find(&floors).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, floors)
}

func (fc *FloorController) View(c *gin.Context) {
	id := c.Param("id")
	user := fc.GetUser(c)

	var floor models.Floor
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = fc.ByCredentials(query, user)

	if err := query.Preload("Building").Preload("Zones").Preload("Equipments").First(&floor).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Floor not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, floor)
}

func (fc *FloorController) Create(c *gin.Context) {
	var floor models.Floor
	if err := c.ShouldBindJSON(&floor); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Create(&floor).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, floor)
}

func (fc *FloorController) Update(c *gin.Context) {
	id := c.Param("id")
	user := fc.GetUser(c)

	var floor models.Floor
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = fc.ByCredentials(query, user)

	if err := query.First(&floor).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Floor not found")
		return
	}

	if err := c.ShouldBindJSON(&floor); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Save(&floor).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, floor)
}

func (fc *FloorController) Delete(c *gin.Context) {
	id := c.Param("id")
	user := fc.GetUser(c)

	var floor models.Floor
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = fc.ByCredentials(query, user)

	if err := query.First(&floor).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Floor not found")
		return
	}

	if err := database.DB.Delete(&floor).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"success": true})
}

