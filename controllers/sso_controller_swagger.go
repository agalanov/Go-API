package controllers

// SSO endpoints

// @Summary      SSO Login
// @Description  Авторизация через SSO
// @Tags         sso
// @Accept       json
// @Produce      json
// @Param        credentials  body      LoginRequest  true  "Credentials"
// @Success      200          {object}  LoginResponse
// @Failure      400          {object}  map[string]string
// @Failure      401          {object}  map[string]string
// @Router       /sso/login [post]
// @Router       /sso/login [get]
func _() {}

// @Summary      SSO JWT
// @Description  Получить JWT токен через SSO
// @Tags         sso
// @Accept       json
// @Produce      json
// @Param        credentials  body      LoginRequest  true  "Credentials"
// @Success      200          {object}  map[string]string
// @Failure      400          {object}  map[string]string
// @Failure      401          {object}  map[string]string
// @Router       /sso/jwt [post]
// @Router       /sso/jwt [get]
func _() {}

// @Summary      SSO Query Param
// @Description  Авторизация через SSO с параметрами запроса
// @Tags         sso
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /sso/qp [post]
// @Router       /sso/qp [get]
func _() {}

// @Summary      SSO OAuth
// @Description  Авторизация через OAuth (SSO)
// @Tags         sso
// @Accept       json
// @Produce      json
// @Param        credentials  body      LoginRequest  true  "Credentials"
// @Success      200          {object}  LoginResponse
// @Failure      400          {object}  map[string]string
// @Failure      401          {object}  map[string]string
// @Router       /sso/oauth [post]
// @Router       /sso/oauth [get]
func _() {}
