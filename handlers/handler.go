package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	data "github.com/varun-muthanna/data"
	protos "github.com/varun-muthanna/products-api/currency"
)

// http handler
type Products struct {
	l  *log.Logger
	cc protos.CurrencyClient
}

func NewProduct(l *log.Logger, cc protos.CurrencyClient) *Products {
	return &Products{l, cc}
}

// func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		p.getProducts(w, r)
// 		return
// 	}

// 	if r.Method == http.MethodPost {
// 		p.addProducts(w, r)
// 		return
// 	}

// 	if r.Method == http.MethodPut {
// 		//expecting an ID
// 		reg := regexp.MustCompile(`/([0-9]+)`)
// 		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

// 		if len(g) != 1 {
// 			http.Error(w, "Invalid", http.StatusBadRequest)
// 			return
// 		}

// 		if len(g[0]) != 2 {
// 			http.Error(w, "Invalid", http.StatusBadRequest)
// 			return
// 		}

// 		idString := g[0][1]

// 		id, err := strconv.Atoi(idString)

// 		if err != nil {
// 			http.Error(w, "Invalid", http.StatusBadRequest)
// 			return
// 		}

// 		p.updateProducts(id, w, r)

// 		return

// 	}

// 	w.WriteHeader(http.StatusMethodNotAllowed)

// }

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Getting Products")

	lp := data.GetProducts()

	jsonData, err := json.Marshal(lp)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)

}

func (p *Products) AddProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Adding Products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	//p.l.Printf("%#v", prod)

	data.AddProduct(&prod)
	return

}

func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err1 := strconv.Atoi(vars["id"])

	if err1 != nil {
		http.Error(w, "Unable to convert", http.StatusBadRequest)
	}

	fmt.Printf("Updating Products with id %d", id)

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	e := data.UpdateProduct(id, &prod)

	if e == data.ErrProductNotFound {
		http.Error(w, "Product Not found", http.StatusNotFound)
	}

	return
}

func (p *Products) ConvertProducts(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Unable to convert", http.StatusBadRequest)
		return
	}

	fmt.Printf("Converting price of Product with id %d\n", id)

	// Safely retrieve the product from the context
	prod, _, err1 := data.FindProduct(id)

	if err1 != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	rr := &protos.RateRequest{
		Source:      protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value["EUR"]),
	}

	resp, err := p.cc.GetRate(context.Background(), rr)

	if err != nil {
		p.l.Println("[Error] error getting new rate", err)
		http.Error(w, "Error getting new rate", http.StatusInternalServerError)
		return
	}

	// Convert the price
	prod.Price = prod.Price * resp.Rate

	lp := data.GetProducts()

	jsonData, err := json.Marshal(lp)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)

	return
}

type KeyProduct struct{}

func (p Products) MiddlewareValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)

		if err != nil {
			http.Error(w, "Unable to unmarshall", http.StatusBadRequest)
			return
		}

		err2 := prod.Validate()

		if err2 != nil {
			p.l.Println("[ERROR] validating product", err2)
			http.Error(w, "Error validating product", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
