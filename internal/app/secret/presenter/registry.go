package presenter

import (
	"github.com/labstack/echo"
	"google.golang.org/grpc"

	"github.com/trueone/beetest/internal/app/secret/usecase"
)

// Registry creates use cases and registers http handlers
type Registry interface {
	RegisterJSONHandlers(e *echo.Echo)
	RegisterGRPCHandlers(s *grpc.Server)
	Create() usecase.Create
	Get() usecase.Get
}
