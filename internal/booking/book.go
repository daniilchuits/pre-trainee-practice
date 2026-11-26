package booking

import (
	"fmt"
	"strconv"
)

const (
	dayStart = 0
	dayEnd   = 24
)

func (s *Service) Book(roomID string, start, end int) (Reservation, error) {

	found := false
	for id := range s.rooms {
		if id == roomID {
			found = true
			break
		}
	}

	if !found {
		return Reservation{}, fmt.Errorf("Error %w: room-id=%s", ErrNotFound, roomID)
	}

	if dayStart > start || start > end || end > dayEnd {
		return Reservation{}, ErrInvalid
	}

	for _, reserv := range s.reservations {
		if roomID == reserv.RoomID && !reserv.Cancelled {
			if reserv.Start < end && start < reserv.End {
				return Reservation{}, fmt.Errorf(
					"Error %w with reservation-id=%s",
					ErrConflict, reserv.ID,
				)
			}
		}
	}

	reservationID := hashReservationID(len(s.reservations))

	reserv := Reservation{
		ID:        reservationID,
		RoomID:    roomID,
		Start:     start,
		End:       end,
		Cancelled: false,
	}
	s.reservations[reservationID] = reserv

	return reserv, nil
}

func hashReservationID(countReservations int) string {
	countReservations++
	return strconv.Itoa(countReservations)
}
