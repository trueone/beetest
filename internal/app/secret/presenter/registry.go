package presenter

import (
	"github.com/labstack/echo"

	"github.com/trueone/beetest/internal/app/secret/usecase"
)

// Registry creates use cases and registers http handlers
type Registry interface {
	RegisterJsonHandlers(e *echo.Echo)
	Create() usecase.Create
	Get() usecase.Get
}
