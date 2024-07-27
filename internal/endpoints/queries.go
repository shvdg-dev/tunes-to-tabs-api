package endpoints

/*
CreateEndpointsTableQuery is a query to create the endpoints table.
It is used to store endpoints, taken from external sources.
+---------------------------------------------------------------+
|   source_id  | category   | type      | url                   |
+---------------------------------------------------------------+
| 1001         | artist     | web       | /artist/{artistID}    |
| 1001         | track      | web       | /track/{trackID}      |
| 1003         | tab        | api       | /tab/{trackID}        |
+---------------------------------------------------------------+

It contains the following columns:
  - 'source_id': This is the ID of the external source from which the data was referenced.
  - 'category': This denotes the category of an external reference.
  - 'type': This denotes the type.
  - 'url': This is the endpoint, which has to be formatted with the corresponding IDs/references, as stored in the 'references' table.
*/
const CreateEndpointsTableQuery = `
	CREATE TABLE IF NOT EXISTS "endpoints" (
	   source_id INT NOT NULL,
	   category VARCHAR(250) NOT NULL,
	   type VARCHAR(250) NOT NULL,
	   url VARCHAR(250) NOT NULL,
	   UNIQUE(source_id, category, type, URL),  
	   CONSTRAINT fk_source FOREIGN KEY(source_id) REFERENCES sources(id)
	);
`

// DropEndpointsTableQuery is a SQL query to drop the 'endpoints' table from the database.
const DropEndpointsTableQuery = `
	DROP TABLE IF EXISTS "endpoints";
`
