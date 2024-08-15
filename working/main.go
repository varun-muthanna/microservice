package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/varun-muthanna/handlers"
	protos "github.com/varun-muthanna/products-api/currency"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	conn, err := grpc.Dial("localhost:9093", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	currencyclient := protos.NewCurrencyClient(conn)

	hh := handlers.NewProduct(l, currencyclient)

	m := mux.NewRouter()

	getrouter := m.Methods("GET").Subrouter() //only GET

	getrouter.HandleFunc("/{id:[0-9]+}", hh.ConvertProducts)
	getrouter.HandleFunc("/", hh.GetProducts)

	putrouter := m.Methods("PUT").Subrouter()
	putrouter.HandleFunc("/{id:[0-9]+}", hh.UpdateProducts)
	putrouter.Use(hh.MiddlewareValidation)

	postrouter := m.Methods("POST").Subrouter()
	postrouter.HandleFunc("/", hh.AddProducts)
	postrouter.Use(hh.MiddlewareValidation)

	//sm := http.NewServeMux() // default HandleFunc calls ServeMux which is implemented by ServeHTTP
	//sm.Handle("/", hh)

	s := &http.Server{
		Addr:         ":8000",           // configure the bind address
		Handler:      m,                 // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	go func() {
		l.Println("Server listening on port 8000")

		err := s.ListenAndServe()

		if err != nil {
			log.Println("Error starting server:", err)
			return
		}

	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	//wait (blocks until signal is recieved)
	sig := <-ch
	log.Println("Graceful Shutdown starting", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}
