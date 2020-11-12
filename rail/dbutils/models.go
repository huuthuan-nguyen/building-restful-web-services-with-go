package dbutils

const train = `
	CREATE TABLE IF NOT EXISTS train (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		driver_name VARCHAR(64) NULL,
		operating_status BOOLEAN
	)
`

const station = `
	CREATE TABLE IF NOT EXISTS station (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(64) NULL,
		opening_time TIME NULL,
		closing_time TIME NULL
	)
`

const schedule = `
	CREATE TABLE IF NOT EXISTS schedule (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		train_id INT,
		station_id INT,
		arrival_time TIME,
		FOREIGN KEY (train_id) REFERENCES train(id),
		FOREIGN KEY (station_id) REFERENCES station(id)
	)
`