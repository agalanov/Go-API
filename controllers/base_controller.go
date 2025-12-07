package controllers

import (
	"diflow/api/database"
	"diflow/api/models"
	"diflow/api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseController struct{}

func (bc *BaseController) GetUser(c *gin.Context) *models.User {
	user, exists := c.Get("user")
	if !exists {
		return nil
	}
	return user.(*models.User)
}

func (bc *BaseController) GetUserID(c *gin.Context) uint {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0
	}
	return userID.(uint)
}

func (bc *BaseController) ByCredentials(query *gorm.DB, user *models.User) *gorm.DB {
	if user != nil && !user.IsSuperuser() {
		return query.Where("id = ?", user.ID)
	}
	return query
}

func (bc *BaseController) Paginate(c *gin.Context, query *gorm.DB, model interface{}) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}

	var total int64
	query.Model(model).Count(&total)

	offset := (page - 1) * pageSize
	query.Offset(offset).Limit(pageSize)

	var items []interface{}
	query.Find(&items)

	utils.CollectionResponse(c, http.StatusOK, items)
}

func (bc *BaseController) Index(c *gin.Context, query *gorm.DB, model interface{}) {
	user := bc.GetUser(c)
	query = bc.ByCredentials(query, user)

	// Apply filters
	if filter := c.Query("filter"); filter != "" {
		// Simple filter implementation - can be extended
		// For now, just apply basic filtering
	}

	bc.Paginate(c, query, model)
}

func (bc *BaseController) View(c *gin.Context, query *gorm.DB, model interface{}, id string) {
	user := bc.GetUser(c)
	query = bc.ByCredentials(query, user)

	var result interface{}
	if err := query.First(&result, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Not found")
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, result)
}

func (bc *BaseController) Create(c *gin.Context, model interface{}) {
	if err := c.ShouldBindJSON(model); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Create(model).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, model)
}

func (bc *BaseController) Update(c *gin.Context, query *gorm.DB, model interface{}, id string) {
	user := bc.GetUser(c)
	query = bc.ByCredentials(query, user)

	var existing interface{}
	if err := query.First(&existing, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Not found")
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.ShouldBindJSON(model); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Model(&existing).Updates(model).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, existing)
}

func (bc *BaseController) Delete(c *gin.Context, query *gorm.DB, model interface{}, id string) {
	user := bc.GetUser(c)
	query = bc.ByCredentials(query, user)

	var existing interface{}
	if err := query.First(&existing, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Not found")
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Soft delete
	if err := database.DB.Delete(&existing).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{"success": true})
}

