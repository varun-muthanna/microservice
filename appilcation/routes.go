package application

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/varun-muthanna/handler"

	"github.com/go-chi/chi/middleware"
)

func loadOrderRoutes(router chi.Router) {

	orderHandler := &handler.Order{}

	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetbyID)
	router.Put("/{id}", orderHandler.UpdatebyID)
	router.Delete("/{id}", orderHandler.DeletebyID)

}

func loadRoutes() *chi.Mux {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", loadOrderRoutes)

	return router

}
