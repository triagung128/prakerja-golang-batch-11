package routes

import (
	authcontroller "day-7/controllers/auth"
	usercontroller "day-7/controllers/user"
	middlewareClaims "day-7/middleware"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	e.POST("/register", authcontroller.RegisterController)

	eAuth := e.Group("")
	// eAuth.Use(echojwt.JWT([]byte("123")))
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(middlewareClaims.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_JWT")),
	}
	eAuth.Use(echojwt.WithConfig(config))
	eAuth.POST("/login", authcontroller.LoginController)
	eAuth.GET("/users", usercontroller.GetUsersController)
}
