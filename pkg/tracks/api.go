package tracks

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// API is for managing tracks of songs.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// CreateTracksTable creates the tracks table if it doesn't already exist.
func (a *API) CreateTracksTable() {
	_, err := a.Database.DB.Exec(createTracksTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'tracks' table")
	}
}

// DropTracksTable drops the tracks table if it exists.
func (a *API) DropTracksTable() {
	_, err := a.Database.DB.Exec(dropTracksTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'tracks' table")
	}
}

// InsertTracks inserts multiple tracks into the tracks table.
func (a *API) InsertTracks(tracks ...*Track) {
	for _, track := range tracks {
		a.InsertTrack(track)
	}
}

// InsertTrack inserts a track into the tracks table.
func (a *API) InsertTrack(track *Track) {
	_, err := a.Database.DB.Exec(insertTrackQuery, track.ID, track.Title, track.Duration)
	if err != nil {
		log.Printf("Failed to insert track with title '%s': %s", track.Title, err.Error())
	} else {
		log.Printf("Successfully inserted track into the 'tracks' table with title '%s'", track.Title)
	}
}
