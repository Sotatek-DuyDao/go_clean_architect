package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"go_clearn_architect/api/route"
	"go_clearn_architect/internal/adapter/mysql"
	"go_clearn_architect/pkg/utils"
	"log"
)

func main() {

	//Start server
	e := echo.New()
	e.Use(middleware.Logger())
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	//Setup env
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	//Auto migrate DB
	mysql.DBAutoMigrate()
	//Setup router
	route.RouteSetup(e)
	e.Logger.Fatal(e.Start(":1213"))

}
