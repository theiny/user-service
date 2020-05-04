package main

import (
	"github.com/theiny/users-service/user/adding"
	"github.com/theiny/users-service/user/deleting"
	"github.com/theiny/users-service/user/editing"
	"github.com/theiny/users-service/user/http/rest"
	"github.com/theiny/users-service/user/listing"
	"github.com/theiny/users-service/user/logger"
	"github.com/theiny/users-service/user/repository"
)

func main() {
	s := rest.NewServer()

	storageType := repository.InMemory()

	s.AddingService = adding.NewService(storageType, logger.New("user-add"))
	s.ListingService = listing.NewService(storageType, logger.New("user-list"))
	s.EditingService = editing.NewService(storageType, logger.New("user-edit"))
	s.DeletingService = deleting.NewService(storageType, logger.New("user-delete"))

	s.NewRouter()

	s.Router.Run(":8080")
}
