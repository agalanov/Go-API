package controllers

// Equipment endpoints

// @Summary      List Equipment
// @Description  Получить список оборудования
// @Tags         equipment
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "items: array of Equipment"
// @Failure      401  {object}  map[string]string
// @Router       /equipment [get]
func _() {}

// @Summary      Get Equipment
// @Description  Получить оборудование по ID
// @Tags         equipment
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "Equipment ID"
// @Success      200  {object}  models.Equipment
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /equipment/{id} [get]
func _() {}

// @Summary      Create Equipment
// @Description  Создать новое оборудование
// @Tags         equipment
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        equipment  body      models.Equipment  true  "Equipment data"
// @Success      201       {object}  models.Equipment
// @Failure      400       {object}  map[string]string
// @Failure      401       {object}  map[string]string
// @Router       /equipment [post]
func _() {}

// @Summary      Update Equipment
// @Description  Обновить оборудование
// @Tags         equipment
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        id         path      int               true  "Equipment ID"
// @Param        equipment  body      models.Equipment  true  "Equipment data"
// @Success      200        {object}  models.Equipment
// @Failure      400        {object}  map[string]string
// @Failure      401        {object}  map[string]string
// @Failure      404        {object}  map[string]string
// @Router       /equipment/{id} [put]
// @Router       /equipment/{id} [patch]
func _() {}

// @Summary      Delete Equipment
// @Description  Удалить оборудование
// @Tags         equipment
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "Equipment ID"
// @Success      200  {object}  map[string]bool
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /equipment/{id} [delete]
func _() {}

// Equipment sub-resources

// @Summary      List Equipment Types
// @Description  Получить список типов оборудования
// @Tags         equipment
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /equipment/type [get]
func _() {}

// @Summary      List Equipment Binds
// @Description  Получить список привязок оборудования
// @Tags         equipment
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /equipment/bind [get]
func _() {}

// @Summary      List Equipment Directions
// @Description  Получить список направлений оборудования
// @Tags         equipment
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /equipment/direction [get]
func _() {}

// @Summary      List Equipment Entities
// @Description  Получить список сущностей оборудования
// @Tags         equipment
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /equipment/entity [get]
func _() {}

// @Summary      List Equipment Limits
// @Description  Получить список лимитов оборудования
// @Tags         equipment
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /equipment/limit [get]
func _() {}

// @Summary      List Equipment Alerts
// @Description  Получить список оповещений оборудования
// @Tags         equipment
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /equipment/alert [get]
func _() {}

// @Summary      List Equipment Attributes
// @Description  Получить список атрибутов оборудования
// @Tags         equipment
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /equipment/attribute [get]
func _() {}

// @Summary      List Equipment Devices
// @Description  Получить список устройств оборудования
// @Tags         equipment
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /equipment/device [get]
func _() {}

// @Summary      List Equipment States
// @Description  Получить список состояний оборудования
// @Tags         equipment
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /equipment/state [get]
func _() {}
