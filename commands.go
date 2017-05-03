package main

import (
  "strings"
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

func onAdd(ctx *Context, args []string) {
  if len(args) != 3 {
    say(ctx.Socket, "Error: Wrong arguments")
    return
  }

  fields := args[1:]
  file := rawFieldsToFile(fields)

  if file.Hash == "" || file.Path == "" {
    say(ctx.Socket, "Error: Wrong arguments")
    return
  }

  file.save(ctx.Database)

  say(ctx.Socket, "Watching: " + file.Hash + ":" + file.Path)
}

func onList(ctx *Context, _ []string) {
  response := ""
  files := allFiles(ctx.Database)

  for _, file := range files {
    response += file.Hash + ":" + file.Path + "\n"
  }

  say(ctx.Socket, response)
}

func onRemove(ctx *Context, args []string) {
  if len(args) < 2 {
    say(ctx.Socket, "Error: Not enough arguments")
    return
  }

  hash := strings.TrimSpace(args[1])
  files := filesByHash(ctx.Database, hash)

  if len(files) == 0 {
    say(ctx.Socket, "Error: No file with hash " + hash)
    return
  }

  file := &files[0]
  file.remove(ctx.Database)

  say(ctx.Socket, "Unwatched: " + file.Hash + ":" + file.Path)
}

func onReset(ctx *Context, _ []string) {
  clearFiles(ctx.Database)

  say(ctx.Socket, "Unwatched: All")
}

func onMd5(ctx *Context, args []string) {
  if len(args) < 2 {
    say(ctx.Socket, "Error: Not enough arguments")
    return
  }

  hash, err := md5By(args[1])

  if err != nil {
    say(ctx.Socket, "Error: " + err.Error())
  } else {
    say(ctx.Socket, "MD5: " + hash)
  }
}

func onStop(ctx *Context, _ []string) {
  ctx.Watcher.stop()
}

func onStart(ctx *Context, _ []string) {
  ctx.Watcher.start()
}

func onHelp(ctx *Context, _ []string) {
  say(ctx.Socket, `Commands:

  'watch hash=[MD5 checksum] path=[File absolute path]' Starts watching for the file

  'unwatch [MD5 checksum]' Stops watching for the file

  'list' Shows the list of files currently being watched

  'reset' Unwatches all the files

  'stop' Pauses the watcher

  'start' Starts the watcher

  'md5 [File absolute path]' Returns MD5 checksum by the absolute path
`)
}

func onDefault(ctx *Context, args []string) {
  say(ctx.Socket, "Error: Command " + args[0] + " not found, type `help` for info about commands")
}
