package main

var products []Product

type Product struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	SKU        string `json:"sku"`
	Advertiser string `json:"advertiser"`
}

func populateProducts() {
	products = append(products, Product{ID: "123", Name: "soap", SKU: "HU12838", Advertiser: "Walmart"})
	products = append(products, Product{ID: "124", Name: "brush", SKU: "HU12836", Advertiser: "Walmart"})
	products = append(products, Product{ID: "125", Name: "rope", SKU: "HU12839", Advertiser: "Walmart"})
}
