package models

import "github.com/jinzhu/gorm"

type SqlStore struct {
	db *gorm.DB
}

func NewSqlStore(dbConnection *gorm.DB) SqlQuerier {
	return &SqlStore{
		db: dbConnection,
	}
}
