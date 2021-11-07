package tests

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetDSN() string {
	db_user := "root"
	db_pass := os.Getenv("TEST_DB_ROOT_PASS")
	db_name := os.Getenv("TEST_DB_NAME")
	db_host := os.Getenv("TEST_DB_HOST")

	dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ":3306)/" + db_name

	return dsn
}

func Connect() *sql.DB {
	dsn := GetDSN()
	db, _ := sql.Open("mysql", dsn)

	return db
}

func TestGetRegisteredDriver(t *testing.T) {
	assert.Equal(t, []string{"mysql"}, sql.Drivers())
}

func TestPingMySql(t *testing.T) {
	err := godotenv.Load("../../.env")

	assert.Nil(t, err)

	dsn := GetDSN()
	db, err := sql.Open("mysql", dsn)

	assert.Nil(t, err)
	assert.NotNil(t, db)

	err = db.Ping()
	assert.Nil(t, err)

	defer db.Close()
}

func TestGetUsers(t *testing.T) {
	db, err := sql.Open("mysql", GetDSN())
	defer db.Close()

	assert.Nil(t, err)
	assert.NotNil(t, db)

	rows, err := db.Query("SELECT * FROM users")
	assert.Nil(t, err)
	assert.NotNil(t, rows)

	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		assert.Nil(t, err)
		assert.NotNil(t, user)

		users = append(users, user)
	}

	assert.NotEmpty(t, users)
}

func TestGetUser(t *testing.T) {
	db, err := sql.Open("mysql", GetDSN())
	defer db.Close()

	assert.Nil(t, err)
	assert.NotNil(t, db)

	var user User

	id := 1
	err = db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.Id, 1)
	assert.Equal(t, user.Name, "John")
	assert.Equal(t, user.Email, "aaa@mail.com")

	id = 5
	err = db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.Email)
	assert.NotNil(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
}
