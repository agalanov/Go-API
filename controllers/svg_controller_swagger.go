package controllers

// SVG Equipment endpoints

// @Summary      Get SVG Equipment Binds
// @Description  Получить привязки оборудования для SVG
// @Tags         svg
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]string
// @Router       /svg/equipment/binds [get]
func _() {}

// @Summary      Get SVG Equipment Tags
// @Description  Получить теги оборудования для SVG
// @Tags         svg
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]string
// @Router       /svg/equipment/tags [get]
func _() {}

// @Summary      Get SVG Equipment Alerts
// @Description  Получить оповещения оборудования для SVG
// @Tags         svg
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]string
// @Router       /svg/equipment/alerts [get]
func _() {}

// @Summary      Batch Edit SVG Equipment
// @Description  Массовое редактирование оборудования SVG
// @Tags         svg
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        data  body      map[string]interface{}  true  "Batch edit data"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Router       /svg/equipment/batch-edit [patch]
func _() {}

// @Summary      Batch Edit SVG Equipment Binds
// @Description  Массовое редактирование привязок оборудования SVG
// @Tags         svg
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        data  body      map[string]interface{}  true  "Batch edit data"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Router       /svg/equipment/bind/batch-edit [patch]
func _() {}

// @Summary      Batch Create SVG Equipment Entities
// @Description  Массовое создание сущностей оборудования SVG
// @Tags         svg
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        data  body      map[string]interface{}  true  "Batch create data"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Router       /svg/equipment/entity/batch [post]
func _() {}

// @Summary      Batch Create SVG Equipment Binds
// @Description  Массовое создание привязок оборудования SVG
// @Tags         svg
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        data  body      map[string]interface{}  true  "Batch create data"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Router       /svg/equipment/binds-batch [post]
func _() {}
