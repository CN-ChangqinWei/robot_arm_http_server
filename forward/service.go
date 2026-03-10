package forward

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

type ForwardRepository interface {
	GetTopics(ctx context.Context) (res []domain.Forward, err error)
	GetTopicInfo(ctx context.Context, topic string) (res domain.Forward, err error)
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
