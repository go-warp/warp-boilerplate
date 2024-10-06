package main

import (
	"context"
	"log"

	"github.com/sitnikovik/go-grpc-api-template/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	if err = a.Run(); err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
