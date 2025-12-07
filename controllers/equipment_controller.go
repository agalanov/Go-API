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

// Equipment sub-resources methods

func (ec *EquipmentController) Types(c *gin.Context) {
	user := ec.GetUser(c)
	query := database.DB.Model(&models.EquipmentType{})
	query = ec.ByCredentials(query, user)

	var types []models.EquipmentType
	if err := query.Find(&types).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, types)
}

func (ec *EquipmentController) Binds(c *gin.Context) {
	user := ec.GetUser(c)
	query := database.DB.Model(&models.EquipmentBind{})
	query = ec.ByCredentials(query, user)

	// Filter by equipment_id if provided
	if equipmentID := c.Query("equipment_id"); equipmentID != "" {
		query = query.Where("equipment_id = ? OR bind_id = ?", equipmentID, equipmentID)
	}

	var binds []models.EquipmentBind
	if err := query.Find(&binds).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, binds)
}

func (ec *EquipmentController) Directions(c *gin.Context) {
	user := ec.GetUser(c)
	query := database.DB.Model(&models.EquipmentDirection{})
	query = ec.ByCredentials(query, user)

	// Filter by building_id if provided
	if buildingID := c.Query("building_id"); buildingID != "" {
		query = query.Where("building_id = ?", buildingID)
	}

	var directions []models.EquipmentDirection
	if err := query.Find(&directions).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, directions)
}

func (ec *EquipmentController) Entities(c *gin.Context) {
	user := ec.GetUser(c)
	query := database.DB.Model(&models.EquipmentEntity{})
	query = ec.ByCredentials(query, user)

	query = query.Preload("Type").Preload("Section")

	// Filter by type_id if provided
	if typeID := c.Query("type_id"); typeID != "" {
		query = query.Where("type_id = ?", typeID)
	}

	// Filter by section_id if provided
	if sectionID := c.Query("section_id"); sectionID != "" {
		query = query.Where("section_id = ?", sectionID)
	}

	var entities []models.EquipmentEntity
	if err := query.Find(&entities).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, entities)
}

func (ec *EquipmentController) Limits(c *gin.Context) {
	user := ec.GetUser(c)
	query := database.DB.Model(&models.EquipmentLimit{})
	query = ec.ByCredentials(query, user)

	// Filter by equipment_id if provided
	if equipmentID := c.Query("equipment_id"); equipmentID != "" {
		query = query.Where("equipment_id = ?", equipmentID)
	}

	var limits []models.EquipmentLimit
	if err := query.Find(&limits).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, limits)
}

func (ec *EquipmentController) Alerts(c *gin.Context) {
	// EquipmentAlert model doesn't exist, return empty array
	utils.CollectionResponse(c, http.StatusOK, []interface{}{})
}

func (ec *EquipmentController) Attributes(c *gin.Context) {
	user := ec.GetUser(c)
	query := database.DB.Model(&models.EquipmentAttribute{})
	query = ec.ByCredentials(query, user)

	// Filter by equipment_id if provided
	if equipmentID := c.Query("equipment_id"); equipmentID != "" {
		query = query.Where("equipment_id = ?", equipmentID)
	}

	var attributes []models.EquipmentAttribute
	if err := query.Find(&attributes).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, attributes)
}

func (ec *EquipmentController) Devices(c *gin.Context) {
	user := ec.GetUser(c)
	query := database.DB.Model(&models.EquipmentDevice{})
	query = ec.ByCredentials(query, user)

	query = query.Preload("Equipment")

	// Filter by equipment_id if provided
	if equipmentID := c.Query("equipment_id"); equipmentID != "" {
		query = query.Where("equipment_id = ?", equipmentID)
	}

	var devices []models.EquipmentDevice
	if err := query.Find(&devices).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, devices)
}

func (ec *EquipmentController) States(c *gin.Context) {
	user := ec.GetUser(c)
	query := database.DB.Model(&models.EquipmentState{})
	query = ec.ByCredentials(query, user)

	// Filter by equipment_id if provided
	if equipmentID := c.Query("equipment_id"); equipmentID != "" {
		query = query.Where("equipment_id = ?", equipmentID)
	}

	var states []models.EquipmentState
	if err := query.Find(&states).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CollectionResponse(c, http.StatusOK, states)
}

