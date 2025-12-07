package controllers

import (
	"diflow/api/database"
	"diflow/api/models"
	"diflow/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BuildingController struct {
	BaseController
}

func (bc *BuildingController) Index(c *gin.Context) {
	user := bc.GetUser(c)
	query := database.DB.Model(&models.Building{}).Where("deleted_at IS NULL")
	query = bc.ByCredentials(query, user)

	var buildings []models.Building
	if err := query.Find(&buildings).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, buildings)
}

func (bc *BuildingController) View(c *gin.Context) {
	id := c.Param("id")
	user := bc.GetUser(c)

	var building models.Building
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = bc.ByCredentials(query, user)

	if err := query.Preload("Customer").Preload("Equipments").First(&building).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Building not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, building)
}

func (bc *BuildingController) Create(c *gin.Context) {
	var building models.Building
	if err := c.ShouldBindJSON(&building); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Create(&building).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, building)
}

func (bc *BuildingController) Update(c *gin.Context) {
	id := c.Param("id")
	user := bc.GetUser(c)

	var building models.Building
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = bc.ByCredentials(query, user)

	if err := query.First(&building).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Building not found")
		return
	}

	if err := c.ShouldBindJSON(&building); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Save(&building).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, building)
}

func (bc *BuildingController) Delete(c *gin.Context) {
	id := c.Param("id")
	user := bc.GetUser(c)

	var building models.Building
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = bc.ByCredentials(query, user)

	if err := query.First(&building).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Building not found")
		return
	}

	// Soft delete
	if err := database.DB.Delete(&building).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"success": true})
}

