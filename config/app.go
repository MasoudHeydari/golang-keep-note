package config

import (
	"github.com/joho/godotenv"
	"log"
)

func GetAppConfig() (appConfig map[string]string, err error) {
	appConfig, err = godotenv.Read()

	if err != nil {
		log.Fatal("error while parsing .env file, Error: ", err)
		return nil, err
	}

	return
}
