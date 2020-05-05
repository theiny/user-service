package deleting

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theiny/user-service/pkg/adding"
	"github.com/theiny/user-service/pkg/listing"
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

func TestDeleteUser(t *testing.T) {
	m := repository.InMemory()
	logger := logger.New("test")

	a := adding.NewService(m, logger)

	user := &models.User{
		FirstName: "John",
		LastName:  "Smith",
		Nickname:  "Pocahontas",
		Password:  "superSecurePassword",
		Email:     "john.smith@somedomain.com",
		Country:   "UK",
	}

	a.AddUser(user)

	l := listing.NewService(m, logger)

	q := url.Values{}
	users, err := l.GetUsers(q)
	assert.Nil(t, err)

	d := NewService(m, logger)
	err = d.DeleteUser(users[0].ID)
	assert.Nil(t, err)
}
