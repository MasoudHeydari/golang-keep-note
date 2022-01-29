package database

import (
	"fmt"
	"github.com/MasoudHeydari/golang-keep-note/config"
	"github.com/MasoudHeydari/golang-keep-note/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func createDatabaseIfNoteExist(dbName string) {
	db, err := gorm.Open("mysql", "root:MySQL*Pass4883@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println("here1")
		panic(err)
	}
	defer db.Close()

	db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	db.Exec("USE " + dbName)
}

func ConnectToDB() (*gorm.DB, error) {
	sqlCredentials, err := getMySqlCredentials()
	if err != nil {
		log.Fatal("err while connecting to mysql database")
		return nil, err
	}

	fmt.Println(sqlCredentials)
	createDatabaseIfNoteExist(config.GetDbName())

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
