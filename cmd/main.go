package main

import (
	"log"

	"github.com/theiny/user-service/pkg/adding"
	"github.com/theiny/user-service/pkg/deleting"
	"github.com/theiny/user-service/pkg/editing"
	"github.com/theiny/user-service/pkg/http/rest"
	"github.com/theiny/user-service/pkg/listing"
	"github.com/theiny/user-service/pkg/logger"
	"github.com/theiny/user-service/pkg/repository"
)

const defaultPort = ":8080"

func main() {
	s := rest.NewServer()

	storageType := repository.InMemory()

	s.AddingService = adding.NewService(storageType, logger.New("user-add"))
	s.ListingService = listing.NewService(storageType, logger.New("user-list"))
	s.EditingService = editing.NewService(storageType, logger.New("user-edit"))
	s.DeletingService = deleting.NewService(storageType, logger.New("user-delete"))

	s.NewRouter()

	log.Printf("Starting app on http://localhost%s\n", defaultPort)

	s.Router.Run(defaultPort)
}
