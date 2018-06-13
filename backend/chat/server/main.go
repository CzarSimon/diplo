package main

import (
	"log"
	"net/http"

	"github.com/CzarSimon/diplo/backend/pkg/httputil"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	config := GetConfig()
	env := setupEnv(config)

	server := registerRoutes(env)
	log.Printf("Starting %s on port: %s\n", SERVER_NAME, config.server.Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

// registerRoutes sets up available routes for the service.
func registerRoutes(env *Env) *http.Server {
	r := gin.Default()
	r.Use(httputil.AuthMiddleware(env.AuthOpts))
	registerChannelRoutes(r, env)
	registerMessageRoutes(r, env)
	registerUserRoutes(r, env)
	registerSocketRoutes(r, env)
	httputil.RegisterHealthCheck(r)

	return &http.Server{
		Addr:    ":" + env.config.server.Port,
		Handler: r,
	}
}
