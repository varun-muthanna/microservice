package main

import (
	"context"
	"fmt"

	"github.com/varun-muthanna/application"
)

func main() {

	app := application.New()

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println("Failed to start App", err)
	}

}
