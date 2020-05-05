package editing

import (
	"github.com/theiny/user-service/pkg/models"
	"go.uber.org/zap"
)

// Service encapsulates the storage type and logger instance.
type Service struct {
	r   Repository
	Log *zap.SugaredLogger
}

// Repository is an interface that provides access to the user storage.
type Repository interface {
	EditUser(id string, u *models.User) error
}

// NewService creates a new service for the adding user behaviour.
func NewService(r Repository, l *zap.SugaredLogger) *Service {
	return &Service{r: r, Log: l}
}

// EditUser updates an existing user with the given id and new user details passed in.
func (s *Service) EditUser(id string, u *models.User) error {
	return s.r.EditUser(id, u)
}
