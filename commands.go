package main

import (
  "strings"
  "log"
  "encoding/json"
  "gopkg.in/olahol/melody.v1"
)

var commands = map[string]func(*Context, []string){
  "watch":   onAdd,
  "list":    onList,
  "unwatch": onRemove,
  "reset":   onReset,
  "md5":     onMd5,
  "stop":    onStop,
  "start":   onStart,
  "help":    onHelp,
  "default": onDefault,
}

type Reply struct {
  Subject string      `json:"subject"`
  Error   string      `json:"error,omitempty"`
  Result  interface{} `json:"result,omitempty"`
}

func onAdd(ctx *Context, args []string) {
  if len(args) != 3 {
    say(ctx.Socket, Reply{Subject:args[0], Error:"Wrong arguments"})
    return
  }

  fields := args[1:]
  file := rawFieldsToFile(fields)

  if file.Hash == "" || file.Path == "" {
    say(ctx.Socket, Reply{Subject:args[0], Error:"Wrong arguments"})
    return
  }

  if err := file.save(ctx.Database); err != nil {
    if strings.HasPrefix(err.Error(), "UNIQUE") {
      say(ctx.Socket, Reply{
        Subject:args[0],
        Error:"The file with this hash is already being watched",
      })
    } else {
      say(ctx.Socket, Reply{
        Subject:args[0],
        Error:err.Error(),
      })
    }

    return
  }

  say(ctx.Socket, Reply{Subject:args[0], Result:file})
}

func onList(ctx *Context, args []string) {
  files := allFiles(ctx.Database)

  say(ctx.Socket, Reply{Subject:args[0], Result:files})
}

func onRemove(ctx *Context, args []string) {
  if len(args) != 2 {
    say(ctx.Socket, Reply{Subject:args[0], Error:"Wrong arguments"})
    return
  }

  hash := strings.TrimSpace(args[1])

  if !validateHash(hash) {
    say(ctx.Socket, Reply{Subject:args[0], Error:"Wrong arguments"})
    return
  }

  files := filesByHash(ctx.Database, hash)

  if len(files) == 0 {
    say(ctx.Socket, Reply{Subject:args[0], Error:"No file with hash " + hash})
    return
  }

  fileToRemove := &files[0]

  if err := fileToRemove.remove(ctx.Database); err != nil {
    say(ctx.Socket, Reply{Subject:args[0], Error:err.Error()})
    return
  }

  say(ctx.Socket, Reply{Subject:args[0], Result:fileToRemove})
}

func onReset(ctx *Context, args []string) {
  all := allFiles(ctx.Database)

  clearFiles(ctx.Database)

  say(ctx.Socket, Reply{Subject: args[0], Result:all})
}

func onMd5(ctx *Context, args []string) {
  if len(args) < 2 {
    say(ctx.Socket, Reply{Subject:args[0], Error:"Not enough arguments"})
    return
  }

  hash, err := md5By(args[1])

  if err != nil {
    say(ctx.Socket, Reply{Subject:args[0], Error:err.Error()})
  } else {
    say(ctx.Socket, Reply{Subject:args[0], Result:hash})
  }
}

func onStop(ctx *Context, _ []string) {
  ctx.Watcher.stop()
}

func onStart(ctx *Context, _ []string) {
  ctx.Watcher.start()
}

func onHelp(ctx *Context, args []string) {
  say(ctx.Socket, Reply{Subject: args[0], Result:[]string{
    "'watch [MD5 checksum] [File absolute path]' Starts watching for the file",
    "'unwatch [MD5 checksum]' Stops watching for the file",
    "'list' Shows the list of files currently being watched",
    "'reset' Unwatches all the files",
    "'stop' Pauses the watcher",
    "'start' Starts the watcher",
    "'md5 [File absolute path]' Returns MD5 checksum by the absolute path",
  }})
}

func onDefault(ctx *Context, args []string) {
  say(ctx.Socket, Reply{
    Subject:"default",
    Error:"Command '" + args[0] + "' not found, use 'help' for info about commands",
  })
}

func say(sock *melody.Melody, reply Reply) {
  data, err := json.Marshal(reply)

  if err != nil {
    log.Fatalf("Could not marshal to json this: %#v", reply)
    return
  }

  sock.Broadcast(data)
}