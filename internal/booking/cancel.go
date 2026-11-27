package booking

import "fmt"

func (s *Service) Cancel(reservationID string) error {

	reservation, ok := s.reservations[reservationID]
	if !ok {
		return fmt.Errorf("Error %w: reservation-id=%s", ErrCancelled, reservationID)
	}
	if reservation.Cancelled {
		return fmt.Errorf("Error %w: reservation-id=%s", ErrNotFound, reservationID)
	}
	return nil
}
