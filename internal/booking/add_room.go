package booking

import (
	"fmt"
	"strings"
)

func (s *Service) AddRoom(id, name string) error {

	id = strings.TrimSpace(id)
	if id == "" {
		return fmt.Errorf("Error %w empty 'id' after trim", ErrInvalid)
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return fmt.Errorf("Error %w empty 'name' after trim", ErrInvalid)
	}

	for _, room := range s.rooms {
		if room.ID == id {
			return fmt.Errorf("Error %w: duplicate id: %s", ErrInvalid, id)
		}
	}
	s.rooms = append(s.rooms, Room{ID: id, Name: name})
	return nil
}
