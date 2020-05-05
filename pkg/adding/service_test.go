package adding

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theiny/user-service/pkg/logger"
	"github.com/theiny/user-service/pkg/models"
	"github.com/theiny/user-service/pkg/repository"
)

func TestNewService(t *testing.T) {
	m := repository.InMemory()
	logger := logger.New("test")

	s := NewService(m, logger)

	assert.NotNil(t, s)
	assert.NotNil(t, s.Log)
}

func TestAddUser(t *testing.T) {
	m := repository.InMemory()
	logger := logger.New("test")

	s := NewService(m, logger)

	user := &models.User{
		FirstName: "John",
		LastName:  "Smith",
		Nickname:  "Pocahontas",
		Password:  "superSecurePassword",
		Email:     "john.smith@somedomain.com",
		Country:   "UK",
	}

	err := s.AddUser(user)
	assert.Nil(t, err)
}
