package main

import (
	"github.com/theiny/users-service/user"
)

func main() {
	s := user.NewServer()

	s.NewLogger()
	s.NewRouter()

	s.Service = &user.UserService{
		Storage: user.InMemory(),
	}

	s.Log.Info("Starting app...")

	s.Router.Run(":8080")
}
