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
	authOpts := setupAuthOptions(config)

	server := registerRoutes(env, authOpts)
	log.Printf("Starting %s on port: %s\n", SERVER_NAME, config.server.Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

// registerRoutes sets up available routes for the service.
func registerRoutes(env *Env, authOpts *httputil.AuthOptions) *http.Server {
	r := gin.Default()
	r.Use(httputil.AuthMiddleware(authOpts))
	registerChannelRoutes(r, env)
	registerMessageRoutes(r, env)
	registerUserRoutes(r, env)
	httputil.RegisterHealthCheck(r)

	return &http.Server{
		Addr:    ":" + env.config.server.Port,
		Handler: r,
	}
}

// setupAuthOptions sets up authentication options for the service.
func setupAuthOptions(config Config) *httputil.AuthOptions {
	return httputil.NewAuthOptions(config.jwtSecret, "directory", config.authExemptedRoutes...)
}
