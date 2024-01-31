package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go_clearn_architect/api/route"
	"go_clearn_architect/pkg/utils"
)

func main() {

	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	route.RouteSetup(e)
	e.Logger.Fatal(e.Start(":1213"))

}
