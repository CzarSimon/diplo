package main

import (
	"log"
	"net/http"
	"time"

	"github.com/CzarSimon/diplo/backend/chat/pkg/chat"
	"github.com/CzarSimon/diplo/backend/pkg/httputil"
	"github.com/gin-gonic/gin"
)

// registerMessageRoutes registers routes that handles chat messages.
func registerMessageRoutes(r *gin.Engine, env *Env) {
	r.POST("/message/:channelID", env.handleNewMessage)
	r.GET("/message/:channelID", env.handleGetMessages)
}

// handleNewMessage handles a new incomming message.
func (env *Env) handleNewMessage(c *gin.Context) {
	userID := httputil.GetUserID(c)
	channelID := c.Param("channelID")
	var message chat.Message
	err := c.BindJSON(&message)
	if err != nil {
		httputil.JSONError(c, httputil.ErrBadRequest)
		return
	}
	message = chat.NewMessage(message.Text, channelID, userID)
	err = env.saveAndBroadcastMessage(message)
	if err != nil {
		httputil.JSONError(c, err)
		return
	}
	httputil.SendOK(c)
}

// handleGetMessages handles retrival of messages for a channel.
func (env *Env) handleGetMessages(c *gin.Context) {
	channelID := c.Param("channelID")
	since, err := httputil.ParseTimestampQuery(c, "since")
	if err != nil {
		log.Println(err)
		since = time.Time{}
	}
	messages, err := env.getMessagesSince(channelID, since)
	if err != nil {
		httputil.JSONError(c, err)
		return
	}
	c.JSON(http.StatusOK, messages)
}
