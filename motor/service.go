package motor

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

// SetMotorStatus 设置电机状态，将 MotorDomain 序列化为 JSON 并通过 MQTT 发送到指定 topic
func (s *Service) SetMotorStatus(topic string, status domain.MotorDomain) error {
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
