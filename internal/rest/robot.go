package rest

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/robot"
	"github.com/labstack/echo/v4"
)

type RobotHandler struct {
	Service *robot.Service
}

type RobotRequest struct {
	DeviceName string            `json:"device_name"`
	Data       domain.RobotDomain `json:"data"`
}

func NewRobotHandler(e *echo.Echo, service *robot.Service) *RobotHandler {
	h := &RobotHandler{
		Service: service,
	}
	// 绑定 POST /robot/control 路由
	e.POST("/robot/control", h.ControlRobot)
	return h
}

// ControlRobot 处理机器人控制请求
func (h *RobotHandler) ControlRobot(c echo.Context) error {
	var req RobotRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, domain.MotorDomainReply{
			Message: "Invalid request: " + err.Error(),
		})
	}

	// 使用设备名作为 topic
	topic := req.DeviceName

	// 调用 service 发送 MQTT 消息
	if err := h.Service.SetRobotStatus(topic, req.Data); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.MotorDomainReply{
			Message: "Failed to send message: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.MotorDomainReply{
		Message: "OK",
	})
}
