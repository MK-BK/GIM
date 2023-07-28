package middleware

import (
	"errors"
	"net/http"
	"strings"

	"GIM/models"
	"GIM/pkg/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "/login") ||
			strings.Contains(c.Request.URL.Path, "/register") ||
			(strings.Contains(c.Request.URL.Path, "/avatar") && c.Request.Method == http.MethodGet) {
			c.Next()
			return
		}

		token := c.GetHeader("token")
		if strings.Contains(c.Request.URL.Path, "/event") {
			token = c.Query("token")
		}

		if token == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("http.StatusUnauthorized"))
			return
		}

		t, err := jwt.ParseWithClaims(token, &models.LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(models.Conf.SecretKey), nil
		})
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New("http.StatusUnauthorized"))
			return
		}

		if claims, ok := t.Claims.(*models.LoginClaims); ok && t.Valid {
			c.Set("userName", claims.UserName)
			c.Set("userID", claims.UserID)

			if err := config.GetRedisClient().Get(claims.UserID).Err(); err != nil {
				c.AbortWithError(http.StatusUnauthorized, errors.New("http.StatusUnauthorized"))
				return
			}
		} else {
			c.AbortWithError(http.StatusUnauthorized, errors.New("http.StatusUnauthorized"))
			return
		}

		c.Next()
	}
}
