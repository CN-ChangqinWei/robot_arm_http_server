package mqttinfo

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	mqtt "github.com/mochi-mqtt/server/v2"
)

var topics = []domain.Forward{}

type ForwardRepository struct {
	Server *mqtt.Server
}

func (f *ForwardRepository) GetTopics(ctx context.Context) (res []domain.Forward, err error) {
	if len(topics) == 0 {
		return
	}
	for _, topic := range topics {
		res = append(res, domain.Forward{
			Topic: topic.Topic,
		})
	}
	return
}
func (f *ForwardRepository) GetTopicInfo(ctx context.Context, topic string) (res domain.Forward, err error) {

	for _, t := range topics {
		if t.Topic == topic {
			res = t
			return
		}
	}
	return
}
