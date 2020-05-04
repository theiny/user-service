package user

import (
	"github.com/gin-gonic/gin"
)

type server struct {
	Router          *gin.Engine
	AddingService   *AddingService
	ListingService  *ListingService
	EditingService  *EditingService
	DeletingService *DeletingService
}

func NewServer() *server {
	return &server{}
}
