package editing

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theiny/users-service/pkg/adding"
	"github.com/theiny/users-service/pkg/listing"
	"github.com/theiny/users-service/pkg/logger"
	"github.com/theiny/users-service/pkg/models"
	"github.com/theiny/users-service/pkg/repository"
)

func TestNewService(t *testing.T) {
	m := repository.InMemory()
	logger := logger.New("test")

	s := NewService(m, logger)

	assert.NotNil(t, s)
	assert.NotNil(t, s.Log)
}

func TestEditUser(t *testing.T) {
	m := repository.InMemory()
	logger := logger.New("test")

	a := adding.NewService(m, logger)

	user := models.User{
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

	updated := "Doe"
	users[0].LastName = updated

	e := NewService(m, logger)
	err = e.EditUser(users[0].ID, users[0])
	assert.Nil(t, err)

	users, err = l.GetUsers(q)
	assert.Nil(t, err)

	assert.Equal(t, updated, users[0].LastName, "User details has not been updated")
}
