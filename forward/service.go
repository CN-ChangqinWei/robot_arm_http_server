package forward

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

type ForwardRepository interface {
	GetTopics(ctx context.Context) (res []domain.Forward, err error)
	GetTopicInfo(ctx context.Context, topic string) (res domain.Forward, err error)
	AddTopic(ctx context.Context, topic string) (err error)
	DelTopic(ctx context.Context, topic string) (err error)
	AddSubscriber(ctx context.Context, topic string, clientId string) (err error)
	DelSubscriber(ctx context.Context, topic string, clientId string) (err error)
}

type Service struct {
	forwardRepository ForwardRepository
}

func NewService(repo ForwardRepository) *Service {
	return &Service{
		forwardRepository: repo,
	}
}

func (s *Service) GetTopics(ctx context.Context) (res []domain.Forward, err error) {
	res, err = s.forwardRepository.GetTopics(ctx)
	if err != nil {
		return
	}
	return
}

func (s *Service) GetTopicInfo(ctx context.Context, topic string) (res domain.Forward, err error) {
	res, err = s.forwardRepository.GetTopicInfo(ctx, topic)
	if err != nil {
		return
	}
	return
}

func (s *Service) Subscribe(ctx context.Context, topic string, clientId string) (err error) {
	err = s.forwardRepository.AddSubscriber(ctx, topic, clientId)

	return
}
func (s *Service) Desubscribe(ctx context.Context, topic string, clientId string) (err error) {

	err = s.forwardRepository.DelSubscriber(ctx, topic, clientId)
	return
}

func (s *Service) Publish(ctx context.Context, topic string) (err error) {
	err = s.forwardRepository.AddTopic(ctx, topic)
	return
}

func (s *Service) Depublish(ctx context.Context, topic string) (err error) {
	err = s.forwardRepository.DelTopic(ctx, topic)
	return
}
