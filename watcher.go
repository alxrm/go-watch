package main

import (
  "time"
)

type Watcher struct {
  IntervalMillis int
  RootDir        string
  Database       *Database
  OnObserved     func(file *File, where string)
  OnStopped      func()
  quit           chan bool
}

func (w *Watcher) watch() {
  w.quit = make(chan bool)

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
    fullPath := fullPathBy(w.RootDir, file.Path)
    hash, err := md5By(fullPath)

    if err != nil {
      continue
    }

    if hash == file.Hash {
      w.OnObserved(&file, fullPath)
    }
  }
}

func (w *Watcher) stop() {
  close(w.quit)

  if w.OnStopped != nil {
    w.OnStopped()
  }
}
