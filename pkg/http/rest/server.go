package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/theiny/user-service/pkg/adding"
	"github.com/theiny/user-service/pkg/deleting"
	"github.com/theiny/user-service/pkg/editing"
	"github.com/theiny/user-service/pkg/listing"
)

// server encapsulates all the dependencies required for the service to run on a REST api.
type server struct {
	Router          *gin.Engine
	AddingService   *adding.Service
	ListingService  *listing.Service
	EditingService  *editing.Service
	DeletingService *deleting.Service
}

// NewServer returns an empty server struct. Dependencies should be injected from elsewhere like main.
func NewServer() *server {
	return &server{}
}
