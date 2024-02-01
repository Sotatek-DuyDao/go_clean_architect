package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"go_clearn_architect/internal/adapter/mysql/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (db *gorm.DB) {
	host := viper.Get("MYSQL_HOST")
	user := viper.Get("MYSQL_USER")
	pass := viper.Get("MYSQL_PASSWORD")
	port := viper.Get("MYSQL_PORT")
	dbName := viper.Get("MYSQL_DB_NAME")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

func DBAutoMigrate() {
	db := Connect()
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return
	}
}
