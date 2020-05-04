package editing

import (
	"github.com/theiny/users-service/pkg/models"
	"go.uber.org/zap"
)

type Service struct {
	r   Repository
	Log *zap.SugaredLogger
}

type Repository interface {
	EditUser(id string, u *models.User) error
}

func NewService(r Repository, l *zap.SugaredLogger) *Service {
	return &Service{r: r, Log: l}
}

func (s *Service) EditUser(id string, u *models.User) error {
	return s.r.EditUser(id, u)
}
