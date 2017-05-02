package main

import (
  _ "github.com/mattn/go-sqlite3"
  "log"
  "github.com/gin-gonic/gin"
  "gopkg.in/olahol/melody.v1"
  "net/http"
)

func main() {
  db, err := makeDB(databaseFile, createStatement)

  if err != nil {
    log.Fatal(err)
    return
  }

  router := gin.Default()
  channel := melody.New()

  router.GET("/", func(c *gin.Context) {
    http.ServeFile(c.Writer, c.Request, "./assets/index.html")
  })

  router.GET("/watcher", func(c *gin.Context) {
    channel.HandleRequest(c.Writer, c.Request)
  })

  channel.HandleConnect(func(s *melody.Session) {
    channel.Broadcast([]byte("Connected, ready to watch!"))
  })

  channel.HandleMessage(func(s *melody.Session, msg []byte) {
    channel.Broadcast(msg)
  })

  watcher := Watcher{
    IntervalMillis: 1000,
    Database: db,
    OnObserved: func(file *File, path string) {
      channel.Broadcast([]byte(path))
    },
    OnStopped: func() {
      channel.Broadcast([]byte("Watcher stopped"))
    },
  }

  watcher.watch()

  router.Run(":5000")
}
