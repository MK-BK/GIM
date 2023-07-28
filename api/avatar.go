package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func uploadAvatar(c *gin.Context) {
	userID := c.GetString("userID")

	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	path, err := getAvatarPath(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
}

func getAvatar(c *gin.Context) {
	var userID string
	if err := SetPathVar(c, &userID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	path, err := getAvatarPath(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	file, err := os.Open(path)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	io.Copy(c.Writer, file)
}

func deleteAvatar(c *gin.Context) {
	var userID string
	if err := SetPathVar(c, &userID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	path, err := getAvatarPath(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	if err := os.Remove(path); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
}

func getAvatarPath(userID string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(pwd, "avatar", fmt.Sprintf("%s.png", userID)), nil
}
