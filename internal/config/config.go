package config

import (
	"context"

	"github.com/joho/godotenv"
)

// Init initializes the configuration
func Init(_ context.Context) error {
	return godotenv.Load(".env")
}
