package controllers

// Zone endpoints

// @Summary      List Zones
// @Description  Получить список зон
// @Tags         zone
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "items: array of Zone"
// @Failure      401  {object}  map[string]string
// @Router       /zone [get]
func _() {}

// @Summary      Get Zone
// @Description  Получить зону по ID
// @Tags         zone
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "Zone ID"
// @Success      200  {object}  models.Zone
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /zone/{id} [get]
func _() {}

// @Summary      Create Zone
// @Description  Создать новую зону
// @Tags         zone
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        zone  body      models.Zone  true  "Zone data"
// @Success      201   {object}  models.Zone
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Router       /zone [post]
func _() {}

// @Summary      Update Zone
// @Description  Обновить зону
// @Tags         zone
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        id    path      int          true  "Zone ID"
// @Param        zone  body      models.Zone  true  "Zone data"
// @Success      200   {object}  models.Zone
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Router       /zone/{id} [put]
// @Router       /zone/{id} [patch]
func _() {}

// Zone Group endpoints

// @Summary      List Zone Groups
// @Description  Получить список групп зон
// @Tags         zone
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /zone-groups [get]
func _() {}

// @Summary      Create Zone Group
// @Description  Создать новую группу зон
// @Tags         zone
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        group  body      map[string]interface{}  true  "Zone Group data"
// @Success      201    {object}  map[string]string
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Router       /zone-groups [post]
func _() {}

// @Summary      Get Zone Group
// @Description  Получить группу зон по ID
// @Tags         zone
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "Zone Group ID"
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /zone-groups/{id} [get]
func _() {}

// @Summary      Update Zone Group
// @Description  Обновить группу зон
// @Tags         zone
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        id     path      int                     true  "Zone Group ID"
// @Param        group  body      map[string]interface{}  true  "Zone Group data"
// @Success      200    {object}  map[string]string
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Failure      404    {object}  map[string]string
// @Router       /zone-groups/{id} [put]
func _() {}
