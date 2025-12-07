package main

import (
	"log"
	"os"

	"diflow/api/database"
	"diflow/api/middleware"
	"diflow/api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	
	_ "diflow/api/docs" // Swagger docs
)

// @title           Diflow API
// @version         1.0
// @description     REST API для управления зданиями, оборудованием и связанными ресурсами
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  admin@admin.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @securityDefinitions.apikey QueryAuth
// @in query
// @name access_token
// @description Access token as query parameter

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize database
	db, err := database.Initialize()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Set Gin mode
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Apply middleware
	r.Use(middleware.CORS())
	r.Use(middleware.ErrorHandler())

	// Health check endpoint
	// @Summary Health check
	// @Description Check if API is running
	// @Tags health
	// @Produce json
	// @Success 200 {object} map[string]string
	// @Router /health [get]
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Swagger documentation
	// Register swagger handler with proper URL
	swaggerURL := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))

	// API routes
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// Public routes (no auth required)
			routes.RegisterSSORoutes(v1)
			routes.RegisterOAuthRoutes(v1)

			// Protected routes (auth required)
			protected := v1.Group("")
			protected.Use(middleware.AuthMiddleware())
			{
				routes.RegisterUserRoutes(protected)
				routes.RegisterBuildingRoutes(protected)
				routes.RegisterEquipmentRoutes(protected)
				routes.RegisterFloorRoutes(protected)
				routes.RegisterCustomerRoutes(protected)
				routes.RegisterZoneRoutes(protected)
				routes.RegisterProjectSectionRoutes(protected)
				routes.RegisterSettingRoutes(protected)
				routes.RegisterFileRoutes(protected)
				routes.RegisterStatRoutes(protected)
				routes.RegisterOPCRoutes(protected)
				routes.RegisterSVGRoutes(protected)
			}
		}
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

