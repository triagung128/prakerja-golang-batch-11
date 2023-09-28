package routes

import (
	usercontroller "day-7/controllers/user"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/users", usercontroller.GetUsersController)
}
