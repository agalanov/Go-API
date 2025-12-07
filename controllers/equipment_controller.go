package controllers

import (
	"diflow/api/database"
	"diflow/api/models"
	"diflow/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EquipmentController struct {
	BaseController
}

func (ec *EquipmentController) Index(c *gin.Context) {
	user := ec.GetUser(c)
	query := database.DB.Model(&models.Equipment{}).Where("deleted_at IS NULL")
	query = ec.ByCredentials(query, user)

	query = query.Preload("Entity").Preload("Entity.Type").
		Preload("Zone").Preload("Building").Preload("Floor").
		Preload("Devices").Preload("Limits")

	var equipments []models.Equipment
	if err := query.Find(&equipments).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, equipments)
}

func (ec *EquipmentController) View(c *gin.Context) {
	id := c.Param("id")
	user := ec.GetUser(c)

	var equipment models.Equipment
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = ec.ByCredentials(query, user)

	if err := query.Preload("Entity").Preload("Entity.Type").
		Preload("Zone").Preload("Building").Preload("Floor").
		Preload("Devices").Preload("Limits").Preload("Binds").
		First(&equipment).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Equipment not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, equipment)
}

func (ec *EquipmentController) Create(c *gin.Context) {
	var equipment models.Equipment
	if err := c.ShouldBindJSON(&equipment); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Create(&equipment).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, equipment)
}

func (ec *EquipmentController) Update(c *gin.Context) {
	id := c.Param("id")
	user := ec.GetUser(c)

	var equipment models.Equipment
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = ec.ByCredentials(query, user)

	if err := query.First(&equipment).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Equipment not found")
		return
	}

	if err := c.ShouldBindJSON(&equipment); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Save(&equipment).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, equipment)
}

func (ec *EquipmentController) Delete(c *gin.Context) {
	id := c.Param("id")
	user := ec.GetUser(c)

	var equipment models.Equipment
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = ec.ByCredentials(query, user)

	if err := query.First(&equipment).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Equipment not found")
		return
	}

	// Delete related binds
	database.DB.Where("equipment_id = ? OR bind_id = ?", id, id).Delete(&models.EquipmentBind{})

	// Delete equipment
	if err := database.DB.Delete(&equipment).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"success": true})
}

