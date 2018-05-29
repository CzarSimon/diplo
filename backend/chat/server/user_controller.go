package main

import (
	"net/http"

	"github.com/CzarSimon/diplo/backend/pkg/httputil"
	"github.com/gin-gonic/gin"
)

// Common errors related to users.
var (
	ErrNoSuchUser = httputil.NewError(http.StatusNotFound, "No such user")
)

// registerUserRoutes registers routes that handles handles users.
func registerUserRoutes(r *gin.Engine, env *Env) {
	r.POST("/user/:userID", env.handleNewUser)
	r.DELETE("/user/:userID", env.handleUserDeletion)
}

// handleNewUser handles requests to add user..
func (env *Env) handleNewUser(c *gin.Context) {
	userID := c.Param("userID")
	err := env.saveUser(userID)
	if err != nil {
		httputil.JSONError(c, err)
		return
	}
	httputil.SendOK(c)
}

// handleUserDeletion handles requests to remove a user.
func (env *Env) handleUserDeletion(c *gin.Context) {
	userID := c.Param("userID")
	err := env.removeUser(userID)
	if err != nil {
		httputil.JSONError(c, err)
		return
	}
	httputil.SendOK(c)
}
