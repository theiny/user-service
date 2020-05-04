package main

import (
	"github.com/theiny/users-service/user"
)

func main() {
	s := user.NewServer()

	storageType := user.InMemory()

	s.AddingService = &user.AddingService{
		Storage: storageType,
		Log:     user.NewLogger("add-user"),
	}

	s.ListingService = &user.ListingService{
		Storage: storageType,
		Log:     user.NewLogger("list-users"),
	}

	s.EditingService = &user.EditingService{
		Storage: storageType,
		Log:     user.NewLogger("edit-user"),
	}

	s.DeletingService = &user.DeletingService{
		Storage: storageType,
		Log:     user.NewLogger("delete-user"),
	}

	s.NewRouter()

	s.Router.Run(":8080")
}
