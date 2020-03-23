package api

import (
	"github.com/labstack/echo"

	"github.com/trueone/beetest/internal/app/secret"
)

func JsonStart() {
	e := echo.New()

	secretRegistry := secret.New()
	secretRegistry.RegisterJsonHandlers(e)

	e.Logger.Fatal(e.Start(":8080"))
}
