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

	if os.Getenv("TESTDATA_FILEPATH_WORDS") == "" {
		log.Panic("Empty variable 'TESTDATA_FILEPATH_WORDS' in .env")
	} else if os.Getenv("TESTDATA_FILEPATH_ACCESS") == "" {
		log.Panic("Empty variable 'TESTDATA_FILEPATH_ACCESS' in .env")
	}
}
