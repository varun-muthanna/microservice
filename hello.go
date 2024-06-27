package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/varun-muthanna/data"
)

// http handler
type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProducts(w, r)
		return
	}

	if r.Method == http.MethodPut {
		//expecting an ID
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(w, "Invalid", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(w, "Invalid", http.StatusBadRequest)
			return
		}

		idString := g[0][1]

		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(w, "Invalid", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, w, r)

		return

	}

	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Getting Products")

	lp := data.GetProducts()

	jsonData, err := json.Marshal(lp)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)

}

func (p *Products) addProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Adding Products")

	prod := &data.Product{}

	d := json.NewDecoder(r.Body)
	err := d.Decode(prod)

	if err != nil {
		http.Error(w, "Unable to unmarshall", http.StatusBadRequest)
		return
	}

	//p.l.Printf("%#v", prod)

	data.AddProduct(prod)
	return

}

func (p *Products) updateProducts(id int, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Updating Products with id %d", id)

	prod := &data.Product{}

	d := json.NewDecoder(r.Body)
	err := d.Decode(prod)

	if err != nil {
		http.Error(w, "Unable to unmarshall", http.StatusBadRequest)
		return
	}

	e := data.UpdateProduct(id, prod)

	if e == data.ErrProductNotFound {
		http.Error(w, "Product Not found", http.StatusNotFound)
	}

	return
}
