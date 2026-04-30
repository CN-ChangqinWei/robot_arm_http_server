package robot

import (
	"encoding/json"

	"github.com/bxcodec/go-clean-arch/domain"
	mqtt "github.com/mochi-mqtt/server/v2"
)

type Service struct {
	Server *mqtt.Server
}

func NewService(server *mqtt.Server) *Service {
	return &Service{
		Server: server,
	}
}

func (s *Service) SetRobotStatus(topic string, status domain.RobotDomain) error {
	if s.Server == nil {
		return nil
	}

	// 将 MotorDomain 序列化为 JSON
	payload, err := json.Marshal(status)
	if err != nil {
		return err
	}

	// 通过 MQTT 发布消息
	return s.Server.Publish(topic, payload, false, 0)
}
