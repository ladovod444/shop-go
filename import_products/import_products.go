package import_products

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	gormbulk "github.com/t-tiger/gorm-bulk-insert/v2"
	"io"
	"log"
	"net/http"
	"os"
	"shop/helper"
	"shop/models"
	"strconv"
	"strings"
)

func ProductsImport() {

	// Минимальный вариант для взаимодействия
	fmt.Println("Enter an count of products you want to import, 'c' to close to app")
	reader := bufio.NewReader(os.Stdin)

	countStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	countStr = strings.TrimSpace(countStr)

	count, err := strconv.Atoi(countStr)
	if err != nil {
		log.Fatal(err)
	}
	processImport(count)

}

func processImport(count int) {
	fmt.Printf("Importing products %d....\n", count)

	//var insertRecords []interface{}
	var insertRecords []interface{}

	// Получим строку подключения к БД
	dsn := helper.GetDsn()

	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := gorm.Open("postgres", dsn)

	api_url, exists := os.LookupEnv("FORTNITE_API_URL")
	if !exists {
		log.Fatal("Wrong API_URL")
	}
	api_key, exists := os.LookupEnv("FORTNITE_API_KEY")
	if !exists {
		log.Fatal("Wrong API_KEY")
	}

	res, err := request(api_url, api_key)
	if err != nil {
		log.Fatal(err)
	}

	// Для обработки json необходимо подготовить соответствующие структуры

	// data["price"]
	type Price struct {
		FinalPrice   int `json:"finalPrice"`
		RegularPrice int `json:"regularPrice"`
	}

	// data["displayAssets"] для получения урла изображения
	type DisplayAssets struct {
		Url string `json:"url"`
	}

	// И сам product c использованием вышеуказанных структур
	type Product struct {
		Sku           string          `json:"mainId"`
		Title         string          `json:"displayName"`
		Price         Price           `json:"price"`
		Description   string          `json:"displayDescription"`
		DisplayAssets []DisplayAssets `json:"displayAssets"`
	}

	type Shop struct {
		//Products []models.Product `json:"users"`
		Products []Product `json:"shop"`
	}

	var shop Shop

	err = json.Unmarshal(res, &shop)
	//fmt.Println(shop)
	pcount := 0
	for _, product := range shop.Products {
		pcount++
		//fmt.Println(product)
		fmt.Println(pcount, product.Sku, product.Title, product.Price.RegularPrice, product.DisplayAssets[0].Url, product.Description)

		insertRecords = append(insertRecords,
			models.Product{
				Title:        product.Title,
				Sku:          product.Sku,
				CurrentPrice: float32(product.Price.RegularPrice),
				RegularPrice: float32(product.Price.FinalPrice),
				//CreatedAt:    time.Time{},
				//UpdatedAt:    time.Time{},
				Image:       product.DisplayAssets[0].Url,
				Description: product.Description,
			},
		)

		if pcount == count {
			break
		}
	}
	err = gormbulk.BulkInsert(db, insertRecords, 3000)
	if err != nil {
		// do something
		log.Fatal(err)
	}

}

// func request(url, api_key string) (string, error) {
func request(url, api_key string) ([]byte, error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle error
	}

	req.Header.Add("Authorization", api_key)
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	result, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	return result, nil
}
