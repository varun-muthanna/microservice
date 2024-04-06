package application

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

func New() *App { //CONSTRUCTOR

	app := &App{
		router: loadRoutes(),
	}

	return app

}

func (a *App) Start(ctx context.Context) error {

	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()

	if err != nil {
		return fmt.Errorf("failed to listen , %w", err)
	}

	return nil

}
