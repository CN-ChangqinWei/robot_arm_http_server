package forward

import (
	"github.com/bxcodec/go-clean-arch/domain"
)

type ForwardRepository interface {
	GetTopics() (res []domain.Forward, err error)
	GetTopicInfo(topic string) (res domain.Forward, err error)
	AddTopic(topic string, publisher string) (err error)
	DelTopic(topic string) (err error)
	AddSubscriber(topic string, clientId string) (err error)
	DelSubscriber(topic string, clientId string) (err error)
}

type Service struct {
	forwardRepository ForwardRepository
}

func NewService(repo ForwardRepository) *Service {
	return &Service{
		forwardRepository: repo,
	}
}

func (s *Service) GetTopics() (res []domain.Forward, err error) {
	res, err = s.forwardRepository.GetTopics()
	if err != nil {
		return
	}
	return
}

func (s *Service) GetTopicInfo(topic string) (res domain.Forward, err error) {
	res, err = s.forwardRepository.GetTopicInfo(topic)
	if err != nil {
		return
	}
	return
}

func (s *Service) Subscribe(topic string, clientId string) (err error) {
	err = s.forwardRepository.AddSubscriber(topic, clientId)

	return
}
func (s *Service) Unsubscribe(topic string, clientId string) (err error) {

	err = s.forwardRepository.DelSubscriber(topic, clientId)
	return
}

func (s *Service) Connect(topic string, publisher string) (err error) {
	err = s.forwardRepository.AddTopic(topic, publisher)
	return
}

func (s *Service) Disconnect(topic string) (err error) {
	err = s.forwardRepository.DelTopic(topic)
	return
}
