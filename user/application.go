package user

import "net/url"

type UserService struct {
	Storage Repository
}

func (s *UserService) AddUser(u User) error {
	return s.Storage.AddUser(u)
}

func (s *UserService) GetUsers(q url.Values) ([]User, error) {
	return s.Storage.GetUsers(q)
}

func (s *UserService) EditUser(u User) error {
	return s.Storage.EditUser(u)
}

func (s *UserService) DeleteUser(id string) error {
	return s.Storage.DeleteUser(id)
}
