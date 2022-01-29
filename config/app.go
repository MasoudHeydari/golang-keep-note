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

func GetDbName() string {
	appConfig, _ := godotenv.Read()
	return appConfig["MYSQL_DB_NAME"]
}

func GetTokenSecretKey() string {
	// TODO cannot read values from .env file
	//s := os.Getenv("MYSQL_PORT")
	//fmt.Println("secret key: ", s)
	appConfig, _ := godotenv.Read()
	return appConfig["API_SECRET"]
}
