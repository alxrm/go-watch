package main

import (
  _ "github.com/mattn/go-sqlite3"
  "log"
)

func main() {
  kill := make(chan bool)
  db, err := makeDB(databaseFile, createStatement)

  if err != nil {
    log.Fatal(err)
    return
  }

  clearFiles(db)

  file := File{path:"/Users/alex/Desktop/Dr. Gross.png", hash:"b5bab97b55d8f8cebefdbb8f54776ba1"}
  file.save(db)

  watcher := Watcher{
    IntervalMillis: 1000,
    Database: db,
    OnObserved: func(file *File, path string) {
      log.Println(path)
      kill <- true
    },
  }

  watcher.watch()

  <-kill
}
