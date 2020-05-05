package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/theiny/user-service/pkg/adding"
	"github.com/theiny/user-service/pkg/deleting"
	"github.com/theiny/user-service/pkg/editing"
	"github.com/theiny/user-service/pkg/listing"
)

type server struct {
	Router          *gin.Engine
	AddingService   *adding.Service
	ListingService  *listing.Service
	EditingService  *editing.Service
	DeletingService *deleting.Service
}

func NewServer() *server {
	return &server{}
}
