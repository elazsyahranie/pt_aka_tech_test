package helpers

import (
	"log"

	"github.com/spf13/viper"
)

const (
	DUPLICATE    string = "duplicate"
	NOT_FOUND    string = "not found"
	UNAUTHORIZED string = "unauthorized"
	
	LOG_IN_REQUIRED string = "Please login first!"

	BOOK_ALREADY_EXIST string = "Book already exists!"
	USER_ALREADY_EXIST string = "User already exists!"
)

var PORT string = GetEnvVariables("PORT")
var DB_CONNECTION string = GetEnvVariables("DB_CONNECTION")
var SECRET_KEY = []byte(GetEnvVariables("TOKEN_SECRET_KEY"))

func GetEnvVariables(variable string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Variable unknown")
	}

	variable, ok := viper.Get(variable).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return variable
}