package database

import (
	"errors"
	"fmt"
	"github.com/MasoudHeydari/golang-keep-note/config"
	"github.com/MasoudHeydari/golang-keep-note/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

// database credentials enumeration
const (
	createDatabase = iota
	connectToDatabase
)

func createDatabaseIfNoteExist(dbName string) {
	// parsing mysql credentials
	sqlCredentials, err := getMySqlCredentials(createDatabase)
	if err != nil {
		log.Fatal("err while parsing mysql credentials")
	}

	fmt.Println(sqlCredentials)
	db, err := gorm.Open("mysql", sqlCredentials)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	db.Exec("USE " + dbName)
}

func ConnectToDB() (*gorm.DB, error) {
	sqlCredentials, err := getMySqlCredentials(connectToDatabase)
	if err != nil {
		log.Fatal("err while parsing mysql credentials")
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

func getMySqlCredentials(mode int) (mySqlCredentials string, err error) {
	appConfig, err := config.GetAppConfig()
	if err != nil {
		log.Fatal("error while getting app config, error: ", err)
		return "", err
	}

	mySqlCredentials = fmt.Sprintf(
		"%s:%s@%s(%s:%s)/",
		appConfig["MYSQL_USER"],
		appConfig["MYSQL_PASSWORD"],
		appConfig["MYSQL_PROTOCOL"],
		appConfig["MYSQL_HOST"],
		appConfig["MYSQL_PORT"],
	)

	if mode == connectToDatabase {
		// Example: user:password@tcp(host:port)/db_name?charset=utf8&parseTime=True&loc=Local
		mySqlCredentials = fmt.Sprintf(
			"%s%s?charset=utf8&parseTime=True&loc=Local",
			mySqlCredentials,
			appConfig["MYSQL_DB_NAME"],
		)
		return
	} else if mode == createDatabase {
		// Example: user:password@tcp(host:port)/
		// do noting, just return base 'mySqlCredentials'
		return
	}

	// invalid mode
	return "", errors.New("provided mode is invalid")

}
