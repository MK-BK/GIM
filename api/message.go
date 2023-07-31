package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

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

	message.Scope = models.Scope{
		Kind:          session.Kind,
		OwnerID:       userID,
		DestinationID: session.DestinationID,
		GroupID:       session.GroupID,
	}

	if err := service.DatacenterAPI.HandlerMessage(sessionID, &message); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
}

func getMessage(c *gin.Context) {
	var sessionID, fileID string
	if err := SetPathVar(c, &sessionID, &fileID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("+++++++++++++++ var:", sessionID, fileID)

	pwd, err := os.Getwd()
	if err != nil {
		return
	}

	path := filepath.Join(pwd, "sessions", sessionID, "images", fileID)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("++++++++++++++++++ path:", path)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	io.Copy(c.Writer, file)
}
