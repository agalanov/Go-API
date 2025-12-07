package controllers

// @Summary      SSO Login
// @Description  Авторизация через SSO
// @Tags         sso
// @Accept       json
// @Produce      json
// @Param        credentials  body      LoginRequest   true  "Credentials"
// @Success      200          {object}  LoginResponse
// @Failure      400          {object}  map[string]string
// @Failure      401          {object}  map[string]string
// @Router       /sso/login [post]
// @Router       /sso/login [get]

// @Summary      Get JWT Token
// @Description  Получить JWT токен
// @Tags         sso
// @Accept       json
// @Produce      json
// @Param        credentials  body      LoginRequest   true  "Credentials"
// @Success      200          {object}  map[string]string
// @Failure      400          {object}  map[string]string
// @Failure      401          {object}  map[string]string
// @Router       /sso/jwt [post]
// @Router       /sso/jwt [get]

// @Summary      Generate Query Param Token
// @Description  Сгенерировать токен для query параметра
// @Tags         sso
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /sso/qp [post]
// @Router       /sso/qp [get]

// @Summary      OAuth Login
// @Description  Авторизация через OAuth
// @Tags         sso
// @Accept       json
// @Produce      json
// @Param        credentials  body      LoginRequest   true  "Credentials"
// @Success      200          {object}  LoginResponse
// @Failure      400          {object}  map[string]string
// @Failure      401          {object}  map[string]string
// @Router       /sso/oauth [post]
// @Router       /sso/oauth [get]
func init() {}

