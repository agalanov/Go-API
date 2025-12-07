package controllers

// Project Section endpoints

// @Summary      List Project Sections
// @Description  Получить список секций проекта
// @Tags         project-section
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "items: array of ProjectSection"
// @Failure      401  {object}  map[string]string
// @Router       /project-section [get]
func _() {}

// @Summary      Get Project Section
// @Description  Получить секцию проекта по ID
// @Tags         project-section
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "Project Section ID"
// @Success      200  {object}  models.ProjectSection
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /project-section/{id} [get]
func _() {}

// @Summary      Create Project Section
// @Description  Создать новую секцию проекта
// @Tags         project-section
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        section  body      models.ProjectSection  true  "Project Section data"
// @Success      201      {object}  models.ProjectSection
// @Failure      400      {object}  map[string]string
// @Failure      401      {object}  map[string]string
// @Router       /project-section [post]
func _() {}

// @Summary      Update Project Section
// @Description  Обновить секцию проекта
// @Tags         project-section
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        id       path      int                    true  "Project Section ID"
// @Param        section  body      models.ProjectSection  true  "Project Section data"
// @Success      200      {object}  models.ProjectSection
// @Failure      400      {object}  map[string]string
// @Failure      401      {object}  map[string]string
// @Failure      404      {object}  map[string]string
// @Router       /project-section/{id} [put]
// @Router       /project-section/{id} [patch]
func _() {}

// @Summary      Delete Project Section
// @Description  Удалить секцию проекта
// @Tags         project-section
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "Project Section ID"
// @Success      200  {object}  map[string]bool
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /project-section/{id} [delete]
func _() {}
