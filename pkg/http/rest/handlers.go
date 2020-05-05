package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theiny/user-service/pkg/adding"
	"github.com/theiny/user-service/pkg/deleting"
	"github.com/theiny/user-service/pkg/editing"
	"github.com/theiny/user-service/pkg/listing"
	"github.com/theiny/user-service/pkg/models"
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

func Error(c *gin.Context, code int, err error, msg string) {
	e := &ErrorResponse{
		Error:   err.Error(),
		Message: msg,
	}
	c.JSON(code, e)
}

func Respond(c *gin.Context, msg string) {
	r := &JSONResponse{
		StatusCode: http.StatusOK,
		Message:    msg,
	}
	c.JSON(r.StatusCode, r)
}

func (s *server) LoadHandlers() {
	users := s.Router.Group("/api/v1/users")
	{
		users.POST("/add", handleUserAdd(s.AddingService))
		users.GET("/get", handleUserGet(s.ListingService))
		users.PUT("/edit/:id", handleUserEdit(s.EditingService))
		users.DELETE("/delete/:id", handleUserDelete(s.DeletingService))
	}
	s.Router.GET("/healthcheck", handleHealthcheck)
}

func handleHealthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "I'm up!"})
}

func handleUserAdd(s *adding.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var u *models.User
		err := json.NewDecoder(c.Request.Body).Decode(&u)
		if err != nil {
			s.Log.Error(err)
			Error(c, http.StatusInternalServerError, err, "Error decoding user JSON")
			return
		}

		err = s.AddUser(u)
		if err != nil {
			s.Log.Error(err)
			Error(c, http.StatusInternalServerError, err, "Error adding new user")
			return
		}

		s.Log.Info("Added new user")
		Respond(c, "Successfully added new user")
	}
}

func handleUserGet(s *listing.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Request.URL.Query()

		users, err := s.GetUsers(query)
		if err != nil {
			s.Log.Error(err)
			Error(c, http.StatusInternalServerError, err, "Error getting users")
			return
		}

		if len(users) == 0 {
			s.Log.Info("No users found")
			c.JSON(http.StatusOK, gin.H{"message": "No users found"})
			return
		}

		s.Log.Debug("Retrieving users")
		c.JSON(http.StatusOK, users)
	}
}

func handleUserEdit(s *editing.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			s.Log.Error(errMissingID)
			Error(c, http.StatusBadRequest, errMissingID, "ID unknown")
			return
		}

		var u *models.User
		err := json.NewDecoder(c.Request.Body).Decode(&u)
		if err != nil {
			s.Log.Error(err)
			Error(c, http.StatusInternalServerError, err, "Error decoding user JSON")
			return
		}

		err = s.EditUser(id, u)
		if err != nil {
			s.Log.Error(err)
			Error(c, http.StatusInternalServerError, err, "Error editing user")
			return
		}

		s.Log.Info("Edited user")
		Respond(c, "Successfully Edited user")
	}
}

func handleUserDelete(s *deleting.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			s.Log.Error(errMissingID)
			Error(c, http.StatusBadRequest, errMissingID, "ID unknown")
			return
		}

		err := s.DeleteUser(id)
		if err != nil {
			s.Log.Error(err)
			Error(c, http.StatusInternalServerError, err, "Error deleting user")
			return
		}

		s.Log.Info("Deleted user")
		Respond(c, "Successfully deleted user")
	}
}
