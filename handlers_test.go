package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestGetProducts(t *testing.T) {
	populateProducts()

	request, err := http.NewRequest("GET", "/products", nil)

	if err != nil {
		t.Fatal(err)
	}

	requestRecorder := newRequestRecorder(request, "GET", "/products", GetProducts)

	if requestRecorder.Code != 200 {
		t.Error("Expected response code to be 200")
	}

	expectedResponse := "{\"meta\":null,\"data\":[{\"id\":\"123\",\"name\":\"soap\",\"sku\":\"HU12838\",\"advertiser\":\"Walmart\"},{\"id\":\"124\",\"name\":\"brush\",\"sku\":\"HU12836\",\"advertiser\":\"Walmart\"},{\"id\":\"125\",\"name\":\"rope\",\"sku\":\"HU12839\",\"advertiser\":\"Walmart\"}]}"

	if strings.Compare(requestRecorder.Body.String(), expectedResponse) == 0 {
		t.Error("Response body does not match")
	}
}

func TestGetProduct(t *testing.T) {
	var products []Product
	products = append(products, Product{ID: "123", Name: "soap", SKU: "HU12838", Advertiser: "Walmart"})

	req, err := http.NewRequest("GET", "/products/123", nil)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", req)

	requestRecorder := newRequestRecorder(req, "GET", "/products/123", GetProducts)

	if requestRecorder.Code != 200 {
		t.Error("Expected response code to be 200")
	}

	expectedResponse := "{\"meta\":null,\"data\":[{\"id\":\"123\",\"name\":\"soap\",\"sku\":\"HU12838\",\"advertiser\":\"Walmart\"}]}"

	fmt.Printf("%+v\n", requestRecorder.Body.String())
	fmt.Printf("%+v\n", expectedResponse)

	if strings.Compare(requestRecorder.Body.String(), expectedResponse) == 0 {
		t.Error("Response body does not match")
	}
}

func TestDeleteProduct(t *testing.T) {
	populateProducts()

	request, err := http.NewRequest("DELETE", "/products/123", nil)

	if err != nil {
		t.Fatal(err)
	}

	requestRecorder := newRequestRecorder(request, "DELETE", "/products/123", DeleteProduct)

	if requestRecorder.Code != 200 {
		t.Error("Expected response code to be 200")
	}

	expectedResponse := "{\"meta\":null,\"data\":[{\"id\":\"124\",\"name\":\"brush\",\"sku\":\"HU12836\",\"advertiser\":\"Walmart\"},{\"id\":\"125\",\"name\":\"rope\",\"sku\":\"HU12839\",\"advertiser\":\"Walmart\"}]}"

	if strings.Compare(requestRecorder.Body.String(), expectedResponse) == 0 {
		t.Error("Response body does not match")
	}
}

func newRequestRecorder(req *http.Request, method string, strPath string, fnHandler func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.Handle(method, strPath, fnHandler)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}
