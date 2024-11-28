package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"shop/helper"
)

var DB *gorm.DB

func init() {
	// Loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func InitDatabase() {
	dsn := helper.GetDsn()

	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := gorm.Open("postgres", dsn)

	fmt.Println(dsn)
	fmt.Println(db)

	if err != nil {
		log.Fatal("Error", err.Error())
	} else {
		DB = db
		//err := db.AutoMigrate(&Product{}, &Verses{})
		//err := db.AutoMigrate(&Product{})
		db.AutoMigrate(&Category{}, &Product{}).AddForeignKey("category_id", "categories (id)", "CASCADE", "CASCADE")
		// .AddForeignKey добавляет внешний (FOREIGN KEY) category_id из product к таблице categories по id
		//if err != nil {
		//	log.Fatal("Automigrate error", err)
		//}
	}
}
