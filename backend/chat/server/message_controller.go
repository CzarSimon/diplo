package main

import (
	"net/http"

	"github.com/CzarSimon/diplo/backend/chat/pkg/chat"
	"github.com/CzarSimon/diplo/backend/pkg/httputil"
	"github.com/gin-gonic/gin"
)

// registerMessageRoutes registers routes that handles chat messages.
func registerMessageRoutes(r *gin.Engine) {
	r.POST("/message/:channelID", handleNewMessage)
}

// handleNewMessage handles a new incomming message.
func handleNewMessage(c *gin.Context) {
	channelID := c.Param("channelID")
	var message chat.Message
	err := c.BindJSON(&message)
	if err != nil {
		httputil.JSONError(c, http.StatusBadRequest, err)
		return
	}
	message = chat.NewMessage(message.Text, channelID, message.AuthorID)
	err = saveAndBroadcastMessage(message)
	if err != nil {
		httputil.JSONError(c, http.StatusInternalServerError, err)
		return
	}
	httputil.SendOK(c)
}
