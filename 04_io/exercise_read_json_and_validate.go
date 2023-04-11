package main

import (
	"encoding/json"
	"log"
)

var productJson = []byte(`
{
	"id": 30,
	"title": "Key Holder",
	"description": "Attractive DesignMetallic materialFour key hooksReliable & DurablePremium Quality",
	"price": 30,
	"discountPercentage": 2.92,
	"rating": 4.92,
	"stock": 54,
	"brand": "Golden",
	"category": "home-decoration",
	"thumbnail": "https://i.dummyjson.com/data/products/30/thumbnail.jpg",
	"images": [
		"https://i.dummyjson.com/data/products/30/1.jpg",
		"https://i.dummyjson.com/data/products/30/2.jpg",
		"https://i.dummyjson.com/data/products/30/3.jpg",
		"https://i.dummyjson.com/data/products/30/thumbnail.jpg"
	]
}`)

type Products struct {
	Products []Product
	Total    int
	Skip     int
	Limit    int
}

// TODO: complete the fields
// TODO: fix the error when unmarshalling/deserializing
type Product struct {
	Id                 int
	Title              string
	Images             []string
	DiscountPercentage float64
}

// Objective read the static json with it's fields
func TestProductJsonRead() {
	var product Product
	if err := json.Unmarshal(productJson, &product); err != nil {
		log.Fatal("error on unmarshall: ", err)
	}
	if product.Id != 30 {
		log.Fatalf("expected product.ID to be 30 but got: %d", product.Id)
	}
	if product.Title != "Key Holder" {
		log.Fatalf("expected product.ID to be Key Holder but got: %s", product.Title)
	}
	if len(product.Images) != 4 {
		log.Fatalf("expected product.Images size to be 4 but got: %d", len(product.Images))
	}
}

func readAllProductsFromFile() Products {
	// TODO: implement me
	return Products{}
}

// Ojective read all the products from the file products.json
func TestReadAllProducts() {
	// TODO: complete this function
	products := readAllProductsFromFile()

	if len(products.Products) != 30 {
		log.Fatalf("expected products.Skip to be 0 but got: %d", len(products.Products))
	}
	if products.Total != 100 {
		log.Fatalf("expected products.Skip to be 0 but got: %d", products.Total)
	}
	if products.Skip != 0 {
		log.Fatalf("expected products.Skip to be 0 but got: %d", products.Skip)
	}
	if products.Limit != 30 {
		log.Fatalf("expected products.Limit to be 30 but got: %d", products.Limit)
	}
}

// Finish the implementation
func findId(products Products, id int) (*Product, error) {
	return nil, nil
}

// Objective do return an error if id 40 is not found
func TestFindIDError() {
	products := readAllProductsFromFile()
	product, err := findId(products, 40)
	if product != nil {
		log.Fatal("product should be nil but got: ", product)
	}
	if err == nil {
		log.Fatal("error should be NOT nil and the error was", err)
	}
}

// Objective do give an error if id 40 is not found
func TestFindProductByID() {
	products := readAllProductsFromFile()
	product, err := findId(products, 12)
	if product == nil {
		log.Fatal("product should NOT be nil but got nil")
	}
	if err != nil {
		log.Fatal("error should be nil but got: ", err)
	}
	if product.Id != 12 {
		log.Fatal("expected product.Id to be 12 but got: ", product.Id)
	}
}

func main() {
	TestProductJsonRead()

	TestReadAllProducts()

	//
	TestFindIDError()
	TestFindProductByID()
}
