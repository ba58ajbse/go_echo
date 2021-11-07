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

func getUser(c echo.Context) error {
	var user User

	id := c.Param("id")
	err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.Email)

	switch {
	case err == sql.ErrNoRows:
		return c.JSON(http.StatusOK, "no record.")
	case err != nil:
		panic(err.Error())
	default:
		return c.JSON(http.StatusOK, user)
	}
}

func main() {
	e := echo.New()
	db = database.Connect()
	defer db.Close()

	e.GET("/", hello)
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUser)
	e.Logger.Fatal(e.Start(":8080"))
}
