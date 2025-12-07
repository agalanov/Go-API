package controllers

// Setting endpoints

// @Summary      List Settings
// @Description  Получить список настроек
// @Tags         setting
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "items: array of Setting"
// @Failure      401  {object}  map[string]string
// @Router       /setting [get]
func _() {}

// @Summary      Get Setting
// @Description  Получить настройку по ID
// @Tags         setting
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "Setting ID"
// @Success      200  {object}  models.Setting
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /setting/{id} [get]
func _() {}

// @Summary      Create Setting
// @Description  Создать новую настройку
// @Tags         setting
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        setting  body      models.Setting  true  "Setting data"
// @Success      201      {object}  models.Setting
// @Failure      400      {object}  map[string]string
// @Failure      401      {object}  map[string]string
// @Router       /setting [post]
func _() {}

// @Summary      Update Setting
// @Description  Обновить настройку
// @Tags         setting
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        id       path      int             true  "Setting ID"
// @Param        setting  body      models.Setting  true  "Setting data"
// @Success      200      {object}  models.Setting
// @Failure      400      {object}  map[string]string
// @Failure      401      {object}  map[string]string
// @Failure      404      {object}  map[string]string
// @Router       /setting/{id} [put]
// @Router       /setting/{id} [patch]
func _() {}
