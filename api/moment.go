package api

import (
	"net/http"

	"GIM/models"

	"github.com/gin-gonic/gin"
)

func listMoments(c *gin.Context) {
	userID := c.GetString("userID")

	moments, err := GE.Moments.List(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, moments)
}

func createMoment(c *gin.Context) {
	var moment models.Moment
	if err := c.BindJSON(&moment); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if moment.OwnerID == "" {
		moment.OwnerID = c.GetString("userID")
	}

	if err := GE.Moments.Create(&moment); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
}

func getMoment(c *gin.Context) {
	var id string
	if err := SetPathVar(c, &id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	moment, err := GE.Moments.Get(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, moment)
}

func deleteMoment(c *gin.Context) {
	var id string
	if err := SetPathVar(c, &id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := GE.Moments.Delete(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
}
