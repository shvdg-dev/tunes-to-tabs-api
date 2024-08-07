package tracks

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// SetupOperations represents operations related to tracks in the database.
type SetupOperations interface {
	CreateTracksTable()
	DropTracksTable()
}

// SetupService is for managing tracks of songs.
type SetupService struct {
	logic.DbOperations
}

// NewSetupService creates a new instance of the SetupService struct.
func NewSetupService(database logic.DbOperations) SetupOperations {
	return &SetupService{DbOperations: database}
}

// CreateTracksTable creates the tracks table if it doesn't already exist.
func (s *SetupService) CreateTracksTable() {
	_, err := s.Exec(CreateTracksTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'tracks' table")
	}
}

// DropTracksTable drops the tracks table if it exists.
func (s *SetupService) DropTracksTable() {
	_, err := s.Exec(DropTracksTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'tracks' table")
	}
}
