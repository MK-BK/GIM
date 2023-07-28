package api

import (
	"net/http"

	"GIM/models"

	"github.com/gin-gonic/gin"
)

func userLogin(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	token, err := GE.Users.Login(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.Header("token", token)
	c.JSON(http.StatusOK, user)
}

func userLogout(c *gin.Context) {
	userID := c.GetString("userID")

	if err := GE.Users.Logout(userID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
}

func register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	if err := GE.Users.CreateUser(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func listUsers(c *gin.Context) {
	var options models.UserListOptions

	if err := c.BindJSON(&options); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	users, err := GE.Users.ListUsers(&options)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	var id string
	if err := SetPathVar(c, &id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	user, err := GE.Users.GetUser(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func updateUser(c *gin.Context) {
	var id string
	if err := SetPathVar(c, &id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	var patch models.UserPatch
	if err := c.BindJSON(&patch); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	if err := GE.Users.UpdateUser(id, &patch); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
}
