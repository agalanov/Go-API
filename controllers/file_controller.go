package controllers

import (
	"diflow/api/database"
	"diflow/api/models"
	"diflow/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	BaseController
}

func (fc *FileController) Index(c *gin.Context) {
	user := fc.GetUser(c)
	query := database.DB.Model(&models.File{}).Where("deleted_at IS NULL")
	query = fc.ByCredentials(query, user)

	var files []models.File
	if err := query.Find(&files).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, files)
}

func (fc *FileController) View(c *gin.Context) {
	id := c.Param("id")
	user := fc.GetUser(c)

	var file models.File
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = fc.ByCredentials(query, user)

	if err := query.First(&file).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "File not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, file)
}

func (fc *FileController) Create(c *gin.Context) {
	var file models.File
	if err := c.ShouldBindJSON(&file); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Create(&file).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, file)
}

func (fc *FileController) Delete(c *gin.Context) {
	id := c.Param("id")
	user := fc.GetUser(c)

	var file models.File
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = fc.ByCredentials(query, user)

	if err := query.First(&file).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "File not found")
		return
	}

	if err := database.DB.Delete(&file).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"success": true})
}

