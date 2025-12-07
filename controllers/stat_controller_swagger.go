package controllers

// Stat endpoints

// @Summary      Get Statistics Table
// @Description  Получить статистику в виде таблицы
// @Tags         stat
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]string
// @Router       /stat/table [get]
func _() {}

// @Summary      Get Statistics Graph
// @Description  Получить статистику в виде графика
// @Tags         stat
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]string
// @Router       /stat/graph [get]
func _() {}

// @Summary      Get Statistics Groups
// @Description  Получить группы статистики
// @Tags         stat
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]string
// @Router       /stat/groups [get]
func _() {}

// @Summary      Get Statistics Table by ID
// @Description  Получить статистику в виде таблицы по ID
// @Tags         stat
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "Stat ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]string
// @Router       /stat/{id}/table [get]
func _() {}

// @Summary      Get Statistics Graph by ID
// @Description  Получить статистику в виде графика по ID
// @Tags         stat
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "Stat ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]string
// @Router       /stat/{id}/graph [get]
func _() {}
