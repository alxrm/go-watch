package main

import (
  "log"

  _ "github.com/mattn/go-sqlite3"
  "fmt"
)

func main() {
  db, err := makeDB("files.db", createFilesCommand)

  if err != nil {
    log.Fatal(err)
    return
  }

  file := File{id:1, hash: "345", path:"123"}
  file.save(db)
  //
  //file.remove(db)
  //
  //clearFiles(db)

  fmt.Println(allFiles(db))
}
