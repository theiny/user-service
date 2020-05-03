package user

import "github.com/gin-gonic/gin"

func (s *server) NewRouter() {
	gin.SetMode(gin.ReleaseMode)
	s.Router = gin.Default()
	s.LoadHandlers()
}
