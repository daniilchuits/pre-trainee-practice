package main

import (
	"fmt"
	"log"

	"github.com/daniilchuits/pre-trainee-practice/internal/booking"
)

func main() {

	s := booking.NewService()
	s.AddRoom("orion", "Orion")

	a, err := s.Book("orion", 9, 11) // ok
	if err != nil {
		log.Fatalln(1, err)
	}
	_, err = s.Book("orion", 11, 12) // ok — touches 11, no overlap
	if err != nil {
		log.Fatalln(2, err)
	}
	_, err = s.Book("orion", 16, 18) // ok
	if err != nil {
		log.Fatalln(3, err)
	}

	if err = s.Cancel(a.ID); err != nil {
		log.Fatalln(4, err)
	}

	_, err = s.Book("orion", 9, 10) // ok
	if err != nil {
		log.Fatalln(5, err)
	}

	fmt.Println(s)

}
