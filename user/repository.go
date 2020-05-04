package user

import "net/url"

type Repository interface {
	AddUser(u User) error
	GetUsers(q url.Values) ([]User, error)
	EditUser(id string, u User) error
	DeleteUser(id string) error
}
