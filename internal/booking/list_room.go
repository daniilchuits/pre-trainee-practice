package booking

import "fmt"

func (s *Service) ListRoom(roomID string) ([]Reservation, error) {

	var reservations []Reservation
	for _, reserv := range s.reservations {
		if reserv.ID == roomID {
			reservations = append(reservations, reserv)
		}
	}
	if len(reservations) == 0 {
		return nil, fmt.Errorf("Error %v; room-id=%s", ErrNotFound, roomID)
	}
	return reservations, nil
}
