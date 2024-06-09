package internal

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	api "github.com/shvdg-dev/tunes-to-tabs-api/pkg"
	"log"
)

// Creator helps with deleting data from the database
type Creator struct {
	API *api.API
}

// NewCreator creates a new instance of Creator
func NewCreator(API *api.API) *Creator {
	return &Creator{API: API}
}

// CreateTables when permitted, creates tables in the database
func (c *Creator) CreateTables() {
	if !logic.GetEnvValueAsBoolean(KeyDatabaseAllowCreatingCommand) {
		log.Fatalf("It is not allowed to create new tables for the database")
	}
	c.API.Artists.CreateArtistsTable()
	c.API.Sessions.CreateSessionsTable()
	c.API.Users.CreateUsersTable()
}