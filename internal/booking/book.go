package booking

import (
	"crypto/sha256"
	"encoding/hex"
)

const (
	dayStart = 0
	dayEnd   = 24
)

func (s *Service) Book(roomID string, start, end int) (Reservation, error) {

	found := false
	for _, room := range s.rooms {
		if room.ID == roomID {
			found = true
			break
		}
	}

	if !found {
		return Reservation{}, ErrNotFound
	}

	if dayStart > start || start > end || end > dayEnd {
		return Reservation{}, ErrInvalid
	}

	for _, reserv := range s.reservations {
		if roomID == reserv.RoomID && !reserv.Cancelled {
			if reserv.Start < end && start < reserv.End {
				return Reservation{}, ErrConflict
			}
		}
	}

	var lastID string
	if len(s.reservations) > 0 {
		lastID = s.reservations[len(s.reservations)-1].ID
	}

	reserv := Reservation{
		ID:        hashReservationID(lastID, roomID),
		RoomID:    roomID,
		Start:     start,
		End:       end,
		Cancelled: false,
	}
	s.reservations = append(s.reservations, reserv)

	return reserv, nil
}

func hashReservationID(lastID, roomID string) string {
	sum := sha256.Sum256(append([]byte(lastID), []byte(roomID)...))
	return hex.EncodeToString(sum[:])
}
