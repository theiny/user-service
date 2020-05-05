package repository

import (
	"errors"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/theiny/user-service/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// store the data in memory as a map. 
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

// InMemory initiates a new map in memory to store user data. 
func InMemory() *memory {
	return &memory{users: make(map[string]*models.User)}
}

// AddUser adds a new user to memory. 
func (m *memory) AddUser(u *models.User) error {
	var err error
	u.Password, err = hashPassword(u.Password)
	u.ID = uuid.New().String()
	m.users[u.ID] = u
	return err
}

// GetUsers retrieves a list of users from memory, matching the criteria given by query params q. 
func (m *memory) GetUsers(q url.Values) ([]*models.User, error) {
	return m.filter(q), nil
}

// EditUser updates an existing user of given id, with new details passed as u. 
func (m *memory) EditUser(id string, u *models.User) error {
	if _, ok := m.users[id]; !ok {
		return errNotFound
	}
	m.users[id] = u
	return nil
}

// DeleteUser removes the user with the given id from memory. 
func (m *memory) DeleteUser(id string) error {
	if _, ok := m.users[id]; !ok {
		return errNotFound
	}
	delete(m.users, id)
	return nil
}

// passwords should be stored as hashes.
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// helper func to check if val is in slice. 
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

// if query params are passed, only retrieve the users that match that given criteria e.g. ?first_name=Brandon returns all users with the first name 'Brandon'. 
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
