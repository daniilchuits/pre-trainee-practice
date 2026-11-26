package booking

import "fmt"

func (s *Service) Cancel(reservationID string) error {

	for i, reserv := range s.reservations {
		if reserv.ID == reservationID {
			if reserv.Cancelled {
				return fmt.Errorf("%v: id=%s", ErrCancelled, reservationID)
			}
			s.reservations[i].Cancelled = true
			return nil
		}
	}
	return fmt.Errorf("%v: id=%s", ErrNotFound, reservationID)
}
