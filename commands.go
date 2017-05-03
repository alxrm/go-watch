package main

import (
  "strings"
)

var commands = map[string]func(*Context, []string){
  "watch": onAdd,
  "list": onList,
  "unwatch": onRemove,
  "reset": onReset,
  "md5": onMd5,
  "default": onDefault,
}

func onAdd(ctx *Context, args []string) {
  if len(args) < 3 {
    ctx.Socket.Broadcast([]byte("Error: Not enough arguments"))
    return
  }

  fields := args[1:]

  file := rawFieldsToFile(fields)
  file.save(ctx.Database)

  ctx.Socket.Broadcast([]byte("Watching: " + file.Hash + ":" + file.Path))
}

func onList(ctx *Context, _ []string) {
  response := ""
  files := allFiles(ctx.Database)

  for _, file := range files {
    response += file.Hash + ":" + file.Path + "\n"
  }

  ctx.Socket.Broadcast([]byte(response))
}

func onRemove(ctx *Context, args []string) {
  if len(args) < 2 {
    ctx.Socket.Broadcast([]byte("Error: Not enough arguments"))
    return
  }

  hash := strings.TrimSpace(args[1])
  files := filesByHash(ctx.Database, hash)

  if len(files) == 0 {
    ctx.Socket.Broadcast([]byte("Error: No file with hash " + hash))
    return
  }

  file := &files[0]
  file.remove(ctx.Database)

  ctx.Socket.Broadcast([]byte("Unwatched: " + file.Hash + ":" + file.Path))
}

func onReset(ctx *Context, _ []string) {
  clearFiles(ctx.Database)

  ctx.Socket.Broadcast([]byte("Unwatched: All"))
}

func onMd5(ctx *Context, args []string) {
  if len(args) < 2 {
    ctx.Socket.Broadcast([]byte("Error: Not enough arguments"))
    return
  }

  hash, err := md5By(args[1])

  if err != nil {
    ctx.Socket.Broadcast([]byte("Error: " + err.Error()))
  } else {
    ctx.Socket.Broadcast([]byte("MD5: " + hash))
  }
}

func onDefault(ctx *Context, args []string) {
  ctx.Socket.Broadcast([]byte("Error: Wrong command " + args[0]))
}