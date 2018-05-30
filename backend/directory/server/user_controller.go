package main

import (
	"net/http"

	"github.com/CzarSimon/diplo/backend/directory/pkg/directory"
	"github.com/CzarSimon/diplo/backend/pkg/httputil"
	"github.com/gin-gonic/gin"
)

var ErrAccessDenied = httputil.NewError(http.StatusForbidden, "Access denied")

// registerUserRoutes registers routes that handles handles users.
func registerUserRoutes(r *gin.Engine, env *Env) {
	r.POST("/user", env.handleNewUser)
	r.POST("/user/login", env.handleUserLogin)
	// r.GET("/user/token/renew", env.handleNewUser)
}

// handleNewUser handles requests to add new users.
func (env *Env) handleNewUser(c *gin.Context) {
	var user directory.User
	err := c.BindJSON(&user)
	if err != nil {
		httputil.JSONError(c, httputil.ErrBadRequest)
		return
	}
	user, err = env.saveUser(user)
	if err != nil {
		httputil.JSONError(c, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// handleUserLogin handles login request for a user.
func (env *Env) handleUserLogin(c *gin.Context) {
	var user directory.User
	err := c.BindJSON(&user)
	if err != nil {
		httputil.JSONError(c, httputil.ErrBadRequest)
		return
	}
	token, err := env.loginUser(user)
	if err != nil {
		httputil.JSONError(c, err)
		return
	}
	c.JSON(http.StatusOK, token)
}
