package controllers

// @Summary      OAuth Login
// @Description  Авторизация через OAuth
// @Tags         oauth
// @Accept       json
// @Produce      json
// @Param        credentials  body      LoginRequest  true  "Credentials"
// @Success      200          {object}  map[string]interface{}
// @Failure      400          {object}  map[string]string
// @Failure      401          {object}  map[string]string
// @Router       /oauth/login [post]

// @Summary      OAuth Register
// @Description  Регистрация нового пользователя
// @Tags         oauth
// @Accept       json
// @Produce      json
// @Param        user  body      RegisterRequest  true  "User data"
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Router       /oauth/register [post]

// @Summary      Refresh Token
// @Description  Обновить access token используя refresh token
// @Tags         oauth
// @Accept       json
// @Produce      json
// @Param        request  body      RefreshTokenRequest  true  "Refresh token request"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  map[string]string
// @Failure      401      {object}  map[string]string
// @Router       /oauth/refresh [post]

// @Summary      Logout
// @Description  Выход из системы
// @Tags         oauth
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /oauth/logout [get]

// @Summary      Get Access Info
// @Description  Получить информацию о текущем пользователе
// @Tags         oauth
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]string
// @Router       /oauth/access [get]

// @Summary      Get JWT Token
// @Description  Получить JWT токен
// @Tags         oauth
// @Accept       json
// @Produce      json
// @Param        credentials  body      LoginRequest  true  "Credentials"
// @Success      200          {object}  map[string]string
// @Failure      400          {object}  map[string]string
// @Failure      401          {object}  map[string]string
// @Router       /oauth/jwt [post]
// @Router       /oauth/jwt [get]
func init() {}

