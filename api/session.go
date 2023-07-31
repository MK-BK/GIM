package api

import (
	"net/http"

	"GIM/models"

	"github.com/gin-gonic/gin"
)

func listSessions(c *gin.Context) {
	id := c.GetString("userID")

	sessions, err := GE.Sessions.List(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, sessions)
}

func createSession(c *gin.Context) {
	var body models.Session
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	body.OwnerID = c.GetString("userID")
	session, err := GE.Sessions.Create(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, session)
}

func getSession(c *gin.Context) {
	var id string
	if err := SetPathVar(c, &id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	session, err := GE.Sessions.Get(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, session)
}

func deleteSession(c *gin.Context) {
	var id string
	if err := SetPathVar(c, &id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := GE.Sessions.Delete(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
}
