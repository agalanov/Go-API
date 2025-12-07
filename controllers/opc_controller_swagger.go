package controllers

// OPC endpoints

// @Summary      List OPC Servers
// @Description  Получить список OPC серверов
// @Tags         opc
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "items: array of OPCServer"
// @Failure      401  {object}  map[string]string
// @Router       /opc/server [get]
func _() {}

// @Summary      Get OPC Server
// @Description  Получить OPC сервер по ID
// @Tags         opc
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "OPC Server ID"
// @Success      200  {object}  models.OPCServer
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /opc/server/{id} [get]
func _() {}

// @Summary      List OPC Tags
// @Description  Получить список OPC тегов
// @Tags         opc
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "items: array of OPCTag"
// @Failure      401  {object}  map[string]string
// @Router       /opc/tag [get]
func _() {}

// @Summary      Get OPC Tag
// @Description  Получить OPC тег по ID
// @Tags         opc
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "OPC Tag ID"
// @Success      200  {object}  models.OPCTag
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /opc/tag/{id} [get]
func _() {}

// @Summary      List OPC Datatypes
// @Description  Получить список типов данных OPC
// @Tags         opc
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "items: array of OPCDatatype"
// @Failure      401  {object}  map[string]string
// @Router       /opc/datatype [get]
func _() {}

// @Summary      Get OPC Datatype
// @Description  Получить тип данных OPC по ID
// @Tags         opc
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "OPC Datatype ID"
// @Success      200  {object}  models.OPCDatatype
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /opc/datatype/{id} [get]
func _() {}
