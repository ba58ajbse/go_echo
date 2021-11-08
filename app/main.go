package main

import (
	"database/sql"
	"echo_app/app/handlers"
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

func main() {
	e := echo.New()

	model.DBInit()
	defer model.DBClose()

	e.GET("/", hello)
	e.GET("/users", handlers.GetUsers)
	e.GET("/users/:id", handlers.GetUser)

	e.Logger.Fatal(e.Start(":8080"))
}
