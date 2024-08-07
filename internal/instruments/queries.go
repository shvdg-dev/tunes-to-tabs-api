package instruments

/*
CreateInstrumentsTableQuery is a query to create an instruments lookup table.
+----+---------------------+
| id |         name        |
+----+---------------------+
| 1  |  distortion guitar  |
| 2  |  acoustic guitar    |
+----+---------------------+

It contains the following columns:
  - 'id': This is an auto-incrementing integer that uniquely identifies a record.
  - 'name': This column records the name of the instrument.
*/
const CreateInstrumentsTableQuery = `
	CREATE TABLE IF NOT EXISTS instruments (
	   id SERIAL PRIMARY KEY,
	   name VARCHAR(255) NOT NULL
	);
`

// DropInstrumentsTableQuery is a SQL query to drop the 'instruments' table if it exists
const DropInstrumentsTableQuery = `
	DROP TABLE IF EXISTS instruments;
`
