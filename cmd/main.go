package main

import (
	"github.com/theiny/users-service/user"
	"github.com/theiny/users-service/user/adding"
	"github.com/theiny/users-service/user/deleting"
	"github.com/theiny/users-service/user/editing"
	"github.com/theiny/users-service/user/listing"
)

func main() {
	s := user.NewServer()

	storageType := user.InMemory()

	s.AddingService = adding.NewService(storageType, user.NewLogger("add-user"))
	s.ListingService = listing.NewService(storageType, user.NewLogger("list-users"))
	s.EditingService = editing.NewService(storageType, user.NewLogger("edit-user"))
	s.DeletingService = deleting.NewService(storageType, user.NewLogger("delete-user"))

	s.NewRouter()

	s.Router.Run(":8080")
}
