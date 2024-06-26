package tracks

/*
+--------------------------------------+--------------------+------------+
|                 ID                   |       Title        |  Duration  |
+--------------------------------------+--------------------+------------+
| 123e4567-e89b-12d3-a456-426614174000 | Sweet Child O'Mine |  356000    |
| 123e4567-e89b-12d3-a456-426614174001 | Stairway to Heaven |  483000    |
+--------------------------------------+--------------------+------------+

This table is used to store the tracks of songs.

It contains the following columns:
  - 'ID': The UUID that uniquely identifies a track in our system.
  - 'Title': The title of the track.
  - 'Duration': The duration of the track in milliseconds.
*/
const createTracksTableQuery = `
	CREATE TABLE IF NOT EXISTS tracks (
	   ID UUID PRIMARY KEY,
	   Title VARCHAR(500) NOT NULL,
	   Duration NUMERIC NOT NULL
	);
`

// dropTracksTableQuery is a SQL query to drop the 'tracks' table from the database.
const dropTracksTableQuery = `
	DROP TABLE IF EXISTS tracks;
`

// insertTrackQuery is a SQL query to insert a track into the 'tracks' table.
const insertTrackQuery = `
	INSERT INTO tracks (id, title, duration)
    VALUES ($1, $2, $3) 
`

/*
+--------------------------------------+--------------------------------------+
|               TrackID                |                 TabID                |
+--------------------------------------+--------------------------------------+
| 123e4567-e89b-12d3-a456-426614174050 | 123e4567-e89b-12d3-a456-426614174000 |
| 123e4567-e89b-12d3-a456-426614174051 | 123e4567-e89b-12d3-a456-426614174001 |
+--------------------------------------+--------------------------------------+

This table is used to link the 'tracks' table and the 'tabs' table.

- 'TrackID': The UUID that uniquely identifies a track from the 'tracks' table.
- 'TabID': The UUID that uniquely identifies a tab from the 'tabs' table.
*/
const createTrackTabTableQuery = `
	CREATE TABLE IF NOT EXISTS track_tab  (
	   TrackID UUID REFERENCES tracks (ID),
	   TabID UUID REFERENCES tabs (ID),
	   PRIMARY KEY (TrackID, TabID)
	);
`

// dropTrackTabTableQuery is a SQL query that drops the 'track_tab' table from the database.
const dropTrackTabTableQuery = `
	DROP TABLE IF EXISTS track_tab;
`

// insertTrackTabQuery is a SQL query to insert a link from a track to a tab in the 'track_tab' table.
const insertTrackTabQuery = `
	INSERT INTO track_tab (trackId, tabId)
    VALUES ($1, $2) 
`
