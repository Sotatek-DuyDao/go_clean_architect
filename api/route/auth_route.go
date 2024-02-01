package route

import (
	"github.com/labstack/echo/v4"
	"go_clearn_architect/api/controller"
)

var loginController = controller.LoginController{}

func NewAuthRoute(r *echo.Group) {
	authRoute := r.Group("/auth")

	authRoute.POST("/sign-up", loginController.Register)
}
