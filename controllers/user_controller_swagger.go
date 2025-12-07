package controllers

// User endpoints

// @Summary      Get User
// @Description  Получить пользователя по username
// @Tags         user
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      string  true  "Username"
// @Success      200  {object}  models.User
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /user/{id} [get]
func _() {}
