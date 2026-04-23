package main

import (
	"log"

	"github.com/SaikatDeb12/ecommerce/internal/utils"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env", err)
	}

	utils.SecretKey = utils.GetEnvVariables("SECET_KEY")
}
