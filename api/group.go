package api

import (
	"net/http"

	"GIM/models"

	"github.com/gin-gonic/gin"
)

func listGroups(c *gin.Context) {
	userID := c.GetString("userID")

	groups, err := GE.Groups.ListGroups(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, groups)
}

func createGroup(c *gin.Context) {
	var options struct {
		Users []string
	}

	if err := c.BindJSON(&options); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	group := &models.Group{
		ManagerID: c.GetString("userID"),
	}

	group, err := GE.Groups.CreateGroup(options.Users, group)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, group)
}

func updateGroup(c *gin.Context) {
	var id string
	if err := SetPathVar(c, &id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	var patch models.GroupPatch
	if err := c.BindJSON(&patch); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := GE.Groups.UpdateGroup(id, &patch); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
}

func getGroup(c *gin.Context) {
	var id string
	if err := SetPathVar(c, &id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	group, err := GE.Groups.GetGroup(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, group)
}

func deleteGroup(c *gin.Context) {
	var id string
	if err := SetPathVar(c, &id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	userID := c.GetString("userID")
	if err := GE.Groups.DeleteGroup(userID, id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
}

func joinGroup(c *gin.Context) {
	userID := c.GetString("userID")

	var groupID string
	if err := SetPathVar(c, &groupID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	var selections struct {
		UserIDs []string
	}

	if err := c.BindJSON(&selections); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := GE.Groups.JoinGroup(userID, groupID, selections.UserIDs); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
}

func leaveGroup(c *gin.Context) {
	userID := c.GetString("userID")

	var groupID string
	if err := SetPathVar(c, &groupID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	var selections struct {
		UserID string
	}

	if err := c.BindJSON(&selections); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := GE.Groups.LeaveGroup(userID, groupID, selections.UserID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
}
