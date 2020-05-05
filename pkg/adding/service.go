package adding

import (
	"github.com/theiny/user-service/pkg/models"
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
	AddUser(u *models.User) error
}

// NewService creates a new service for the adding user behaviour.
func NewService(r Repository, l *zap.SugaredLogger) *Service {
	return &Service{r: r, Log: l}
}

// AddUser stores a new user to the repository.
func (s *Service) AddUser(u *models.User) error {
	return s.r.AddUser(u)
}
