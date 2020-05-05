package listing

import (
	"net/url"

	"go.uber.org/zap"

	"github.com/theiny/user-service/pkg/models"
)

// Ideally we might want an iteration to interface the service layer as well as the repository to make it a truly hex architecture.

// Service encapsulates the storage type and logger instance.
type Service struct {
	r   Repository
	Log *zap.SugaredLogger
}

// Repository is an interface that provides access to the user storage.
type Repository interface {
	GetUsers(q url.Values) ([]*models.User, error)
}

// NewService creates a new service for the adding user behaviour.
func NewService(r Repository, l *zap.SugaredLogger) *Service {
	return &Service{r: r, Log: l}
}

// GetUsers retrieves all users from storage that match the criteria specified by query params q.
func (s *Service) GetUsers(q url.Values) ([]*models.User, error) {
	return s.r.GetUsers(q)
}
