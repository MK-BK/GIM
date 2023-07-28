package api

import (
	"net/http"

	"GIM/models"

	"github.com/gin-gonic/gin"
)

func listFriends(c *gin.Context) {
	userID := c.GetString("userID")

	users, err := GE.Users.ListFriends(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func addFriend(c *gin.Context) {
	userID := c.GetString("userID")

	var option models.UserAddOption
	if err := c.BindJSON(&option); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := GE.Users.AddFriend(userID, &option); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
}

func deleteFriend(c *gin.Context) {
	userID := c.GetString("userID")

	var id string
	if err := SetPathVar(c, &id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := GE.Users.DeleteFriend(userID, id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
}
