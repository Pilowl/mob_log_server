package repository

import (
	"log"

	"../models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

var IsInitialized = false

func GetDb() *gorm.DB {
	if IsInitialized {
		return db
	}
	log.Println("DB is not initialized.")
	return nil
}

func Init() {
	var err error
	db, err = gorm.Open("mysql", "root:1234@tcp(:3306)/log_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		IsInitialized = false
		panic("failed to connect to DB@. " + err.Error())
	} else {
		IsInitialized = true
	}

	db.AutoMigrate(&models.LogModel{})
}

func Close() {
	db.Close()
	log.Println("Database is closed.")
}