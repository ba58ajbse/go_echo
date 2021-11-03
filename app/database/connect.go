package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// var DB *gorm.DB

// func Connect() *gorm.DB {
func Connect() *sql.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
	}

	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	db_host := os.Getenv("DB_HOST")

	dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ":3306)/" + db_name

	// DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err.Error())
	}

	return DB
}
