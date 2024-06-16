package internal

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	api "github.com/shvdg-dev/tunes-to-tabs-api/pkg"
	diff "github.com/shvdg-dev/tunes-to-tabs-api/pkg/difficulties"
	inst "github.com/shvdg-dev/tunes-to-tabs-api/pkg/instruments"
	"log"
)

// Seeder helps with deleting data from the database
type Seeder struct {
	API     *api.API
	Factory *Factory
}

// NewSeeder creates a new instance of Seeder
func NewSeeder(api *api.API) *Seeder {
	return &Seeder{API: api}
}

// SeedTables attempts to seed the database with the minimally required values and dummy data.
func (s *Seeder) SeedTables() {
	s.minimumSeed()
	s.initFactory()
	s.dummySeed()
}

func (s *Seeder) initFactory() {
	instruments := s.API.Instruments.GetInstruments()
	difficulties := s.API.Difficulties.GetDifficulties()
	s.Factory = NewFactory(instruments, difficulties)
}

// minimumSeed when permitted, seeds the database with the minimally required values.
func (s *Seeder) minimumSeed() {
	if !logic.GetEnvValueAsBoolean(KeyDatabaseAllowMinimumSeedingCommand) {
		log.Println("It is not allowed to seed the database with the minimally required values.")
		return
	}
	s.seedAdmin()
	s.seedInstruments()
	s.seedDifficulties()
}

// seedAdmin inserts an administrator user into the database.
func (s *Seeder) seedAdmin() {
	email := logic.GetEnvValueAsString(KeyAdminInitialEmail)
	password := logic.GetEnvValueAsString(KeyAdminInitialPassword)
	if email != "" && password != "" {
		s.API.Users.InsertUser(email, password)
	} else {
		log.Println("Did not insert the initial admin account as no credentials were defined")
	}
}

// seedInstruments seeds the instruments table with the default instruments.
func (s *Seeder) seedInstruments() {
	s.API.Instruments.InsertInstruments(
		inst.NewInstrument(InstrumentElectricGuitar),
		inst.NewInstrument(InstrumentAcousticGuitar),
		inst.NewInstrument(InstrumentBassGuitar),
		inst.NewInstrument(InstrumentDrums))
}

// seedDifficulties seeds the difficulties table with the default difficulties.
func (s *Seeder) seedDifficulties() {
	s.API.Difficulties.InsertDifficulties(
		diff.NewDifficulty(DifficultyEasy),
		diff.NewDifficulty(DifficultyIntermediate),
		diff.NewDifficulty(DifficultyHard),
		diff.NewDifficulty(DifficultyExpert))
}

// dummySeed when permitted, seeds the database with dummy data.
func (s *Seeder) dummySeed() {
	if !logic.GetEnvValueAsBoolean(KeyDatabaseAllowDummySeedingCommand) {
		log.Println("It is not allowed to seed the database with dummy data.")
		return
	}
	s.seedDummyArtists()
}

// seedDummyArtists inserts dummy artists, tracks, and tabs into the database.
func (s *Seeder) seedDummyArtists() {
	artists := s.Factory.CreateDummyArtists(1)
	// Insert the artist
	s.API.Artists.InsertArtists(artists...)
	for _, artist := range artists {
		// Insert the tracks of an artist
		s.API.Tracks.InsertTracks(artist.Tracks...)
		for _, track := range artist.Tracks {
			// Insert the tabs of a track
			s.API.Tabs.InsertTabs(track.Tabs...)
		}
	}
}
