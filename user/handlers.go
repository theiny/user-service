package user

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var errMissingID = errors.New("Missing id param")

// ErrorResponse is Error response template
type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"reason"`
}

type JSONResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (s *server) Error(c *gin.Context, code int, err error, msg string) {
	log.Println(err)
	s.Log.Errorw(msg, "error", err)
	e := &ErrorResponse{
		Error:   err.Error(),
		Message: msg,
	}
	c.JSON(code, e)
}

func (s *server) Respond(c *gin.Context, msg string) {
	s.Log.Info(msg)
	r := &JSONResponse{
		StatusCode: http.StatusOK,
		Message:    msg,
	}
	c.JSON(r.StatusCode, r)
}

func (s *server) LoadHandlers() {
	users := s.Router.Group("/api/v1/users")
	{
		users.POST("/add", s.handleUserAdd)
		users.GET("/get", s.handleUserGet)
		users.PUT("/edit", s.handleUserEdit)
		users.DELETE("/delete/:id", s.handleUserDelete)
	}
	s.Router.GET("/healthcheck", s.handleHealthcheck)
}

func (s *server) handleHealthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "I'm up!"})
}

func (s *server) handleUserAdd(c *gin.Context) {
	var u User
	err := json.NewDecoder(c.Request.Body).Decode(&u)
	if err != nil {
		s.Error(c, http.StatusInternalServerError, err, "Error decoding user JSON")
		return
	}

	err = s.Service.Storage.AddUser(u)
	if err != nil {
		s.Error(c, http.StatusInternalServerError, err, "Error adding new user")
		return
	}

	s.Respond(c, "Successfully added new user")
}

func (s *server) handleUserGet(c *gin.Context) {
	query := c.Request.URL.Query()

	users, err := s.Service.Storage.GetUsers(query)
	if err != nil {
		s.Error(c, http.StatusInternalServerError, err, "Error getting users")
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No users found"})
		return
	}

	s.Log.Debug("Retrieving users")
	c.JSON(http.StatusOK, users)
}

func (s *server) handleUserEdit(c *gin.Context) {
	var u User
	err := json.NewDecoder(c.Request.Body).Decode(&u)
	if err != nil {
		s.Error(c, http.StatusInternalServerError, err, "Error decoding user JSON")
		return
	}

	err = s.Service.Storage.EditUser(u)
	if err != nil {
		s.Error(c, http.StatusInternalServerError, err, "Error editing user")
		return
	}

	s.Respond(c, "Successfully Edited user")
}

func (s *server) handleUserDelete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		s.Error(c, http.StatusBadRequest, errMissingID, "ID unknown")
		return
	}

	err := s.Service.Storage.DeleteUser(id)
	if err != nil {
		s.Error(c, http.StatusInternalServerError, err, "Error deleting user")
		return
	}

	s.Respond(c, "Successfully deleted user")
}
