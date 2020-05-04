package user

import (
	"net/url"

	"go.uber.org/zap"
)

type AddingService struct {
	Storage Repository
	Log     *zap.SugaredLogger
}

type ListingService struct {
	Storage Repository
	Log     *zap.SugaredLogger
}

type EditingService struct {
	Storage Repository
	Log     *zap.SugaredLogger
}

type DeletingService struct {
	Storage Repository
	Log     *zap.SugaredLogger
}

func (s *AddingService) AddUser(u User) error {
	return s.Storage.AddUser(u)
}

func (s *ListingService) GetUsers(q url.Values) ([]User, error) {
	return s.Storage.GetUsers(q)
}

func (s *EditingService) EditUser(id string, u User) error {
	return s.Storage.EditUser(id, u)
}

func (s *DeletingService) DeleteUser(id string) error {
	return s.Storage.DeleteUser(id)
}
