package deleting

import (
	"go.uber.org/zap"
)

type Service struct {
	r   Repository
	Log *zap.SugaredLogger
}

type Repository interface {
	DeleteUser(id string) error
}

func NewService(r Repository, l *zap.SugaredLogger) *Service {
	return &Service{r: r, Log: l}
}

func (s *Service) DeleteUser(id string) error {
	return s.r.DeleteUser(id)
}
