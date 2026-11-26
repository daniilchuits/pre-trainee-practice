package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/daniilchuits/pre-trainee-practice/internal/booking"
)

func main() {

	// now I understand why table-drivern tests are so important
	// they wouldn't have saved me all the time
	// but they would have saved a large part of it

	// and I understood on practise how important is
	// to use correct data structure

	s := booking.NewService()
	if err := s.AddRoom("", "x"); err != nil {
		if !errors.Is(err, booking.ErrInvalid) {
			log.Panicln(1, err)
		}
	}

	if err := s.AddRoom("orion", "orion"); err != nil {
		log.Panicln(3, err)
	}

	if err := s.AddRoom("orion", "dup"); err != nil {
		if !errors.Is(err, booking.ErrInvalid) {
			log.Panicln(3, err)
		}
	}

	if _, err := s.Book("missing", 9, 10); err != nil {
		if !errors.Is(err, booking.ErrNotFound) {
			log.Panicln(4, err)
		}
	}

	if _, err := s.Book("orion", 9, 9); err != nil {
		if !errors.Is(err, booking.ErrInvalid) {
			log.Panicln(5, err)
		}
	}

	if _, err := s.Book("orion", -1, 2); err != nil {
		if !errors.Is(err, booking.ErrInvalid) {
			log.Panicln(6, err)
		}
	}

	if _, err := s.Book("orion", 10, 25); err != nil {
		if !errors.Is(err, booking.ErrInvalid) {
			log.Panicln(7, err)
		}
	}

	//Book("orion", 10, 12) when [9,11) is active
	a, err := s.Book("orion", 9, 11)
	if err != nil {
		log.Panicln(8, err)
	}
	if _, err = s.Book("orion", 10, 12); err != nil {
		if !errors.Is(err, booking.ErrConflict) {
			log.Panicln(9, err)
		}
	}

	if err = s.Cancel("nope"); err != nil {
		if !errors.Is(err, booking.ErrNotFound) {
			log.Panicln(10, err)
		}
	}

	if err = s.Cancel(a.ID); err != nil {
		log.Panicln("11. Error must be nil, have:", err)
	}
	if err = s.Cancel(a.ID); err != nil {
		if !errors.Is(err, booking.ErrCancelled) {
			log.Panicln(12, err)
		}
	}

	if _, err := s.Get("nope"); err != nil {
		if !errors.Is(err, booking.ErrNotFound) {
			log.Panicln(13, err)
		}
	}

	if _, err := s.ListRoom("missing"); err != nil {
		if !errors.Is(err, booking.ErrNotFound) {
			log.Panicln(14, err)
		}
	}

	if _, err = s.Book("orion", 9, 11); err != nil {
		log.Panicln(15, err)
	}
	if _, err = s.Book("orion", 11, 12); err != nil {
		log.Panicln(16, err)
	}
	fmt.Println("OK")

}
