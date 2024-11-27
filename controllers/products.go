package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/swaggo/http-swagger"
	"io"
	"log"
	"net/http"
	"shop/helper"
	"shop/models"
	"strconv"
)

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

// GetProducts godoc
// @Summary Retrieves products
// @Produce json
// @Param page query string false "paginating results - ?page=1"
// @Param title query string  false "product search - ?title=Some title"
// @Param description query string  false "product search - ?description=Some descr"
// @Success 200 {object} models.Product
// @Failure 500
// @Router /products [get]
func GetProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Receiving products data")
	w.Header().Set("Content-Type", "application/json")
	var products []models.Product
	//
	//db := model.DB.Preload("Verses")
	db := models.DB
	//
	// Check get parameters for filtration
	query := r.URL.Query()
	if len(query) != 0 {
		for key, value := range query {
			if key != "page" && key != "items_per_page" && key != "release_date" {
				//db.Where(key+" LIKE ?", "%"+value[0]+"%")
				db = db.Where(key+"=?", value[0]).Find(&products)
				log.Printf("Using filter %s, value=%s", key, value[0])
			}
			if key == "release_date" {
				db = db.Where("CAST(release_date AS date)=?", value[0])
			}
		}
	}

	// Exclude song with "deleted_at"
	//models.DB.Where("deleted_at IS NULL").Find(&products)

	//models.DB.Where("deleted_at IS NOT NULL").Find(&products)
	//models.DB.Where()

	// Add a pagination
	//models.DB.Scopes(helper.Paginate(r)).Find(&products)
	//models.DB.Scopes(helper.Paginate(r)).Find(&products)

	db.Scopes(helper.Paginate(r)).Find(&products)
	//fmt.Println(models.DB.Where())

	//models.DB.Debug()

	//models.DB.Find(&products)

	res, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(res)
	}
}

// CreateProduct godoc
// @Summary Create a product
// @Produce json
// @Accept json
// @Param body body models.Product true "body"
// @Success 201
// @Failure 500
// @Router /product [POST]
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating product...")
	var product models.Product
	body, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(body, &product)
	if err != nil {
		log.Println("Error on data: ", err.Error())
	}

	models.DB.Create(&product)

	w.Header().Set("Content-Type", "application/json")

	// Send song's data in response
	res, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		log.Println("Product was created:")
		log.Println(product)
		_, err = w.Write(res)
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println(params["id"])
	var product models.Product
	// Receiving song's id
	id, ok := params["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
	}

	db := models.DB.Where("id=?", id).Find(&product)
	//fmt.Println(db)
	if db.RecordNotFound() {
		w.WriteHeader(http.StatusNotFound)
	} else {
		res := models.DB.Delete(&product, id)
		log.Println(res.Error)
		w.WriteHeader(http.StatusOK)
		log.Printf("Product with id=%s was deleted", id)
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println(params["id"])
	var productData models.Product
	var product models.Product
	// Receiving product's id
	id, ok := params["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
	}

	// Находим товар
	db := models.DB.Where("id=?", id).Find(&product)

	if db.RecordNotFound() {
		w.WriteHeader(http.StatusNotFound)
	} else {
		//log.Println(product.Sku)
		log.Printf("Updating product %s...", product.Sku)
		body, _ := io.ReadAll(r.Body)
		// Заполним товар для обновления данными, которые пришли от клиента
		err := json.Unmarshal(body, &productData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			// Если не было ошибки обновим товар
			models.DB.Model(&product).Updates(productData)
			w.WriteHeader(http.StatusOK)
		}
	}
}
