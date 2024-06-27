package data

import (
	"fmt"
	"time"
)

type Product struct {
	ID          int     `json:"id"` //repsentation in JSON
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"` //ommitted in JSON
	DeletedOn   string  `json:"-"`
}

func GetProducts() []*Product {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNext()
	productList = append(productList, p)
}

func getNext() int {
	return productList[len(productList)-1].ID + 1
}

func UpdateProduct(id int, p *Product) error {
	_, i, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[i] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},

	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
