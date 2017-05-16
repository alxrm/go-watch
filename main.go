package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/olahol/melody.v1"
	"log"
	"strings"
)

type Context struct {
	Watcher  *Watcher
	Database *Database
	Socket   *melody.Melody
}

//go:generate file2const -package main assets/index.html:indexHtml index_html.go

func main() {
	db, err := newDB(databaseFile, createStatement)

	if err != nil {
		log.Fatal(err)
		return
	}

	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	socket := melody.New()

	watcher := &Watcher{
		IntervalMillis: 1000,
		Database:       db,
		OnObserved: func(file *File, path string) {
			file.remove(db)
			say(socket, Reply{Subject: "found", Result: path})
		},
		OnStarted: func() {
			say(socket, Reply{Subject: "started"})
		},
		OnStopped: func() {
			say(socket, Reply{Subject: "stopped"})
		},
	}

	context := &Context{
		Watcher:  watcher,
		Database: db,
		Socket:   socket,
	}

	router.GET("/", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/html")
		c.String(200, indexHtml)
	})

	router.GET("/watcher", func(c *gin.Context) {
		socket.HandleRequest(c.Writer, c.Request)
	})

	socket.HandleConnect(func(s *melody.Session) {
		say(socket, Reply{Subject: "connected"})
	})

	socket.HandleMessage(func(s *melody.Session, msg []byte) {
		args := strings.Split(strings.TrimSpace(string(msg)), " ")

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

	router.Run(":5000")
}
