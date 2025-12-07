package controllers

// Building endpoints

// @Summary      List Buildings
// @Description  Получить список зданий
// @Tags         building
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "items: array of Building"
// @Failure      401  {object}  map[string]string
// @Router       /building [get]
func _() {}

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
func _() {}

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
func _() {}

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
func _() {}

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
func _() {}

// Building Project Section endpoints

// @Summary      List Building Project Sections
// @Description  Получить список секций проекта для здания
// @Tags         building
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "Building ID"
// @Success      200  {object}  map[string]interface{}  "items: array of ProjectSection"
// @Failure      401  {object}  map[string]string
// @Router       /building/{id}/project-section [get]
func _() {}

// @Summary      Create Building Project Section
// @Description  Создать новую секцию проекта для здания
// @Tags         building
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        id       path      int                   true  "Building ID"
// @Param        section  body      models.ProjectSection true  "Project Section data"
// @Success      201      {object}  models.ProjectSection
// @Failure      400      {object}  map[string]string
// @Failure      401      {object}  map[string]string
// @Router       /building/{id}/project-section [post]
func _() {}

// @Summary      Get Building Project Section
// @Description  Получить секцию проекта здания по ID
// @Tags         building
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id          path      int  true  "Building ID"
// @Param        section_id  path      int  true  "Project Section ID"
// @Success      200         {object}  models.ProjectSection
// @Failure      401         {object}  map[string]string
// @Failure      404         {object}  map[string]string
// @Router       /building/{id}/project-section/{section_id} [get]
func _() {}

// @Summary      Update Building Project Section
// @Description  Обновить секцию проекта здания
// @Tags         building
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        id          path      int                   true  "Building ID"
// @Param        section_id  path      int                   true  "Project Section ID"
// @Param        section     body      models.ProjectSection true  "Project Section data"
// @Success      200         {object}  models.ProjectSection
// @Failure      400         {object}  map[string]string
// @Failure      401         {object}  map[string]string
// @Failure      404         {object}  map[string]string
// @Router       /building/{id}/project-section/{section_id} [put]
func _() {}

// @Summary      Delete Building Project Section
// @Description  Удалить секцию проекта здания
// @Tags         building
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id          path      int  true  "Building ID"
// @Param        section_id  path      int  true  "Project Section ID"
// @Success      200         {object}  map[string]bool
// @Failure      401         {object}  map[string]string
// @Failure      404         {object}  map[string]string
// @Router       /building/{id}/project-section/{section_id} [delete]
func _() {}
