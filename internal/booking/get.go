package booking

func (s *Service) Get(reservationID string) (Reservation, error) {

	for _, reserv := range s.reservations {
		if reserv.ID == reservationID {
			return reserv, nil
		}
	}
	return Reservation{}, ErrNotFound
}
