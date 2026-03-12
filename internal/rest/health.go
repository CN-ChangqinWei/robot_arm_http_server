package rest

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/labstack/echo/v4"
)

type HealthService interface {
	IsHttpHealthy() (res domain.Health)
}
type HealthHandler struct {
	Service HealthService
}

func NewHealthHandler(e *echo.Echo, svc HealthService) {
	handler := HealthHandler{
		Service: svc,
	}
	e.GET("/health", handler.IsHttpHealthy)
}

func (h *HealthHandler) IsHttpHealthy(e echo.Context) (err error) {
	res := h.Service.IsHttpHealthy()
	e.JSON(http.StatusOK, res)
	return
}
