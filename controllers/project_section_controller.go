package controllers

import (
	"diflow/api/database"
	"diflow/api/models"
	"diflow/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjectSectionController struct {
	BaseController
}

func (psc *ProjectSectionController) Index(c *gin.Context) {
	user := psc.GetUser(c)
	query := database.DB.Model(&models.ProjectSection{}).Where("deleted_at IS NULL")
	query = psc.ByCredentials(query, user)

	var sections []models.ProjectSection
	if err := query.Preload("Buildings").Find(&sections).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, sections)
}

func (psc *ProjectSectionController) View(c *gin.Context) {
	id := c.Param("id")
	user := psc.GetUser(c)

	var section models.ProjectSection
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = psc.ByCredentials(query, user)

	if err := query.Preload("Buildings").First(&section).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Project section not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, section)
}

func (psc *ProjectSectionController) Create(c *gin.Context) {
	var section models.ProjectSection
	if err := c.ShouldBindJSON(&section); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Create(&section).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, section)
}

func (psc *ProjectSectionController) Update(c *gin.Context) {
	id := c.Param("id")
	user := psc.GetUser(c)

	var section models.ProjectSection
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = psc.ByCredentials(query, user)

	if err := query.First(&section).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Project section not found")
		return
	}

	if err := c.ShouldBindJSON(&section); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Save(&section).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, section)
}

func (psc *ProjectSectionController) Delete(c *gin.Context) {
	id := c.Param("id")
	user := psc.GetUser(c)

	var section models.ProjectSection
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = psc.ByCredentials(query, user)

	if err := query.First(&section).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Project section not found")
		return
	}

	if err := database.DB.Delete(&section).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"success": true})
}

