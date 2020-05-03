package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server struct {
	Router  *gin.Engine
	Service *UserService
	Log     *zap.SugaredLogger
}

func NewServer() *server {
	return &server{}
}
