package api

import (
	"net/http"

	"GIM/models"
	"GIM/service"

	"github.com/gin-gonic/gin"
)

func listMessages(c *gin.Context) {
	var sessionID string
	if err := SetPathVar(c, &sessionID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	messages, err := service.DatacenterAPI.GetMessages(sessionID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, messages)
}

func sendMessage(c *gin.Context) {
	userID := c.GetString("userID")

	var sessionID string
	if err := SetPathVar(c, &sessionID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	session, err := GE.Sessions.Get(sessionID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	var message models.Message
	if err := c.BindJSON(&message); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	message.Kind = session.Kind
	message.OwnerID = userID
	message.DestinationID = session.DestinationID
	message.GroupID = session.GroupID

	if err := service.DatacenterAPI.HandlerMessage(&message); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
}
