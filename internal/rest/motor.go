package rest

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/motor"
	"github.com/labstack/echo/v4"
)

type MotorHandler struct {
	Service *motor.Service
}

type MotorRequest struct {
	DeviceName string             `json:"device_name"`
	Data       domain.MotorDomain `json:"data"`
}

func NewHandler(e *echo.Echo, service *motor.Service) *MotorHandler {
	h := &MotorHandler{
		Service: service,
	}
	// 绑定 POST /motor/control 路由
	e.POST("/motor/control", h.ControlMotor)
	return h
}

// ControlMotor 处理电机控制请求
func (h *MotorHandler) ControlMotor(c echo.Context) error {
	var req MotorRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, domain.MotorDomainReply{
			Message: "Invalid request: " + err.Error(),
		})
	}

	// 使用设备名作为 topic
	topic := req.DeviceName

	// 调用 service 发送 MQTT 消息
	if err := h.Service.SetMotorStatus(topic, req.Data); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.MotorDomainReply{
			Message: "Failed to send message: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.MotorDomainReply{
		Message: "OK",
	})
}
