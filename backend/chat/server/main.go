package main

import (
	"log"

	"github.com/CzarSimon/diplo/backend/pkg/httputil"
	"github.com/gin-gonic/gin"
)

// registerRoutes sets up available routes for the service.
func registerRoutes(r *gin.Engine) {
	registerMessageRoutes(r)
	httputil.RegisterHealthCheck(r)
}

func main() {
	config := GetConfig()

	r := gin.Default()
	registerRoutes(r)

	log.Printf("Starting %s on port: %s\n", SERVER_NAME, config.server.Port)
	err := r.Run(":" + config.server.Port)
	if err != nil {
		log.Println(err)
	}
}
