package main

import (
	"kbot/cmd"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file found")
	}
	cmd.Execute()
}
