package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type Product struct {
	ID          int     `json:"id" ` //repsentation in JSON
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku"   validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"` //ommitted in JSON
	DeletedOn   string  `json:"-"`
}

func (p *Product) Validate() error {
	val := validator.New()
	val.RegisterValidation("sku", validateSKU)

	return val.Struct(p)
}

func validateSKU(f validator.FieldLevel) bool { //interface providing context of the field to be validated
	//sku = abs-abd-abcd
	regex := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]`)
	matches := regex.FindAllString(f.Field().String(), -1)

	return len(matches) == 1
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
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func UpdateProduct(id int, p *Product) error {
	_, i, err := FindProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[i] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func FindProduct(id int) (*Product, int, error) {
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
