package controllers

import (
	"diflow/api/database"
	"diflow/api/models"
	"diflow/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OPCController struct {
	BaseController
}

func (oc *OPCController) ServerIndex(c *gin.Context) {
	var servers []models.OPCServer
	if err := database.DB.Find(&servers).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, servers)
}

func (oc *OPCController) ServerView(c *gin.Context) {
	id := c.Param("id")
	var server models.OPCServer
	if err := database.DB.First(&server, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "OPC Server not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, server)
}

func (oc *OPCController) TagIndex(c *gin.Context) {
	var tags []models.OPCTag
	query := database.DB.Preload("Server").Preload("DataType")
	if err := query.Find(&tags).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, tags)
}

func (oc *OPCController) TagView(c *gin.Context) {
	id := c.Param("id")
	var tag models.OPCTag
	if err := database.DB.Preload("Server").Preload("DataType").First(&tag, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "OPC Tag not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, tag)
}

func (oc *OPCController) DatatypeIndex(c *gin.Context) {
	var datatypes []models.OPCDatatype
	if err := database.DB.Find(&datatypes).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, datatypes)
}

func (oc *OPCController) DatatypeView(c *gin.Context) {
	id := c.Param("id")
	var datatype models.OPCDatatype
	if err := database.DB.First(&datatype, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "OPC Datatype not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, datatype)
}

