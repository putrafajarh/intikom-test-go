package utils

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var ErrInvalidRefreshToken = errors.New("invalid refresh token")
var ErrInvalidAccessToken = errors.New("invalid access token")

func getSecret(tokenType string) string {
	if tokenType == "access" {
		return viper.GetString("JWT_ACCESS_SECRET")
	}
	if tokenType == "refresh" {
		return viper.GetString("JWT_REFRESH_SECRET")
	}
	panic("invalid token type")
}

// GenerateAccessToken generates a new access token for a user with a 1 hour expiration
func GenerateAccessToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
	})

	t, err := token.SignedString([]byte(getSecret("access")))
	if err != nil {
		return "", err
	}

	return t, nil
}

// Return claims if token is valid, otherwise return false
func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(getSecret("access")), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, ErrInvalidAccessToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidAccessToken
	}

	return claims, nil
}

// GenerateRefreshToken generates a new refresh token for a user with a 7 day expiration
func GenerateRefreshToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
	})

	t, err := token.SignedString([]byte(getSecret("refresh")))
	if err != nil {
		return "", err
	}

	return t, nil
}

func GenerateAccessTokenFromRefreshToken(tokenString string) (string, error) {
	refreshToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
		return []byte(getSecret("refresh")), nil
	})
	if err != nil || !refreshToken.Valid {
		return "", ErrInvalidRefreshToken
	}

	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok || !refreshToken.Valid {
		return "", ErrInvalidRefreshToken
	}

	return GenerateAccessToken(claims["sub"].(string))
}

func ParseRefreshToken(tokenString string) (bool, error) {
	refreshToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
		return []byte(getSecret("refresh")), nil
	})
	if err != nil || !refreshToken.Valid {
		return false, ErrInvalidRefreshToken
	}

	return refreshToken.Valid, nil
}
