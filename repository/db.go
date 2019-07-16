package repository

import (
	"fmt"
	"log"

	"../config"
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
	var username = config.GetConfig().DB.Username
	var password = config.GetConfig().DB.Password
	var port = config.GetConfig().DB.Port
	var name = config.GetConfig().DB.Name

	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, port, name))
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
