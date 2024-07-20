package main

import (
	"fmt"
	logic "github.com/shvdg-dev/base-logic/pkg"
	inter "github.com/shvdg-dev/tunes-to-tabs-api/internal"
	"github.com/shvdg-dev/tunes-to-tabs-api/internal/services"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg"
	"log"
	"os"
)

var (
	config *inter.Config
	tables *services.TableService
	api    *pkg.API
)

// init instantiates all app requirements.
func init() {
	config = initConfig()
	database := initDatabase()
	tables = services.NewTableService(database)
	api = pkg.NewAPI(database)
}

// initConfig initializes the application configuration.
func initConfig() *inter.Config {
	conf, err := inter.NewConfig(inter.PathConfig)
	if err != nil {
		log.Fatalf("Could not load config")
	}
	return conf
}

// initDatabase initializes the database manager.
func initDatabase() *logic.DatabaseManager {
	URL := logic.GetEnvValueAsString(inter.KeyDatabaseURL)
	database := logic.NewDatabaseManager(inter.ValueDatabaseDriver, URL)
	return database
}

// main is the entry point of the application.
func main() {
	handleArgs(os.Args[1:])
}

// handleArgs handles each argument individually.
func handleArgs(args []string) {
	for _, arg := range args {
		handleArg(arg)
	}
}

// handleArgs handles the command line argument and performs the corresponding action.
func handleArg(arg string) {
	switch arg {
	case inter.CommandCreate:
		services.NewCreateService(tables).CreateAll()
	case inter.CommandPurge:
		services.NewPurgeService(tables).Purge()
	case inter.CommandSeed:
		services.NewSeedService(config.Seeding, api).Seed()
	default:
		printErrorAndExit()
	}
}

// printErrorAndExit prints an error message and exits the program with an exit code of 1.
func printErrorAndExit() {
	fmt.Println("Failed to run app, expected 'create', 'purge', or 'seed'")
	os.Exit(1)
}
