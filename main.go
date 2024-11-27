package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"shop/controllers"
	_ "shop/docs"
	"shop/import_products"
	"shop/models"
	"strings"
)

// @title Products API
// @version 1.0
// @description Swagger API for Songs library API.
// @termsOfService http://swagger.io/terms/

// @contact.name Dmitrii
// @contact.email ladovod@gmail.com

// @BasePath /api/v1
func main() {

	models.InitDatabase()
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/v1/product", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/v1/product/{id}", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/api/v1/product/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/v1/product/{id}", controllers.DeleteProduct).Methods("DELETE")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Минимальный вариант для взаимодействия
	fmt.Println("Enter an option 'i'- for import, 's'- to start server, 'c' to close to app")
	reader := bufio.NewReader(os.Stdin)

	option, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	option = strings.TrimSpace(option)
	switch option {
	// В случае есл пользователь ввел i произведем импорт товаров
	case "i":
		import_products.ProductsImport()
	case "c":
		log.Fatal("Closing app...")
	case "s":
		log.Println("Starting server")
		//log.Fatal(http.ListenAndServe(":7070", router))
		log.Fatal(http.ListenAndServe(":7080", router))
	}

	//log.Println(option)
}
