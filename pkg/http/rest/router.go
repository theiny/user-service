package rest

import "github.com/gin-gonic/gin"

// NewRouter creates a new Gin engine and loads the handlers. 
func (s *server) NewRouter() {
	gin.SetMode(gin.ReleaseMode)
	s.Router = gin.Default()
	s.LoadHandlers()
}
