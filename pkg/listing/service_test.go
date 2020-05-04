package listing

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theiny/users-service/pkg/adding"
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

func TestGetUsers(t *testing.T) {
	m := repository.InMemory()
	logger := logger.New("test")

	a := adding.NewService(m, logger)

	users := []models.User{
		models.User{
			FirstName: "John",
			LastName:  "Smith",
			Nickname:  "Pocahontas",
			Password:  "superSecurePassword",
			Email:     "john.smith@somedomain.com",
			Country:   "UK",
		},
		models.User{
			FirstName: "Jane",
			LastName:  "Doe",
			Nickname:  "JD",
			Password:  "superSecurePassword",
			Email:     "jane.doe@somedomain.com",
			Country:   "US",
		},
		models.User{
			FirstName: "John",
			LastName:  "Doe",
			Nickname:  "JD",
			Password:  "superSecurePassword",
			Email:     "john.doe@somedomain.com",
			Country:   "US",
		},
	}

	tests := []struct {
		key     string
		value   string
		results int
	}{
		{
			key:     "country",
			value:   "UK",
			results: 1,
		},
		{
			key:     "country",
			value:   "US",
			results: 2,
		},
		{
			key:     "first_name",
			value:   "Jane",
			results: 1,
		},
		{
			key:     "last_name",
			value:   "Doe",
			results: 2,
		},
		{
			key:     "nickname",
			value:   "Boaty McBoatface",
			results: 0,
		},
	}

	for _, user := range users {
		a.AddUser(user)
	}

	l := NewService(m, logger)

	var filtered []models.User
	for _, test := range tests {
		q := url.Values{}
		q.Add(test.key, test.value)
		filtered, _ = l.GetUsers(q)
		assert.Equal(t, test.results, len(filtered))
	}
}
