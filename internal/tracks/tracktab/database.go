package tracktab

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// SetupOperations represents setup operations related to 'track to tab' links.
type SetupOperations interface {
	CreateTrackTabTable()
	DropTrackTabTable()
}

// SetupService is for settings up tracks tables in the database.
type SetupService struct {
	logic.DbOperations
}

// NewSetupService creates a new instance of the SetupService struct.
func NewSetupService(database logic.DbOperations) *SetupService {
	return &SetupService{DbOperations: database}
}

// CreateTrackTabTable creates a track_tab table if it doesn't already exist.
func (s *SetupService) CreateTrackTabTable() {
	_, err := s.Exec(CreateTrackTabTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'track_tab' table")
	}
}

// DropTrackTabTable drops the track_tab table if it exists.
func (s *SetupService) DropTrackTabTable() {
	_, err := s.Exec(DropTrackTabTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'track_tab' table")
	}
}
