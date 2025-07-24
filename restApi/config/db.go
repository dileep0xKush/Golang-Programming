package config

import (
	"fmt"
	"log"
	"restApi/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Replace with your actual MySQL credentials
	dsn := "root:root@tcp(127.0.0.1:3306)/golang?parseTime=true"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to MySQL database: %v", err)
	}

	fmt.Println("✅ Connected to MySQL database!")

	// Auto migrate models
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("❌ AutoMigration failed: %v", err)
	}
}
