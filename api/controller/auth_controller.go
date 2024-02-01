package controller

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go_clearn_architect/internal/adapter/repository"
	"go_clearn_architect/internal/domain/auth/presenter"
	"go_clearn_architect/internal/domain/auth/usecase"
	"go_clearn_architect/pkg/utils"
	"net/http"
)

var authRepository = repository.Auth{}

type LoginController struct{}

func (lc LoginController) Register(c echo.Context) error {
	u := new(presenter.SignUpRequest)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		errMsg := map[string]interface{}{}
		for _, e := range err.(validator.ValidationErrors) {
			tag := utils.GetJSONTag(*u, e.Field())

			switch e.Tag() {
			case "required":
				errMsg[tag] = fmt.Sprintf("%s is required", tag)
			case "min":
				errMsg[tag] = fmt.Sprintf("%s must be at least %s characters", tag, e.Param())
			case "max":
				errMsg[tag] = fmt.Sprintf("%s must be at most %s characters", tag, e.Param())
			case "email":
				errMsg[tag] = fmt.Sprintf("Invalid email address for %s", tag)
			}
		}
		return c.JSON(400, map[string]interface{}{"error": "Validation error", "details": errMsg})
	}

	accessToken, err := usecase.SignUp(u, authRepository)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"success": true, "data": accessToken})
}
