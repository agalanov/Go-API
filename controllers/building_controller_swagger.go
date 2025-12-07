package controllers

// @Summary      List Buildings
// @Description  Получить список зданий
// @Tags         building
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "items: array of Building"
// @Failure      401  {object}  map[string]string
// @Router       /building [get]

// @Summary      Get Building
// @Description  Получить здание по ID
// @Tags         building
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "Building ID"
// @Success      200  {object}  models.Building
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /building/{id} [get]

// @Summary      Create Building
// @Description  Создать новое здание
// @Tags         building
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        building  body      models.Building  true  "Building data"
// @Success      201      {object}  models.Building
// @Failure      400      {object}  map[string]string
// @Failure      401      {object}  map[string]string
// @Router       /building [post]

// @Summary      Update Building
// @Description  Обновить здание
// @Tags         building
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        id        path      int              true  "Building ID"
// @Param        building  body      models.Building  true  "Building data"
// @Success      200       {object}  models.Building
// @Failure      400       {object}  map[string]string
// @Failure      401       {object}  map[string]string
// @Failure      404       {object}  map[string]string
// @Router       /building/{id} [put]
// @Router       /building/{id} [patch]

// @Summary      Delete Building
// @Description  Удалить здание (soft delete)
// @Tags         building
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "Building ID"
// @Success      200  {object}  map[string]bool
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /building/{id} [delete]
func init() {}

