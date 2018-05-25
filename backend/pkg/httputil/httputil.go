package httputil

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Common errors
var (
	ErrBadRequest          = NewError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	ErrInternalServerError = NewError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
)

// Error http error
type Error struct {
	StatusCode int
	Message    string
}

// NewError creates a new error.
func NewError(statusCode int, message string) Error {
	return Error{
		StatusCode: statusCode,
		Message:    message,
	}
}

// Error returns a string representation of an http error.
func (err Error) Error() string {
	return fmt.Sprintf("%d - %s", err.StatusCode, err.Message)
}

// JSONError sends an error response back to a client.
func JSONError(c *gin.Context, err error) {
	switch httpErr := err.(type) {
	case Error:
		c.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Message})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// SendOK sends an ok status and message to the client.
func SendOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"error": nil, "status": "OK"})
}

// RegisterHealthCheck registers a health check endpoint.
func RegisterHealthCheck(r *gin.Engine) {
	r.GET("/health", HandleHealthCheck)
}

// HandleHealthCheck responds possitively to a health check.
func HandleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
