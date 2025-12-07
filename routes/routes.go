package routes

import (
	"diflow/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterSSORoutes(r *gin.RouterGroup) {
	ssoController := &controllers.SSOController{}
	sso := r.Group("/sso")
	{
		sso.POST("/login", ssoController.Login)
		sso.GET("/login", ssoController.Login)
		sso.POST("/jwt", ssoController.JWT)
		sso.GET("/jwt", ssoController.JWT)
		sso.POST("/qp", ssoController.QueryParam)
		sso.GET("/qp", ssoController.QueryParam)
		sso.POST("/oauth", ssoController.OAuth)
		sso.GET("/oauth", ssoController.OAuth)
	}
}

func RegisterOAuthRoutes(r *gin.RouterGroup) {
	oauthController := &controllers.OAuthController{}
	oauth := r.Group("/oauth")
	{
		oauth.POST("/login", oauthController.Login)
		oauth.GET("/logout", oauthController.Logout)
		oauth.GET("/access", oauthController.Access)
		oauth.POST("/refresh", oauthController.Refresh)
		oauth.POST("/register", oauthController.Register)
		oauth.POST("/jwt", oauthController.JWT)
		oauth.GET("/jwt", oauthController.JWT)
	}
}

func RegisterUserRoutes(r *gin.RouterGroup) {
	userController := &controllers.UserController{}
	users := r.Group("/user")
	{
		users.GET("/:id", userController.View)
	}
}

func RegisterBuildingRoutes(r *gin.RouterGroup) {
	buildingController := &controllers.BuildingController{}
	buildings := r.Group("/building")
	{
		buildings.GET("", buildingController.Index)
		buildings.POST("", buildingController.Create)
		buildings.GET("/:id", buildingController.View)
		buildings.PUT("/:id", buildingController.Update)
		buildings.PATCH("/:id", buildingController.Update)
		buildings.DELETE("/:id", buildingController.Delete)
	}

	// Building project sections
	projectSectionController := &controllers.ProjectSectionController{}
	buildings.GET("/:id/project-section", projectSectionController.Index)
	buildings.POST("/:id/project-section", projectSectionController.Create)
	buildings.GET("/:id/project-section/:section_id", projectSectionController.View)
	buildings.PUT("/:id/project-section/:section_id", projectSectionController.Update)
	buildings.DELETE("/:id/project-section/:section_id", projectSectionController.Delete)
}

func RegisterEquipmentRoutes(r *gin.RouterGroup) {
	equipmentController := &controllers.EquipmentController{}
	equipments := r.Group("/equipment")
	{
		equipments.GET("", equipmentController.Index)
		equipments.POST("", equipmentController.Create)
		equipments.GET("/:id", equipmentController.View)
		equipments.PUT("/:id", equipmentController.Update)
		equipments.PATCH("/:id", equipmentController.Update)
		equipments.DELETE("/:id", equipmentController.Delete)
	}

	// Equipment sub-resources
	equipments.GET("/type", equipmentController.Types)
	equipments.GET("/bind", equipmentController.Binds)
	equipments.GET("/direction", equipmentController.Directions)
	equipments.GET("/entity", equipmentController.Entities)
	equipments.GET("/limit", equipmentController.Limits)
	equipments.GET("/alert", equipmentController.Alerts)
	equipments.GET("/attribute", equipmentController.Attributes)
	equipments.GET("/device", equipmentController.Devices)
	equipments.GET("/state", equipmentController.States)
}

func RegisterFloorRoutes(r *gin.RouterGroup) {
	floorController := &controllers.FloorController{}
	floors := r.Group("/floor")
	{
		floors.GET("", floorController.Index)
		floors.POST("", floorController.Create)
		floors.GET("/:id", floorController.View)
		floors.PUT("/:id", floorController.Update)
		floors.PATCH("/:id", floorController.Update)
		floors.DELETE("/:id", floorController.Delete)
	}

	floors.GET("/mode", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Floor modes endpoint"})
	})
}

func RegisterCustomerRoutes(r *gin.RouterGroup) {
	customerController := &controllers.CustomerController{}
	customers := r.Group("/customer")
	{
		customers.GET("", customerController.Index)
		customers.POST("", customerController.Create)
		customers.GET("/:id", customerController.View)
		customers.PUT("/:id", customerController.Update)
		customers.PATCH("/:id", customerController.Update)
	}
}

func RegisterZoneRoutes(r *gin.RouterGroup) {
	zoneController := &controllers.ZoneController{}
	zones := r.Group("/zone")
	{
		zones.GET("", zoneController.Index)
		zones.POST("", zoneController.Create)
		zones.GET("/:id", zoneController.View)
		zones.PUT("/:id", zoneController.Update)
		zones.PATCH("/:id", zoneController.Update)
	}

	zoneGroups := r.Group("/zone-groups")
	{
		zoneGroups.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Zone groups endpoint"})
		})
		zoneGroups.POST("", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Zone groups create endpoint"})
		})
		zoneGroups.GET("/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Zone groups view endpoint"})
		})
		zoneGroups.PUT("/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Zone groups update endpoint"})
		})
	}
}

func RegisterProjectSectionRoutes(r *gin.RouterGroup) {
	projectSectionController := &controllers.ProjectSectionController{}
	sections := r.Group("/project-section")
	{
		sections.GET("", projectSectionController.Index)
		sections.POST("", projectSectionController.Create)
		sections.GET("/:id", projectSectionController.View)
		sections.PUT("/:id", projectSectionController.Update)
		sections.PATCH("/:id", projectSectionController.Update)
		sections.DELETE("/:id", projectSectionController.Delete)
	}
}

func RegisterSettingRoutes(r *gin.RouterGroup) {
	settingController := &controllers.SettingController{}
	settings := r.Group("/setting")
	{
		settings.GET("", settingController.Index)
		settings.POST("", settingController.Create)
		settings.GET("/:id", settingController.View)
		settings.PUT("/:id", settingController.Update)
		settings.PATCH("/:id", settingController.Update)
	}
}

func RegisterFileRoutes(r *gin.RouterGroup) {
	fileController := &controllers.FileController{}
	files := r.Group("/file")
	{
		files.GET("", fileController.Index)
		files.POST("", fileController.Create)
		files.GET("/:id", fileController.View)
		files.DELETE("/:id", fileController.Delete)
		files.POST("/table", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "File table endpoint"})
		})
		files.POST("/pysvc", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "File pysvc endpoint"})
		})
	}
}

func RegisterStatRoutes(r *gin.RouterGroup) {
	statController := &controllers.StatController{}
	stats := r.Group("/stat")
	{
		stats.GET("/table", statController.Table)
		stats.GET("/graph", statController.Graph)
		stats.GET("/groups", statController.Groups)
		stats.GET("/:id/table", statController.Table)
		stats.GET("/:id/graph", statController.Graph)
	}
}

func RegisterOPCRoutes(r *gin.RouterGroup) {
	opcController := &controllers.OPCController{}
	opc := r.Group("/opc")
	{
		servers := opc.Group("/server")
		{
			servers.GET("", opcController.ServerIndex)
			servers.GET("/:id", opcController.ServerView)
		}

		tags := opc.Group("/tag")
		{
			tags.GET("", opcController.TagIndex)
			tags.GET("/:id", opcController.TagView)
		}

		datatypes := opc.Group("/datatype")
		{
			datatypes.GET("", opcController.DatatypeIndex)
			datatypes.GET("/:id", opcController.DatatypeView)
		}
	}
}

func RegisterSVGRoutes(r *gin.RouterGroup) {
	svgController := &controllers.SVGController{}
	svg := r.Group("/svg/equipment")
	{
		svg.GET("/binds", svgController.Binds)
		svg.GET("/tags", svgController.Tags)
		svg.GET("/alerts", svgController.Alerts)
		svg.PATCH("/batch-edit", svgController.BatchEdit)
		svg.PATCH("/bind/batch-edit", svgController.BatchEdit)
		svg.POST("/entity/batch", svgController.BatchEdit)
		svg.POST("/binds-batch", svgController.BatchEdit)
	}
}

