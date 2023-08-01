package api

import (
	"net/http"

	"GIM/pkg/event"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func handlerEvent(c *gin.Context) {
	userID := c.GetString("userID")

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ch := event.Sub()
	for {
		select {
		case evt := <-ch:
			if e, ok := evt.(event.ScopeEvent); ok {
				if e.GetUserID() == userID {
					ws.WriteJSON(map[string]interface{}{
						"Topic": evt.Type(),
						"Body":  evt,
					})
				}
			} else {
				ws.WriteJSON(map[string]interface{}{
					"Topic": evt.Type(),
					"Body":  evt,
				})
			}
		}
	}
}
