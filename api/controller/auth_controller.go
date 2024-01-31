package controller

import "github.com/labstack/echo/v4"

type LoginController struct {
	loginUserCase interface{}
}

func (lc LoginController) Login(c echo.Context) {

}
