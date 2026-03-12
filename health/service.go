package health

import (
	"github.com/bxcodec/go-clean-arch/domain"
)

type HealthRepository interface {
}

type HealthService struct {
	repo *HealthRepository
}

func NewService() *HealthService {
	return &HealthService{}
}

func (*HealthService) IsHttpHealthy() (res domain.Health) {
	res.HttpServerOk = true
	return
}
