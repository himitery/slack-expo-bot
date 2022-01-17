package main

import (
	"github.com/himitery/slack-expo-bot/src"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load("env/.env")
	if err != nil {
		log.Fatalf("[godotenv.Load() Error] %s\n", err)
	}

	app.Init()
}
