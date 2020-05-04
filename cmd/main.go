package main

import (
	"github.com/theiny/users-service/pkg/adding"
	"github.com/theiny/users-service/pkg/deleting"
	"github.com/theiny/users-service/pkg/editing"
	"github.com/theiny/users-service/pkg/http/rest"
	"github.com/theiny/users-service/pkg/listing"
	"github.com/theiny/users-service/pkg/logger"
	"github.com/theiny/users-service/pkg/repository"
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
