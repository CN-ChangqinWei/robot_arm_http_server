package forward

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

type ForwardMqtt interface {
	GetTopics(ctx context.Context) (res []domain.Forward, err error)
	GetTopicInfo(ctx context.Context, topic string) (res domain.Forward, err error)
}

type Service struct {
	forwardMqtt ForwardMqtt
}

func NewService(fMqtt ForwardMqtt) *Service {
	return &Service{
		forwardMqtt: fMqtt,
	}
}

func (s *Service) GetTopics(ctx context.Context) (res []domain.Forward, err error) {
	res, err = s.forwardMqtt.GetTopics(ctx)
	if err != nil {
		return
	}
	return
}

func (s *Service) GetTopicInfo(ctx context.Context, topic string) (res domain.Forward, err error) {
	res, err = s.forwardMqtt.GetTopicInfo(ctx, topic)
	if err != nil {
		return
	}
	return
}
