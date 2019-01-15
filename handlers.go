package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetProducts(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response := &JsonResponse{Data: &products}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func GetProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	for _, product := range products {

		fmt.Printf("%+v\n", params.ByName("id"))
		fmt.Println("hello")
		if product.ID == params.ByName("id") {
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	response := JsonErrorResponse{Error: &ApiError{Status: 404, Title: "Record Not Found"}}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	for index, product := range products {
		if product.ID == params.ByName("id") {
			fmt.Printf("%+v\n", params.ByName("id"))

			products = append(products[:index], products[index+1:]...)
			json.NewEncoder(w).Encode(products)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	response := JsonErrorResponse{Error: &ApiError{Status: 404, Title: "Record Not Found"}}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var product Product
	json.NewDecoder(r.Body).Decode(&product)

	products = append(products, product)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	json.NewEncoder(w).Encode(product)
}
