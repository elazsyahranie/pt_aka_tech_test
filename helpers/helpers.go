package helpers

import (
	"log"

	"github.com/spf13/viper"
)

var PORT string = GetEnvVariables("PORT")
var DB_CONNECTION string = GetEnvVariables("DB_CONNECTION")

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