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
	r.GET("/user/token/renew", env.handleTokenRenewal)
	r.GET("/me", env.handleGetUser)
	r.GET("/users", env.handleGetUsers)
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

	token, err := env.createJWTToken(user)
	if err != nil {
		httputil.JSONError(c, err)
		return
	}
	c.JSON(http.StatusOK, directory.NewSignupResponse(sanitizeUser(user), token))
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

// handleTokenRenewal handles request to renew a user token if it is not expired.
func (env *Env) handleTokenRenewal(c *gin.Context) {
	tokenString, err := httputil.GetTokenString(c)
	if err != nil {
		httputil.JSONError(c, httputil.ErrBadRequest)
		return
	}

	newToken, err := env.renewUserToken(tokenString)
	if err != nil {
		httputil.JSONError(c, err)
		return
	}
	c.JSON(http.StatusOK, newToken)
}

// handleGetUsers handles requests to get user info for multiple users.
func (env *Env) handleGetUsers(c *gin.Context) {
	userIDs, err := httputil.ParseQueryValues(c, "userId")
	if err != nil {
		httputil.JSONError(c, err)
		return
	}

	users, err := env.getUsersByID(userIDs)
	if err != nil {
		httputil.JSONError(c, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

// handleGetUser handles requests to get full user info.
func (env *Env) handleGetUser(c *gin.Context) {
	userID := httputil.GetUserID(c)
	user, err := env.getUserByID(userID)
	if err != nil {
		httputil.JSONError(c, err)
		return
	}
	c.JSON(http.StatusOK, sanitizeUser(user))
}
