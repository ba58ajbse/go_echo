package main

import (
	"database/sql"
	"echo_app/app/database"
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
	users := []User{}
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, users)
}

func main() {
	e := echo.New()
	db = database.Connect()
	defer db.Close()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo World!")
	})
	e.GET("/hello", hello)
	e.GET("/users", getUsers)
	e.Logger.Fatal(e.Start(":8080"))
}
