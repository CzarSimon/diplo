package main

import (
	"github.com/CzarSimon/diplo/backend/chat/pkg/chat"
	"github.com/CzarSimon/diplo/backend/pkg/httputil"
	"github.com/CzarSimon/diplo/backend/pkg/id"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// registerSocketRoutes registers routes handles websocket connections.
func registerSocketRoutes(r *gin.Engine, env *Env) {
	r.GET("/connect/:userId", env.handleConnectionRequest)
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// handleConnectionRequest handles creation and managing of websockets.
func (env *Env) handleConnectionRequest(c *gin.Context) {
	userID := c.Param("userId")
	err := env.checkUserExists(userID)
	if err != nil {
		httputil.JSONError(c, err)
		return
	}

	msgChan := make(chan chat.Message)
	sessionID := id.New()
	err = env.Channels.Add(userID, sessionID, msgChan)
	if err != nil {
		httputil.JSONError(c, err)
		return
	}

	ws, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		httputil.JSONError(c, err)
		return
	}
	defer ws.Close()

	for {
		msg := <-msgChan
		err = ws.WriteJSON(msg)
		if err != nil {
			env.Channels.Delete(userID, sessionID)
			return
		}
	}
}
