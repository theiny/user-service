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

// app will run on this port.
const defaultPort = ":8080"

func main() {
	// NewServer returns a new server struct.
	// This encapsulates all the dependencies required for the service to run.
	s := rest.NewServer()

	/* We can declare the storage type we want to use here.
	** In this example, we use an in-memory data store, but we can swap it out with a sql database,
	** so long as the implementation satisfies the Repository interfaces.
	 */
	storageType := repository.InMemory()

	/* A new service encapsulates the storage, as well as it's own instance of a logger for that behaviour.
	** We create a new logger for each one so we can add labels to them which can be used by an external actor e.g. notifying a search service that a 'user-add' event has occurred.
	 */
	s.AddingService = adding.NewService(storageType, logger.New("user-add"))
	s.ListingService = listing.NewService(storageType, logger.New("user-list"))
	s.EditingService = editing.NewService(storageType, logger.New("user-edit"))
	s.DeletingService = deleting.NewService(storageType, logger.New("user-delete"))

	// creates a new Gin router engine, and loads the handlers for a REST API implementation.
	s.NewRouter()

	log.Printf("Starting app on http://localhost%s\n", defaultPort)

	// run the server.
	s.Router.Run(defaultPort)
}
