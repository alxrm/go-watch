package main

import "log"

type File struct {
  Hash string
  Path string
}

func (file *File) save(db *Database) {
  err := db.exec(`INSERT INTO files(hash, path) VALUES(?, ?);`, fileToRaw(file)...)

  if err != nil {
    log.Fatal(err)
  }
}

func (file *File) remove(db *Database) {
  err := db.exec(`DELETE FROM files WHERE hash = ? and path = ?`, file.Hash, file.Path)

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
    return []File{}
  }

  return toFiles(res)
}
