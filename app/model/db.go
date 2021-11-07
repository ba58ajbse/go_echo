package model

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	sqlDB *sql.DB
	err   error
)

func DBInit() {
	err = godotenv.Load("../.env")
	if err != nil {
		log.Fatalln(err.Error())
	}

	db_user := "root"
	db_pass := os.Getenv("DB_ROOT_PASS")
	db_name := os.Getenv("DB_NAME")
	db_host := os.Getenv("DB_HOST")

	dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ":3306)/" + db_name

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	sqlDB, err = db.DB()
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalln("ping db error: ", err.Error())
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)

	db.AutoMigrate(&User{})
}

func DBClose() {
	sqlDB.Close()
}
