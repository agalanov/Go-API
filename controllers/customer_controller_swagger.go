package controllers

// Customer endpoints

// @Summary      List Customers
// @Description  Получить список клиентов
// @Tags         customer
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "items: array of Customer"
// @Failure      401  {object}  map[string]string
// @Router       /customer [get]
func _() {}

// @Summary      Get Customer
// @Description  Получить клиента по ID
// @Tags         customer
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "Customer ID"
// @Success      200  {object}  models.Customer
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /customer/{id} [get]
func _() {}

// @Summary      Create Customer
// @Description  Создать нового клиента
// @Tags         customer
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        customer  body      models.Customer  true  "Customer data"
// @Success      201       {object}  models.Customer
// @Failure      400       {object}  map[string]string
// @Failure      401       {object}  map[string]string
// @Router       /customer [post]
func _() {}

// @Summary      Update Customer
// @Description  Обновить клиента
// @Tags         customer
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        id        path      int              true  "Customer ID"
// @Param        customer  body      models.Customer  true  "Customer data"
// @Success      200       {object}  models.Customer
// @Failure      400       {object}  map[string]string
// @Failure      401       {object}  map[string]string
// @Failure      404       {object}  map[string]string
// @Router       /customer/{id} [put]
// @Router       /customer/{id} [patch]
func _() {}
