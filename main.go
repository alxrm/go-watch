package main

import (
  "log"

  _ "github.com/mattn/go-sqlite3"
  "fmt"
)

func main() {
  db, err := makeDB(databaseFile, createStatement)

  if err != nil {
    log.Fatal(err)
    return
  }

  fmt.Println(allFiles(db))
}
