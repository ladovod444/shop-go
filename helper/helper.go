package helper

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetDsn() string {
	host, exists := os.LookupEnv("DB_HOST")
	if !exists {
		log.Fatal("Wrong DB_HOST")
	}
	port, exists := os.LookupEnv("DB_PORT")
	if !exists {
		log.Fatal("Wrong DB_PORT")
	}
	user, exists := os.LookupEnv("DB_USER")
	if !exists {
		log.Fatal("Wrong DB_USER")
	}
	pass, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		log.Fatal("Wrong DB_PASSWORD")
	}
	database, exists := os.LookupEnv("DATABASE_NAME")
	if !exists {
		log.Fatal("Wrong DATABASE_NAME")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, pass, database)

	return dsn
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		pageNum, err := strconv.Atoi(q.Get("page"))
		if err == nil {
			pageSize, _ := strconv.Atoi(q.Get("items_per_page"))
			switch {
			case pageSize > 100:
				pageSize = 100
			case pageSize <= 0:
				pageSize = 5
			}

			offset := (pageNum - 1) * pageSize
			return db.Offset(offset).Limit(pageSize)
		}

		return db
	}
}
