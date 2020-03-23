package api

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo"

	"github.com/trueone/beetest/internal/app/secret"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func JsonStart() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	secretRegistry := secret.New()
	secretRegistry.RegisterJsonHandlers(e)

	e.Logger.Fatal(e.Start(":8080"))
}
