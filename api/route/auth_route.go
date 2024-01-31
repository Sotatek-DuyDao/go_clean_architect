package route

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go_clearn_architect/internal/auth/presenter"
	"net/http"
)

func NewAuthRoute(r *echo.Group) {
	authRoute := r.Group("/auth")

	authRoute.POST("/sign-up", func(c echo.Context) error {
		u := new(presenter.SignUpRequest)

		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}

		validate := validator.New()
		if err := validate.Struct(u); err != nil {
			return c.JSONPretty(http.StatusBadRequest, err.Error(), "")
		}

		return nil
	})
}
