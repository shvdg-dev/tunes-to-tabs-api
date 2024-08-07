package endpoints

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DataOperations represents operations related to endpoints in the database.
type DataOperations interface {
	InsertEndpoints(endpoints ...*Endpoint)
	InsertEndpoint(endpoint *Endpoint)
}

// DataService is for managing endpoints.
type DataService struct {
	logic.DbOperations
}

// NewDataService creates a new instance of DataService.
func NewDataService(database logic.DbOperations) DataOperations {
	return &DataService{DbOperations: database}
}

// InsertEndpoints inserts multiple records into the endpoints table.
func (d *DataService) InsertEndpoints(endpoints ...*Endpoint) {
	for _, endpoint := range endpoints {
		d.InsertEndpoint(endpoint)
	}
}

// InsertEndpoint inserts a record into the endpoints table.
func (d *DataService) InsertEndpoint(endpoint *Endpoint) {
	_, err := d.Exec(insertEndpointQuery, endpoint.SourceID, endpoint.Category, endpoint.Type, endpoint.URL)
	if err != nil {
		log.Printf(
			"Failed to insert endpoint with Source '%d', Category '%s', Type '%s', and URL '%s': %s",
			endpoint.SourceID, endpoint.Category, endpoint.Type, endpoint.URL, err.Error(),
		)
	} else {
		log.Printf(
			"Successfully inserted endpoint into the 'endpoints' table with Source '%d', Category '%s', Type '%s', and URL '%s'",
			endpoint.SourceID, endpoint.Category, endpoint.Type, endpoint.URL,
		)
	}
}
