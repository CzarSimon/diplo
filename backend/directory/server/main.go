package main

import (
	"log"
	"net/http"

	"github.com/CzarSimon/diplo/backend/pkg/httputil"
	"github.com/gin-gonic/gin"
)

// registerRoutes sets up available routes for the service.
func registerRoutes(env *Env) *http.Server {
	r := gin.Default()
	httputil.RegisterHealthCheck(r)

	return &http.Server{
		Addr:    ":" + env.config.server.Port,
		Handler: r,
	}
}

func main() {
	config := GetConfig()
	env := setupEnv(config)
	defer env.db.Close()

	server := registerRoutes(env)
	log.Printf("Starting %s on port: %s\n", SERVER_NAME, config.server.Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
