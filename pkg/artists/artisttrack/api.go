package artisttrack

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// API is for managing artists.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// LinkArtistToTrack inserts a link between an artist and a track into the artist_track table.
func (a *API) LinkArtistToTrack(artistId, trackId string) {
	_, err := a.Database.DB.Exec(insertArtistTrackQuery, artistId, trackId)
	if err != nil {
		log.Printf("Failed linking artist with ID '%s' and track with ID '%s': %s", artistId, trackId, err.Error())
	} else {
		log.Printf("Successfully linked artist with ID '%s' and track with ID '%s'", artistId, trackId)
	}
}

// GetTrackIDs retrieves the track IDs for the provided internal artist IDs.
func (a *API) GetTrackIDs(artistID ...string) ([]string, error) {
	rows, err := a.Database.DB.Query(getTrackIDs, artistID)
	if err != nil {
		return nil, err
	}

	var trackIDs []string

	for rows.Next() {
		var trackID string
		err := rows.Scan(&trackID)
		if err != nil {
			return nil, err
		}
		trackIDs = append(trackIDs, trackID)
	}

	return trackIDs, nil
}
