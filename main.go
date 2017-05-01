package main

import (
  "log"

  _ "github.com/mattn/go-sqlite3"
)

func main() {
  _, err := makeDB(databaseFile, createStatement)

  if err != nil {
    log.Fatal(err)
    return
  }

  //del := 1000
  //
  //watcher := Watcher{
  //  IntervalMillis: del,
  //  RootDir: "/",
  //  Database: db,
  //  OnCheck: func(hash string) bool {
  //    log.Println("Checking...")
  //    return true
  //  },
  //}
  //
  //watcher.watch()
  //
  //time.Sleep(5 * time.Second)
  //
  //watcher.stop()


  //printDirContents("/Users/alex/Desktop/Dr. Gross.png")

  if err != nil {
    log.Fatal(err)
    return
  }

  log.Println(s)
}
