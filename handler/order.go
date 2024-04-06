package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Order")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all Orders")
}

func (o *Order) GetbyID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an Order ID")
}

func (o *Order) UpdatebyID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an Order ID")
}

func (o *Order) DeletebyID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an Order ID")
}
