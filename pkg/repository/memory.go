package repository

import (
	"errors"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/theiny/users-service/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type memory struct {
	users map[string]*models.User
}

const (
	firstName = "first_name"
	lastName  = "last_name"
	nickname  = "nickname"
	email     = "email"
	country   = "country"
)

var errNotFound = errors.New("User not found")

func InMemory() *memory {
	return &memory{users: make(map[string]*models.User)}
}

func (m *memory) AddUser(u *models.User) error {
	var err error
	u.Password, err = hashPassword(u.Password)
	u.ID = uuid.New().String()
	m.users[u.ID] = u
	return err
}

func (m *memory) GetUsers(q url.Values) ([]*models.User, error) {
	return m.filter(q), nil
}

func (m *memory) EditUser(id string, u *models.User) error {
	if _, ok := m.users[id]; !ok {
		return errNotFound
	}
	m.users[id] = u
	return nil
}

func (m *memory) DeleteUser(id string) error {
	if _, ok := m.users[id]; !ok {
		return errNotFound
	}
	delete(m.users, id)
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// return it's key, otherwise it will return -1 and a bool of false.
func found(slice []string, val string) bool {
	var exists bool
	for _, item := range slice {
		if strings.ToLower(item) == strings.ToLower(val) {
			exists = true
			break
		}
	}
	return exists
}

func (m *memory) filter(q url.Values) []*models.User {
	var users []*models.User

	for _, u := range m.users {
		if len(q) != 0 {
			for filterBy, v := range q {
				switch filterBy {
				case firstName:
					if found(v, u.FirstName) {
						users = append(users, u)
					}
				case lastName:
					if found(v, u.LastName) {
						users = append(users, u)
					}
				case nickname:
					if found(v, u.Nickname) {
						users = append(users, u)
					}
				case email:
					if found(v, u.Email) {
						users = append(users, u)
					}
				case country:
					if found(v, u.Country) {
						users = append(users, u)
					}
				}
			}
		} else {
			users = append(users, u)
		}
	}

	return users
}
