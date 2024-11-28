package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/swaggo/http-swagger"
	"io"
	"log"
	"net/http"
	"shop/helper"
	"shop/models"
)

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

	db.Scopes(helper.Paginate(r)).Find(&products)

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

	fmt.Println(product)

	models.DB.Create(&product)

	w.Header().Set("Content-Type", "application/json")

	// Send product's data in response
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

// DeleteProduct godoc
// @Summary Delete a product by id
// @Param id path integer true "Product ID"
// @Success 200
// @Failure 400
// @Router /product/{id} [DELETE]
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println(params["id"])
	var product models.Product
	// Receiving product's id
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

// UpdateProduct godoc
// @Summary Updates product based on given ID
// @Produce json
// @Accept json
// @Param id path integer true "Product ID"
// @Param body body models.Product true "body"
// @Success 200
// @Failure 400
// @Failure 404
// @Router /product/{id} [PUT]
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

// GetProduct godoc
// @Summary Retrieves product based on given ID
// @Produce json
// @Param id path integer true "Product ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /product/{id} [GET]
func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println(params["id"])
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
		log.Printf("Product with id=%s was found", id)
		res, err := json.Marshal(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			_, err = w.Write(res)
		}
	}
}
