package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/varun-muthanna/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewProduct(l)

	sm := http.NewServeMux() // default HandleFunc calls ServeMux which is implemented by ServeHTTP
	sm.Handle("/", hh)

	s := &http.Server{
		Addr:         ":8000",           // configure the bind address
		Handler:      sm,                // set the default handler
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
