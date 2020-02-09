package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product feindes the structure for the API request
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products ..
type Products []*Product

// ToJSON ...
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	fmt.Printf("ToJSON: %v\n", p)
	return e.Encode(p)
}

// GetProducts ..
//func GetProducts() []*Product {
func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "yada",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Americano",
		Description: "Espresso and loads of water",
		Price:       1.90,
		SKU:         "foo",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
