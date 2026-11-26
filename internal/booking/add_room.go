package booking

import (
	"errors"
	"fmt"
	"strings"
)

func (s *Service) AddRoom(id, name string) error {

	id = strings.TrimSpace(id)
	if id == "" {
		return errors.New("Empty 'id' after trim")
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("Empty 'name' after trim")
	}

	for _, room := range s.rooms {
		if room.ID == id {
			return fmt.Errorf("Error %w: duplicate id: %s", ErrInvalid, id)
		}
	}
	s.rooms = append(s.rooms, Room{ID: id, Name: name})
	return nil
}
