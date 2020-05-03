package user

import "net/url"

type Service interface {
	AddUser(u User) error
	GetUsers(q url.Values) ([]User, error)
	EditUser(u User) error
	DeleteUser(id string) error
}
