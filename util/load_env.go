package util

import (
				"os"
				"fmt"

				"github.com/joho/godotenv"
)


func LoadEnv() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}
}