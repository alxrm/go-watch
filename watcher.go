package main

import (
  "time"
)

type Watcher struct {
  IntervalMillis int
  Database       *Database
  OnObserved     func(file *File, where string)
  OnStarted      func()
  OnStopped      func()
  quit           chan bool
}

func (w *Watcher) start() {
  if w.quit != nil {
    return
  }

  w.quit = make(chan bool)

  if w.OnStarted != nil {
    w.OnStarted()
  }

  go func() {
    delay := time.Duration(w.IntervalMillis) * time.Millisecond

    for {
      w.performCheck()

      select {
      case <-time.After(delay):
      case <-w.quit:
        return
      }
    }
  }()
}

func (w *Watcher) performCheck() {
  files := allFiles(w.Database)

  for _, file := range files {
    hash, err := md5By(file.Path)

    if err != nil {
      continue
    }

    if hash == file.Hash {
      w.OnObserved(&file, file.Path)
    }
  }
}

func (w *Watcher) stop() {
  if w.quit == nil {
    return
  }

  w.quit <- true
  w.quit = nil

  if w.OnStopped != nil {
    w.OnStopped()
  }
}
