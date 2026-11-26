package booking

func (s *Service) Cancel(reservationID string) error {

	for i, reserv := range s.reservations {
		if reserv.ID == reservationID {
			if reserv.Cancelled {
				return ErrCancelled
			}
			s.reservations[i].Cancelled = true
			return nil
		}
	}
	return ErrNotFound
}
