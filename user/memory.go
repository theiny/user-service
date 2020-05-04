package user

import (
	"errors"
	"net/url"
	"strings"
	"log"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type memRepository struct {
	users map[string]User
}

var errNotFound = errors.New("User not found")

func InMemory() *memRepository {
	return &memRepository{users: make(map[string]User)}
}

func (m *memRepository) AddUser(u User) error {
	var err error
	u.Password, err = hashPassword(u.Password)
	u.ID = uuid.New().String()
	m.users[u.ID] = u
	return err
}

func (m *memRepository) GetUsers(q url.Values) ([]User, error) {
	return m.filter(q), nil
}

func (m *memRepository) EditUser(id string, u User) error {
	if _, ok := m.users[id]; !ok {
		return errNotFound
	}
	m.users[id] = u
	return nil
}

func (m *memRepository) DeleteUser(id string) error {
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

func (m *memRepository) filter(q url.Values) []User {
	var users []User

	for _, u := range m.users {
		if len(q) != 0 {
			for k, v := range q {
				switch k {
				case "first_name":
					if found(v, u.FirstName) {
						users = append(users, u)
					}
				case "last_name":
					if found(v, u.LastName) {
						users = append(users, u)
					}
				case "nickname":
					if found(v, u.Nickname) {
						users = append(users, u)
					}
				case "email":
					if found(v, u.Email) {
						users = append(users, u)
					}
				case "country":
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
