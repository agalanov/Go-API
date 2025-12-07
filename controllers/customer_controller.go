package controllers

import (
	"diflow/api/database"
	"diflow/api/models"
	"diflow/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	BaseController
}

func (cc *CustomerController) Index(c *gin.Context) {
	user := cc.GetUser(c)
	query := database.DB.Model(&models.Customer{}).Where("deleted_at IS NULL")
	query = cc.ByCredentials(query, user)

	var customers []models.Customer
	if err := query.Find(&customers).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, customers)
}

func (cc *CustomerController) View(c *gin.Context) {
	id := c.Param("id")
	user := cc.GetUser(c)

	var customer models.Customer
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = cc.ByCredentials(query, user)

	if err := query.Preload("Buildings").First(&customer).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Customer not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, customer)
}

func (cc *CustomerController) Create(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Create(&customer).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, customer)
}

func (cc *CustomerController) Update(c *gin.Context) {
	id := c.Param("id")
	user := cc.GetUser(c)

	var customer models.Customer
	query := database.DB.Where("id = ? AND deleted_at IS NULL", id)
	query = cc.ByCredentials(query, user)

	if err := query.First(&customer).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Customer not found")
		return
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Save(&customer).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, customer)
}

