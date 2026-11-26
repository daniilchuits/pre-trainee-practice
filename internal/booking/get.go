package booking

import "fmt"

func (s *Service) Get(reservationID string) (Reservation, error) {

	reservation, ok := s.reservations[reservationID]
	if ok {
		return reservation, nil
	}
	return Reservation{}, fmt.Errorf(
		"Error %w: reservation-id=%s",
		ErrNotFound, reservationID,
	)
}
