package main

import (
  _ "github.com/mattn/go-sqlite3"
  "log"
  "github.com/gin-gonic/gin"
  "gopkg.in/olahol/melody.v1"
  "net/http"
  "strings"
)

type Context struct {
  Watcher  *Watcher
  Database *Database
  Socket   *melody.Melody
}

func main() {
  db, err := newDB(databaseFile, createStatement)

  if err != nil {
    log.Fatal(err)
    return
  }

  router := gin.Default()
  socket := melody.New()

  watcher := &Watcher{
    IntervalMillis: 1000,
    Database: db,
    OnObserved: func(file *File, path string) {
      socket.Broadcast([]byte("Found: " + path))
    },
    OnStopped: func() {
      socket.Broadcast([]byte("Stopped"))
    },
  }

  context := &Context{
    Watcher: watcher,
    Database: db,
    Socket: socket,
  }

  router.GET("/", func(c *gin.Context) {
    http.ServeFile(c.Writer, c.Request, "./assets/index.html")
  })

  router.GET("/watcher", func(c *gin.Context) {
    socket.HandleRequest(c.Writer, c.Request)
  })

  socket.HandleConnect(func(s *melody.Session) {
    socket.Broadcast([]byte("Connected, ready to watch!"))
  })

  socket.HandleMessage(func(s *melody.Session, msg []byte) {
    args := strings.Split(string(msg), " ")

    if len(args) == 0 {
      return
    }

    reaction, ok := commands[args[0]]

    if ok {
      reaction(context, args)
    } else {
      commands["default"](context, args)
    }
  })

  watcher.watch()

  router.Run(":5000")
}
