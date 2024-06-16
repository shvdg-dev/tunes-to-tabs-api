package difficulties

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// API is for managing difficulties.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// CreateDifficultiesTable creates a difficulties table.
func (a *API) CreateDifficultiesTable() {
	_, err := a.Database.DB.Exec(createDifficultiesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'difficulties' table.")
	}
}

// DropDifficultiesTable drops the difficulties table.
func (a *API) DropDifficultiesTable() {
	_, err := a.Database.DB.Exec(dropDifficultiesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'difficulties' table.")
	}
}

// InsertDifficulties inserts multiple difficulty levels.
func (a *API) InsertDifficulties(difficulties ...*Difficulty) {
	for _, difficulty := range difficulties {
		a.InsertDifficulty(difficulty)
	}
}

// InsertDifficulty inserts a new difficulty level.
func (a *API) InsertDifficulty(difficulty *Difficulty) {
	_, err := a.Database.DB.Exec(insertDifficultyQuery, difficulty.Name)
	if err != nil {
		log.Printf("Failed inserting difficulty level with name: '%s': %s", difficulty.Name, err.Error())
	} else {
		log.Printf("Successfully inserted difficulty level with name: '%s'", difficulty.Name)
	}
}
