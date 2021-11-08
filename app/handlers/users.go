package handlers

import (
	"echo_app/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUsers(c echo.Context) error {
	users, err := model.Find(&model.User{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, "record not found")
	}

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := model.Select(&model.User{}, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "record not found")
	}

	return c.JSON(http.StatusOK, user)
}
