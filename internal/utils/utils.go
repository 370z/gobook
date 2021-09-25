package utils

import (
	"gobook/internal/auth"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetJWTClaims(c echo.Context) *auth.JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	return user.Claims.(*auth.JwtCustomClaims)
}