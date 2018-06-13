package httputil

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Common errors
var (
	ErrBadRequest          = NewError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	ErrInternalServerError = NewError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	ErrUserNotFound        = NewError(http.StatusBadRequest, "User not found")
	ErrAccessDenied        = NewError(http.StatusForbidden, "Access denied")
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
	log.Println(err)
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

// GetUserID gets userID in header set by the authentication middleware.
func GetUserID(c *gin.Context) string {
	return c.GetString(UserHeader)
}

// RegisterHealthCheck registers a health check endpoint.
func RegisterHealthCheck(r *gin.Engine) {
	r.GET("/health", HandleHealthCheck)
}

// HandleHealthCheck responds possitively to a health check.
func HandleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}

// ParseQueryValue parses a query value from request.
func ParseQueryValue(c *gin.Context, key string) (string, error) {
	value, ok := c.GetQuery(key)
	if !ok {
		return "", NewError(http.StatusBadRequest,
			fmt.Sprintf("No value found for key: %s", key))
	}
	return value, nil
}

// ParseQueryValues parses query values from a request
func ParseQueryValues(c *gin.Context, key string) ([]string, error) {
	values, ok := c.GetQueryArray(key)
	if !ok {
		return nil, NewError(http.StatusBadRequest,
			fmt.Sprintf("No values found for key: %s", key))
	}
	return values, nil
}

// ParseTimestampQuery parses a unix timestamp.
func ParseTimestampQuery(c *gin.Context, key string) (time.Time, error) {
	value, err := ParseQueryValue(c, key)
	if err != nil {
		return time.Time{}, err
	}

	timestamp, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(timestamp, 0), nil
}
