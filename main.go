package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	populateProducts()
	router := httprouter.New()

	router.GET("/products", GetProducts)
	router.GET("/products/:id", GetProduct)
	router.POST("/products", CreateProduct)
	router.DELETE("/products/:id", DeleteProduct)

	log.Fatal(http.ListenAndServe(":8000", router))
}
