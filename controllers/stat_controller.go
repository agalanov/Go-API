package controllers

import (
	"diflow/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatController struct {
	BaseController
}

func (sc *StatController) Table(c *gin.Context) {
	// Placeholder for stat table endpoint
	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "Stat table endpoint"})
}

func (sc *StatController) Graph(c *gin.Context) {
	// Placeholder for stat graph endpoint
	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "Stat graph endpoint"})
}

func (sc *StatController) Groups(c *gin.Context) {
	// Placeholder for stat groups endpoint
	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "Stat groups endpoint"})
}

