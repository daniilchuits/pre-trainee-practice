package booking

import "errors"

type Room struct {
	ID   int
	Name string
}

type Service struct {
	rooms        []Room
	reservations []Reservation
}

func NewService() *Service {
	return &Service{}
}

type Reservation struct {
	ID        string
	RoomID    string
	Start     int // inclusive
	End       int // exclusive
	Cancelled bool
}

var (
	ErrNotFound  = errors.New("booking: not found")
	ErrConflict  = errors.New("booking: slot conflict")
	ErrInvalid   = errors.New("booking: invalid argument")
	ErrCancelled = errors.New("booking: already cancelled")
)
