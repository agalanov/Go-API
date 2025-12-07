package controllers

// @Summary      List Equipment
// @Description  Получить список оборудования
// @Tags         equipment
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "items: array of Equipment"
// @Failure      401  {object}  map[string]string
// @Router       /equipment [get]

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

// @Summary      Create Equipment
// @Description  Создать новое оборудование
// @Tags         equipment
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        equipment  body      models.Equipment  true  "Equipment data"
// @Success      201        {object}  models.Equipment
// @Failure      400        {object}  map[string]string
// @Failure      401        {object}  map[string]string
// @Router       /equipment [post]

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
func init() {}

