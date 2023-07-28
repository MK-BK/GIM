package api

import (
	"net/http"

	"GIM/models"

	"github.com/gin-gonic/gin"
)

func createComment(c *gin.Context) {
	var comment models.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if comment.OwnerID == "" {
		comment.OwnerID = c.GetString("userID")
	}

	if err := GE.Comments.Create(&comment); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
}

func deleteComment(c *gin.Context) {
	var id string
	if err := SetPathVar(c, &id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := GE.Comments.Delete(id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
}
