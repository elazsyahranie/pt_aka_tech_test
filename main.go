package main

import (
	"fmt"
	"log"
	"time"

	"pt_aka_tech_test/helpers"
	"pt_aka_tech_test/users"

	"github.com/gin-gonic/gin"
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("Cannot connect to database!")
	}
	// db.AutoMigrate(&book.Book{})
	fmt.Println("Succesfuly connected to the database!")

	router := gin.Default()

	userRepository := users.NewRepository(db) 
	userService := users.NewService(userRepository)
	userHandler := users.NewUserHandler(userService)

	userRoute := router.Group("/users")

	userRoute.POST("/", userHandler.Create)

	router.Run(helpers.PORT) // Will use default port 8000 if it's left empty
}