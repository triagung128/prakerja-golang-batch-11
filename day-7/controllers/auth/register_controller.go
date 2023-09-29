package auth

import (
	"day-7/configs"
	"day-7/middleware"
	"day-7/models/base"
	usermodel "day-7/models/user"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(c echo.Context) error {
	var user usermodel.User
	c.Bind(&user)

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(hashPassword)

	result := configs.DB.Create(&user)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), fmt.Sprintf("Duplicate entry '%s' for key 'idx_users_email", user.Email)) {
			return c.JSON(http.StatusConflict, base.BaseResponse{
				Status:  false,
				Message: "User with that email already exists",
				Data:    nil,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, base.BaseResponse{
				Status:  false,
				Message: "Failed add data to database",
				Data:    nil,
			})
		}
	}

	authResponse := usermodel.ResponseAuth{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Token: middleware.GenerateTokenJWT(user.Id, user.Name),
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success Register",
		Data:    authResponse,
	})
}
