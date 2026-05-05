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

	payload, err := json.Marshal(status)
	if err != nil {
		return err
	}

	return s.Server.Publish(topic, payload, false, 0)
}

// SendMotionPositions 将运动轨迹位置列表转换为 RobotDomain 数组，通过 MQTT 发送给指定设备
func (s *Service) SendMotionPositions(motion domain.RobotMotionDomain) error {
	if s.Server == nil {
		return nil
	}

	// 将 Positions 展开为 RobotDomain 数组，protocol 设为 PROTO_ROBOT_POSITION
	positions := make([]domain.RobotDomain, 0, len(motion.Positions))
	for _, pos := range motion.Positions {
		positions = append(positions, domain.RobotDomain{
			Protocol: int(domain.PROTO_ROBOT_POSITION),
			X:        pos.X,
			Y:        pos.Y,
			Z:        pos.Z,
		})
	}

	// 序列化为 JSON
	payload, err := json.Marshal(positions)
	if err != nil {
		return err
	}

	// 通过 MQTT 发送给目标设备
	return s.Server.Publish(motion.Dev, payload, false, 0)
}
