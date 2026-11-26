package booking

import "errors"

type Room struct {
	ID   string
	Name string
}

type Service struct {
	rooms        map[string]Room        // <room-id>:<Room>
	reservations map[string]Reservation //	<reservation-id>:<Reservation>
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
