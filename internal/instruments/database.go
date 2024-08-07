package instruments

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// SetupOperations represents setup operations related to instruments in the database.
type SetupOperations interface {
	CreateInstrumentsTable()
	DropInstrumentsTable()
}

// SetupService is for setting up the instruments table in the database.
type SetupService struct {
	logic.DbOperations
}

// NewSetupService creates a new instance of the SetupService struct.
func NewSetupService(database logic.DbOperations) SetupOperations {
	return &SetupService{DbOperations: database}
}

// CreateInstrumentsTable creates an instruments table if it doesn't already exist.
func (t *SetupService) CreateInstrumentsTable() {
	_, err := t.Exec(CreateInstrumentsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'instruments' table.")
	}
}

// DropInstrumentsTable drops the instruments table if it exists.
func (t *SetupService) DropInstrumentsTable() {
	_, err := t.Exec(DropInstrumentsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'instruments' table.")
	}
}
