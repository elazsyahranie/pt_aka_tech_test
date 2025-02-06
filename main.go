package main

import (
	"fmt"
	"log"
	"time"

	"pt_aka_tech_test/helpers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


func main() {
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // Output logs to terminal
		logger.Config{
			SlowThreshold:             time.Second, // Highlight slow queries
			LogLevel:                  logger.Info, // Set log level to Info
			IgnoreRecordNotFoundError: true,        // Ignore 'record not found' errors
			Colorful:                  true,        // Colorful output
		},
	)

	dsn := helpers.DB_CONNECTION
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("Cannot connect to database!")
	}
	// db.AutoMigrate(&book.Book{})
	fmt.Println("Succesfuly connected to the database!")
	
}