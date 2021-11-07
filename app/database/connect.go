package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func Connect() *sql.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
	}

	db_user := "root"
	db_pass := os.Getenv("DB_ROOT_PASS")
	db_name := os.Getenv("DB_NAME")
	db_host := os.Getenv("DB_HOST")

	dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ":3306)/" + db_name

	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}

	return DB
}
