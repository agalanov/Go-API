package controllers

// Floor endpoints

// @Summary      List Floors
// @Description  Получить список этажей
// @Tags         floor
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "items: array of Floor"
// @Failure      401  {object}  map[string]string
// @Router       /floor [get]
func _() {}

// @Summary      Get Floor
// @Description  Получить этаж по ID
// @Tags         floor
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "Floor ID"
// @Success      200  {object}  models.Floor
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /floor/{id} [get]
func _() {}

// @Summary      Create Floor
// @Description  Создать новый этаж
// @Tags         floor
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        floor  body      models.Floor  true  "Floor data"
// @Success      201    {object}  models.Floor
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Router       /floor [post]
func _() {}

// @Summary      Update Floor
// @Description  Обновить этаж
// @Tags         floor
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        id     path      int           true  "Floor ID"
// @Param        floor  body      models.Floor  true  "Floor data"
// @Success      200    {object}  models.Floor
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Failure      404    {object}  map[string]string
// @Router       /floor/{id} [put]
// @Router       /floor/{id} [patch]
func _() {}

// @Summary      Delete Floor
// @Description  Удалить этаж
// @Tags         floor
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "Floor ID"
// @Success      200  {object}  map[string]bool
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /floor/{id} [delete]
func _() {}

// @Summary      List Floor Modes
// @Description  Получить список режимов этажей
// @Tags         floor
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /floor/mode [get]
func _() {}
