package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"shop/helper"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type Product struct {
	gorm.Model
	Title       string `json:"Title" example:"Product title"`
	Sku         string `json:"Sku" example:"Product sku"`
	Description string `json:"Description" example:"Product description"`
	//CurrentPrice float32 `json:"CurrentPrice" example:"10.01" sql:"type:decimal(10,2);"`
	//RegularPrice float32 `json:"RegularPrice" example:"10.01" sql:"type:decimal(10,2);"`
	//CreatedAt    time.Time `json:"CreatedAt" example:"2006-02-01T15:04:05Z" gorm:"default:current_timestamp"`
	//UpdatedAt    time.Time `json:"UpdatedAt" example:"2006-02-01T15:04:05Z" gorm:"default:current_timestamp"`
	CurrentPrice float32 `json:"CurrentPrice" example:"10.01"`
	RegularPrice float32 `json:"RegularPrice" example:"10.01"`
	//CreatedAt    time.Time `gorm:"autoCreateTime:true" json:"createdAt"`
	//UpdatedAt    time.Time `gorm:"autoUpdateTime:true" json:"updatedAt"`
	Image string `json:"Image" example:"Product Image"`
	//Verses      []Verses  `json:"Verses" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

//CurrentPrice float32   `json:"CurrentPrice" example:"10.01" sql:"type:decimal(10,2);"`
//RegularPrice float32   `json:"RegularPrice" example:"10.01" sql:"type:decimal(10,2);"`

//type Verses struct {
//	gorm.Model
//	Text   string `json:"Text" example:"Verse text"`
//	ProductID uint   `json:"ProductID" swaggerignore:"true"`
//}

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
		db.AutoMigrate(&Product{})
		//if err != nil {
		//	log.Fatal("Automigrate error", err)
		//}
	}
}
