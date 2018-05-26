package main

import (
	"net/http"

	"github.com/CzarSimon/diplo/backend/chat/pkg/chat"
	"github.com/CzarSimon/diplo/backend/pkg/httputil"
	"github.com/gin-gonic/gin"
)

// registerChannelRoutes registers routes that handles chat messages.
func registerChannelRoutes(r *gin.Engine, env *Env) {
	r.POST("/channel/:gameID/:name", env.handleNewChannel)
	r.PUT("/channel/:channelID/:userID", env.handleNewChannelUser)
	r.DELETE("/channel/:channelID/:userID", env.handleChannelUserRemoval)
}

// handleNewChannel handles a request of creating a new channel.
func (env *Env) handleNewChannel(c *gin.Context) {
	name := c.Param("name")
	gameID := c.Param("gameID")
	recievers, err := httputil.ParseQueryValues(c, "recieverId")
	if err != nil {
		httputil.JSONError(c, err)
		return
	}
	channel := chat.NewChannel(name, gameID, recievers)
	err = env.saveAndRegisterChannel(channel)
	if err != nil {
		httputil.JSONError(c, err)
		return
	}
	c.JSON(http.StatusOK, channel)
}

// handleNewChannelUser handles requests for adding a user to a channel.
func (env *Env) handleNewChannelUser(c *gin.Context) {
	userID := c.Param("userID")
	channelID := c.Param("channelID")
	err := env.addUserToChannel(userID, channelID)
	if err != nil {
		httputil.JSONError(c, err)
		return
	}
	httputil.SendOK(c)
}

// handleChannelUserRemoval handles requests for removing a user to a channel.
func (env *Env) handleChannelUserRemoval(c *gin.Context) {
	userID := c.Param("userID")
	channelID := c.Param("channelID")
	err := env.removeUserFromChannel(userID, channelID)
	if err != nil {
		httputil.JSONError(c, err)
		return
	}
	httputil.SendOK(c)
}
