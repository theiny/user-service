package listing

import (
	"net/url"

	"go.uber.org/zap"

	"github.com/theiny/users-service/user/models"
)

type Service struct {
	Storage Repository
	Log     *zap.SugaredLogger
}

type Repository interface {
	GetUsers(q url.Values) ([]models.User, error)
}

func NewService(r Repository, l *zap.SugaredLogger) *Service {
	return &Service{Storage: r, Log: l}
}

func (s *Service) GetUsers(q url.Values) ([]models.User, error) {
	return s.Storage.GetUsers(q)
}
