package mqttinfo

import (
	"sync"

	"github.com/bxcodec/go-clean-arch/domain"
)

type ForwardRepository struct {
	mu     sync.RWMutex
	topics map[string]*domain.Forward
}

func NewForwardRepository() *ForwardRepository {
	return &ForwardRepository{
		topics: make(map[string]*domain.Forward),
	}
}

func (f *ForwardRepository) GetTopics() (res []domain.Forward, err error) {
	if f == nil {
		return nil, nil
	}
	f.mu.RLock()
	defer f.mu.RUnlock()

	n := len(f.topics)
	if n == 0 {
		return []domain.Forward{}, nil
	}

	res = make([]domain.Forward, 0, n)
	for _, v := range f.topics {
		res = append(res, *v)
	}
	return res, nil
}
func (f *ForwardRepository) GetTopicInfo(topic string) (res domain.Forward, err error) {
	if f == nil {
		return res, nil
	}
	f.mu.RLock()
	defer f.mu.RUnlock()

	if v, ok := f.topics[topic]; ok {
		return *v, nil
	}
	return res, nil
}
func (f *ForwardRepository) AddTopic(topic string, publisher string) error {
	if f == nil {
		return nil
	}
	f.mu.Lock()
	defer f.mu.Unlock()

	if _, ok := f.topics[topic]; ok {
		return nil
	}
	f.topics[topic] = &domain.Forward{
		Topic:     topic,
		Publisher: publisher,
	}
	return nil
}
func (f *ForwardRepository) DelTopic(topic string) error {
	if f == nil {
		return nil
	}
	f.mu.Lock()
	defer f.mu.Unlock()

	delete(f.topics, topic)
	return nil
}
func (f *ForwardRepository) AddSubscriber(topic string, clientId string) error {
	if f == nil {
		return nil
	}
	f.mu.Lock()
	defer f.mu.Unlock()

	if t, ok := f.topics[topic]; ok {
		t.Subscribers = append(t.Subscribers, clientId)
	}
	return nil
}
func (f *ForwardRepository) DelSubscriber(topic string, clientId string) error {
	if f == nil {
		return nil
	}
	f.mu.Lock()
	defer f.mu.Unlock()

	t, ok := f.topics[topic]
	if !ok {
		return nil
	}

	subs := t.Subscribers
	for i, id := range subs {
		if id == clientId {
			t.Subscribers = append(subs[:i], subs[i+1:]...)
			return nil
		}
	}
	return nil
}
