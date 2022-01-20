package database

import (
	"fmt"
	"github.com/MasoudHeydari/golang-keep-note/config"
	"github.com/MasoudHeydari/golang-keep-note/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
	"log"
)

func ConnectToDB() (*gorm.DB, error) {
	sqlCredentials, err := getMySqlCredentials()
	if err != nil {
		log.Fatal("err while connecting to mysql database")
		return nil, err
	}

	fmt.Println(sqlCredentials)

	// connect to db
	dbConnection, err := gorm.Open("mysql", sqlCredentials)

	if err != nil {
		fmt.Println("err: ", err)
		panic("Failed to connect to the database")
		return nil, err
	}

	fmt.Println("mysql db opened successfully")

	// database migration
	dbConnection.AutoMigrate(&models.User{}, &models.Note{})
	return dbConnection, err
}

func getMySqlCredentials() (mySqlCredentials string, err error) {

	appConfig, err := config.GetAppConfig()
	if err != nil {
		log.Fatal("error while getting app config, error: ", err)
		return "", err
	}

	// Example: user:password@tcp(host:port)/db_name
	mySqlCredentials = fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		appConfig["MYSQL_USER"],
		appConfig["MYSQL_PASSWORD"],
		appConfig["MYSQL_PROTOCOL"],
		appConfig["MYSQL_HOST"],
		appConfig["MYSQL_PORT"],
		appConfig["MYSQL_DB_NAME"],
	)
	return
}
