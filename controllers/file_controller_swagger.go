package controllers

// File endpoints

// @Summary      List Files
// @Description  Получить список файлов
// @Tags         file
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "items: array of File"
// @Failure      401  {object}  map[string]string
// @Router       /file [get]
func _() {}

// @Summary      Get File
// @Description  Получить файл по ID
// @Tags         file
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "File ID"
// @Success      200  {object}  models.File
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /file/{id} [get]
func _() {}

// @Summary      Upload File
// @Description  Загрузить новый файл
// @Tags         file
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "File to upload"
// @Success      201   {object}  models.File
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Router       /file [post]
func _() {}

// @Summary      Delete File
// @Description  Удалить файл
// @Tags         file
// @Security     BearerAuth
// @Security     QueryAuth
// @Produce      json
// @Param        id   path      int  true  "File ID"
// @Success      200  {object}  map[string]bool
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /file/{id} [delete]
func _() {}

// @Summary      File Table
// @Description  Обработка файла таблицы
// @Tags         file
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        data  body      map[string]interface{}  true  "Table data"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Router       /file/table [post]
func _() {}

// @Summary      File Pysvc
// @Description  Обработка файла через Python сервис
// @Tags         file
// @Security     BearerAuth
// @Security     QueryAuth
// @Accept       json
// @Produce      json
// @Param        data  body      map[string]interface{}  true  "Pysvc data"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Router       /file/pysvc [post]
func _() {}
