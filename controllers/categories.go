package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"shop/models"
)

// CreateCategory godoc
// @Summary Create a category
// @Produce json
// @Accept json
// @Param body body models.Category true "body"
// @Success 201
// @Failure 500
// @Router /category [POST]
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating category...")
	var category models.Category
	body, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(body, &category)
	if err != nil {
		log.Println("Error on data: ", err.Error())
	}

	models.DB.Create(&category)

	w.Header().Set("Content-Type", "application/json")

	// Send product's data in response
	res, err := json.Marshal(category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		log.Println("Category was created:")
		log.Println(category)
		_, err = w.Write(res)
	}
}

// DeleteCategory godoc
// @Summary Delete a category by id
// @Param id path integer true "Category ID"
// @Success 200
// @Failure 400
// @Router /category/{id} [DELETE]
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println(params["id"])
	var category models.Category
	// Receiving category's id
	id, ok := params["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
	}

	db := models.DB.Where("id=?", id).Find(&category)
	//fmt.Println(db)
	if db.RecordNotFound() {
		w.WriteHeader(http.StatusNotFound)
	} else {
		res := models.DB.Delete(&category, id)
		log.Println(res.Error)
		w.WriteHeader(http.StatusOK)
		log.Printf("Category with id=%s was deleted", id)
	}
}
