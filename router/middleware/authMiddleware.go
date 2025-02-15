package middleware

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"intikom-test-go/repository"
	"intikom-test-go/utils"

	"github.com/gin-gonic/gin"
)

var (
	TokenHeaderPrefix = "Bearer"

	ErrEmptyAuthHeader   = errors.New("authorization header is required")
	ErrInvalidAuthHeader = errors.New("authorization header is invalid")
)

func jwtFromHeader(c *gin.Context) (string, error) {
	token := c.GetHeader("Authorization")
	if token == "" {
		return "", ErrEmptyAuthHeader
	}

	parts := strings.Split(token, " ")
	if len(parts) != 2 || parts[0] != TokenHeaderPrefix {
		return "", ErrInvalidAuthHeader
	}

	return parts[1], nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := jwtFromHeader(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, err := utils.VerifyAccessToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		userId, err := strconv.ParseUint(claims["sub"].(string), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		userRepository := repository.NewUserRepository()
		user, err := userRepository.FindById(uint(userId))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("user", user)
		c.Set("user_id", uint(userId))
		c.Next()
	}
}
