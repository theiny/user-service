package user

import (
	"net/url"
	"github.com/theiny/users-service/user/models"
)

type Service interface {
	AddUser(u models.User) error
	GetUsers(q url.Values) ([]models.User, error)
	EditUser(u models.User) error
	DeleteUser(id string) error
}
