package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func listRequests(c *gin.Context) {
	userID := c.GetString("userID")

	requests, err := GE.Requests.ListRequests(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, requests)
}

func getRequest(c *gin.Context) {
	var id string
	if err := SetPathVar(c, &id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	request, err := GE.Requests.GetRequest(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, request)
}

func updateRequest(c *gin.Context) {
	userID := c.GetString("userID")

	var id string
	if err := SetPathVar(c, &id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := GE.Requests.UpdateRequest(userID, id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
}

func deleteRequest(c *gin.Context) {
	var id string
	if err := SetPathVar(c, &id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := GE.Requests.DeleteRequest(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
}
