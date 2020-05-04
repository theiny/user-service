package adding

import (
	"github.com/theiny/users-service/pkg/models"
	"go.uber.org/zap"
)

type Service struct {
	r   Repository
	Log *zap.SugaredLogger
}

type Repository interface {
	AddUser(u *models.User) error
}

func NewService(r Repository, l *zap.SugaredLogger) *Service {
	return &Service{r: r, Log: l}
}

func (s *Service) AddUser(u *models.User) error {
	return s.r.AddUser(u)
}
