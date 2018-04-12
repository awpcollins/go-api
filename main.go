package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	SKU        string `json:"sku,omitempty"`
	Advertiser string `json:"advertiser,omitempty"`
}

var products []Product

func main() {
	populateProducts()
	router := mux.NewRouter()

	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/products", CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, product := range products {
		if product.ID == params["id"] {
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	json.NewDecoder(r.Body).Decode(&product)

	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, product := range products {
		if product.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			break
		}

	}
	fmt.Printf("%+v\n", products)
	json.NewEncoder(w).Encode(products)
}

func populateProducts() {
	products = append(products, Product{ID: "123", Name: "soap", SKU: "HU12838", Advertiser: "Walmart"})
	products = append(products, Product{ID: "124", Name: "brush", SKU: "HU12836", Advertiser: "Walmart"})
	products = append(products, Product{ID: "125", Name: "rope", SKU: "HU12839", Advertiser: "Walmart"})
}
