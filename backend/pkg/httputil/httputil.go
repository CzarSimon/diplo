package httputil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONError sends an error response back to a client.
func JSONError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
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
