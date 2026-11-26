package booking

import "fmt"

func (s *Service) Get(reservationID string) (Reservation, error) {

	for _, reserv := range s.reservations {
		if reserv.ID == reservationID {
			return reserv, nil
		}
	}
	return Reservation{}, fmt.Errorf(
		"Error %w: reservation-id=%s",
		ErrNotFound, reservationID,
	)
}
