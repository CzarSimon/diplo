package httputil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONError sends an error response back to a client.
func JSONError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}

// RegisterHealthCheck registers a health check endpoint.
func RegisterHealthCheck(r *gin.Engine) {
	r.GET("/health", HandleHealthCheck)
}

// HandleHealthCheck responds possitively to a health check.
func HandleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
