package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.Start(":8000")
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetUsersController(c echo.Context) error {
	var users = []User{
		{
			Id:   1,
			Name: "Iqbal",
		},
		{
			Id:   2,
			Name: "Dedi",
		},
		{
			Id:   3,
			Name: "Dian",
		},
	}

	return c.JSON(http.StatusOK, BaseResponse{
		Status:  true,
		Message: "success",
		Data:    users,
	})
}
