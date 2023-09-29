package auth

import (
	"day-7/middleware"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(middleware.JwtCustomClaims)
	// name := claims.Name
	// return c.String(http.StatusOK, "Welcome "+name+"!")

	if temp := c.Get("user"); temp != nil {
		u := temp.(*jwt.Token)
		claims := u.Claims.(*middleware.JwtCustomClaims)
		name := strconv.Itoa(int(claims.UserId))
		return c.String(http.StatusOK, "Welcome "+name+"!")
	}

	return c.String(http.StatusInternalServerError, "Error")
}
