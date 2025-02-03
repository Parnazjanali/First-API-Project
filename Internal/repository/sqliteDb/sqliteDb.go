package sqliteDb

import (
	"log"
	"simple-api/Internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDb() {
	var err error
	Db, err = gorm.Open(sqlite.Open("contact.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected Successfully")
	err = Db.AutoMigrate(&models.Contact{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
