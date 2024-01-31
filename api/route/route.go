package route

import "github.com/labstack/echo/v4"

func RouteSetup(routerGroup *echo.Echo) {
	publicRouter := routerGroup.Group("")

	NewAuthRoute(publicRouter)
	//protectedRouter := routerGroup.Group("")

}
