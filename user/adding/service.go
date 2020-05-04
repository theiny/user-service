package adding

import (
	"github.com/theiny/users-service/user/models"
	"go.uber.org/zap"
)

type Service struct {
	Storage Repository
	Log     *zap.SugaredLogger
}

type Repository interface {
	AddUser(u models.User) error
}

func NewService(r Repository, l *zap.SugaredLogger) *Service {
	return &Service{Storage: r, Log: l}
}

func (s *Service) AddUser(u models.User) error {
	return s.Storage.AddUser(u)
}
