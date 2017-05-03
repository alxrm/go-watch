package main

import (
  "github.com/gin-gonic/gin"
  _ "github.com/mattn/go-sqlite3"
  "gopkg.in/olahol/melody.v1"
  "log"
  "net/http"
  "strings"
)

//go:generate file2const -package main assets/index.html:indexHtml index_html.go

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
    Database:       db,
    OnObserved: func(file *File, path string) {
      say(socket, "Found: " + path)
    },
    OnStopped: func() {
      say(socket, "Stopped")
    },
  }

  context := &Context{
    Watcher:  watcher,
    Database: db,
    Socket:   socket,
  }

  router.GET("/", func(c *gin.Context) {
    http.ServeFile(c.Writer, c.Request, "./assets/index.html")
  })

  router.GET("/watcher", func(c *gin.Context) {
    socket.HandleRequest(c.Writer, c.Request)
  })

  socket.HandleConnect(func(s *melody.Session) {
    say(socket, "Connected, ready to watch!")
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

  watcher.start()
  router.Run(":5000")
}
