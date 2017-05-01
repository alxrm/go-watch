package main

import (
  "time"
  "log"
)

type Watcher struct {
  IntervalMillis int
  RootDir        string
  Database       *Database
  OnObserved     func(file *File, where string)
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
  log.Println("Checking...")
  files := allFiles(w.Database)

  for _, file := range files {
    fullPath := fullPathBy(w.RootDir, file.path)
    hash, err := md5sum(fullPath)

    if err != nil {
      continue
    }

    if hash == file.hash {
      w.OnObserved(&file, fullPath)
    }
  }
}

func (w *Watcher) stop() {
  close(w.quit)
  log.Println("Stopped")
}
