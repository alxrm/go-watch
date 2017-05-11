package main

import "log"

type File struct {
  Hash string `json:"hash"`
  Path string `json:"path"`
}

func (file *File) save(db *Database) error {
  return db.exec(`INSERT INTO files(hash, path) VALUES(?, ?);`, fileToRaw(file)...)
}

func (file *File) remove(db *Database) error {
  return db.exec(`DELETE FROM files WHERE hash = ? and path = ?`, file.Hash, file.Path)
}

func clearFiles(db *Database) error {
  return db.exec(`DELETE FROM files`)
}

func allFiles(db *Database) []File {
  res, err := db.queryAll(`SELECT hash, path FROM files;`, fileToFields, fieldsToFile)

  if err != nil {
    log.Println(err)
    return []File{}
  }

  return toFiles(res)
}

func filesByHash(db *Database, hash string) []File {
  args := []interface{}{hash}
  res, err := db.query(`SELECT hash, path FROM files WHERE hash = ?;`, args, fileToFields, fieldsToFile)

  if err != nil {
    log.Println(err)
    return []File{}
  }

  return toFiles(res)
}
