package main

import (
	"log"

	"github.com/daniilchuits/pre-trainee-practice/internal/booking"
)

func main() {

	s := booking.NewService()

	if err := s.AddRoom("or", "or"); err != nil {
		log.Fatalln("Adding error:", err)
	}
	if err := s.AddRoom("or", "or"); err != nil {
		log.Fatalln("Adding error:", err)
	}

}
