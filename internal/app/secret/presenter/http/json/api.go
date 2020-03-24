package json

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/trueone/beetest/internal/app/secret/data"
	"github.com/trueone/beetest/internal/app/secret/presenter"
)

type Handler struct {
	registry presenter.Registry
}

func New(registry presenter.Registry) *Handler {
	return &Handler{
		registry: registry,
	}
}

func Register(e *echo.Echo, h *Handler) {
	e.POST("/v1/secret", h.Create)
	e.GET("/v1/secret/:hash", h.Get)
}

func (h *Handler) Create(c echo.Context) error {
	request := new(data.SecretRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusMethodNotAllowed, "Invalid input")
	}

	s, err := h.registry.Create().Execute(*request)
	if err != nil {
		return c.JSON(http.StatusMethodNotAllowed, "Invalid input")
	}

	response := new(data.SecretResponse)
	response.Deserialize(s)

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) Get(c echo.Context) error {
	hash := data.HashRequest(c.Param("hash"))

	s, err := h.registry.Get().Execute(hash)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Secret not found")
	}

	response := new(data.SecretResponse)
	response.Deserialize(s)

	return c.JSON(http.StatusOK, response)
}
