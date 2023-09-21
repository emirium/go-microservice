package main

import (
	"context"
	"fmt"

	"github.com/emirium/go-microservice/application"
)
func main() {
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start the app:", err)
	}
}