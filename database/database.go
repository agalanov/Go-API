package database

import (
	"fmt"
	"log"

	"diflow/api/config"
	"diflow/api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Initialize() (*gorm.DB, error) {
	cfg := config.Load()
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db
	models.SetDB(db)
	log.Println("Database connection established")
	return db, nil
}

func Migrate(db *gorm.DB) error {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&models.User{},
		&models.Token{},
		&models.Customer{},
		&models.Building{},
		&models.Floor{},
		&models.FloorMode{},
		&models.Zone{},
		&models.ZoneGroup{},
		&models.ProjectSection{},
		&models.Equipment{},
		&models.EquipmentType{},
		&models.EquipmentEntity{},
		&models.EquipmentBind{},
		&models.EquipmentDirection{},
		&models.EquipmentLimit{},
		&models.EquipmentAttribute{},
		&models.EquipmentDevice{},
		&models.EquipmentState{},
		&models.Setting{},
		&models.File{},
		&models.OPCServer{},
		&models.OPCTag{},
		&models.OPCDatatype{},
	)

	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Database migrations completed")
	return nil
}

