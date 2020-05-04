package user

import (
	"github.com/gin-gonic/gin"
	"github.com/theiny/users-service/user/adding"
	"github.com/theiny/users-service/user/deleting"
	"github.com/theiny/users-service/user/editing"
	"github.com/theiny/users-service/user/listing"
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
