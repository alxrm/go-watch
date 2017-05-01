package main

import "log"

type File struct {
	id   int
	hash string
	path string
}

func (file *File) save(db *Database) {
	err := db.exec(`INSERT INTO files(id, hash, path) VALUES(?, ?, ?);`, fileToRaw(file)...)

	if err != nil {
		log.Fatal(err)
	}
}

func (file *File) remove(db *Database) {
	err := db.exec(`DELETE FROM files WHERE hash = ? and path = ?`, file.hash, file.path)

	if err != nil {
		log.Fatal(err)
	}
}

func clearFiles(db *Database) {
	err := db.exec(`DELETE FROM files`)

	if err != nil {
		log.Fatal(err)
	}
}

func allFiles(db *Database) []File {
	res, err := db.queryAll(`SELECT * FROM files;`, fileToFields, fieldsToFile)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return toFiles(res)
}

func filesByHash(db *Database, hash string) []File {
	args := []interface{}{hash}
	res, err := db.query(`SELECT * FROM files WHERE hash = ?;`, args, fileToFields, fieldsToFile)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return toFiles(res)
}
