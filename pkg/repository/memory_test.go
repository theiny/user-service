package repository

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theiny/users-service/pkg/models"
)

func TestAddUser(t *testing.T) {
	m := InMemory()

	user := &models.User{
		FirstName: "John",
		LastName:  "Doe",
		Nickname:  "JD",
		Password:  "superSecurePassword",
		Email:     "john.doe@somedomain.com",
		Country:   "US",
	}

	err := m.AddUser(user)
	assert.Nil(t, err)
}

func TestGetUsers(t *testing.T) {
	m := InMemory()

	user := &models.User{
		FirstName: "John",
		LastName:  "Doe",
		Nickname:  "JD",
		Password:  "superSecurePassword",
		Email:     "john.doe@somedomain.com",
		Country:   "US",
	}

	err := m.AddUser(user)
	assert.Nil(t, err)

	q := url.Values{}
	users, err := m.GetUsers(q)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(users), "Number of entries do not match")
}

func TestEditUser(t *testing.T) {
	m := InMemory()

	user := &models.User{
		FirstName: "John",
		LastName:  "Doe",
		Nickname:  "JD",
		Password:  "superSecurePassword",
		Email:     "john.doe@somedomain.com",
		Country:   "US",
	}

	err := m.AddUser(user)
	assert.Nil(t, err)

	q := url.Values{}
	users, err := m.GetUsers(q)

	updated := "Tyrian Lannister"
	users[0].Nickname = updated
	err = m.EditUser(users[0].ID, users[0])

	assert.Nil(t, err)
	assert.Equal(t, updated, users[0].Nickname, "User field has not been updated")
}

func TestDeleteUser(t *testing.T) {
	m := InMemory()

	user := &models.User{
		FirstName: "John",
		LastName:  "Doe",
		Nickname:  "JD",
		Password:  "superSecurePassword",
		Email:     "john.doe@somedomain.com",
		Country:   "US",
	}

	err := m.AddUser(user)
	assert.Nil(t, err)

	q := url.Values{}
	users, err := m.GetUsers(q)

	err = m.DeleteUser(users[0].ID)
	assert.Nil(t, err)
}
