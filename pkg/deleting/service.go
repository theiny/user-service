package deleting

import (
	"go.uber.org/zap"
)

// Ideally we might want an iteration to interface the service layer as well as the repository to make it a truly hex architecture.

// Service encapsulates the storage type and logger instance.
type Service struct {
	r   Repository
	Log *zap.SugaredLogger
}

// Repository is an interface that provides access to the user storage.
type Repository interface {
	DeleteUser(id string) error
}

// NewService creates a new service for the adding user behaviour.
func NewService(r Repository, l *zap.SugaredLogger) *Service {
	return &Service{r: r, Log: l}
}

// DeleteUser removes the user with the given id from the repository.
func (s *Service) DeleteUser(id string) error {
	return s.r.DeleteUser(id)
}
