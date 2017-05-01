package main

const (
	databaseFile    = "files.db"
	createStatement = `
		CREATE TABLE IF NOT EXISTS files (
			id int not null,
			hash varchar(255) not null,
			path varchar(255) not null
		);
	`
)
