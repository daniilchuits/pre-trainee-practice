package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Panicln("'.env' is not loaded:", err)
	}

	if os.Getenv("TESTDATA_FILEPATH") == "" {
		log.Panic("Empty variable 'TESTDATA_FILEPATH' in .env")
	}
}
