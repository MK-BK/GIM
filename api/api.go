package api

import (
	"fmt"

	"GIM/middleware"
	"GIM/models"

	"github.com/gin-gonic/gin"
)

var GE = &models.GlobalEnvironment

func InitRouter() error {
	engine := gin.Default()

	engine.Use(middleware.AuthMiddleware())

	engine.GET("/event", handlerEvent)

	engine.POST("/login", userLogin)
	engine.POST("/logout", userLogout)
	engine.POST("/register", register)

	engine.POST("/upload", upload)

	users := engine.Group("/users")
	{
		users.POST("", listUsers)
		users.GET("/:id", getUser)
		users.POST("/:id", updateUser)
	}

	friends := engine.Group("/friends")
	{
		friends.GET("", listFriends)
		friends.POST("", addFriend)
		friends.DELETE("/:id", deleteFriend)
	}

	requests := engine.Group("/requests")
	{
		requests.GET("", listRequests)
		requests.GET("/:id", getRequest)
		requests.POST("/:id", updateRequest)
		requests.DELETE("/:id", deleteRequest)
	}

	groups := engine.Group("/groups")
	{
		groups.GET("", listGroups)
		groups.POST("", createGroup)
		groups.POST("/:id", updateGroup)
		groups.GET("/:id", getGroup)
		groups.DELETE("/:id", deleteGroup)
		groups.POST("/:id/join", joinGroup)
		groups.POST("/:id/leave", leaveGroup)
	}

	sessions := engine.Group("/sessions")
	{
		sessions.GET("", listSessions)
		sessions.POST("", createSession)
		sessions.GET("/:id", getSession)
		sessions.DELETE("/:id", deleteSession)
		sessions.GET("/:id/messages", listMessages)
		sessions.POST("/:id/messages", sendMessage)
		sessions.GET("/:id/messages/files/:fileID", getMessage)
	}

	moments := engine.Group("/moments")
	{
		moments.GET("", listMoments)
		moments.POST("", createMoment)
		moments.GET("/:id", getMoment)
		moments.DELETE("/:id", deleteMoment)
	}

	comments := engine.Group("/comments")
	{
		comments.POST("", createComment)
		comments.DELETE("/:id", deleteComment)
	}

	avatars := engine.Group("/avatar")
	{
		avatars.POST("", uploadAvatar)
		avatars.GET("/:id", getAvatar)
		avatars.DELETE("/:id", deleteAvatar)
	}

	return engine.Run(":" + models.Conf.ListenPort)
}

func SetPathVar(c *gin.Context, values ...*string) error {
	if len(c.Params) != len(values) {
		return fmt.Errorf("path variables count mismatch, expected %d, actual %d", len(c.Params), len(values))
	}

	for i, param := range c.Params {
		*values[i] = param.Value
	}

	return nil
}
