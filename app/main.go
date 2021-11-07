package main

import (
	"database/sql"
	"echo_app/app/model"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

var db *sql.DB

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUsers(c echo.Context) error {
	users := model.Find(&model.User{})

	return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	user := model.Select(&model.User{}, id)

	return c.JSON(http.StatusOK, user)
}

func main() {
	e := echo.New()

	model.DBInit()
	defer model.DBClose()

	e.GET("/", hello)
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUser)
	e.Logger.Fatal(e.Start(":8080"))
}
