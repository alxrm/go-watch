package main

import (
  "time"
  "log"
)

type Watcher struct {
  IntervalMillis int
  RootDir        string
  Database       *Database
  OnObserved     func(file File)
  OnCheck        func(hash string) bool
  quit           chan bool
}

func (watcher *Watcher) watch() {
  watcher.quit = make(chan bool)

  go func() {
    delay := time.Duration(watcher.IntervalMillis) * time.Millisecond

    for {
      watcher.performCheck()

      select {
      case <-time.After(delay):
      case <-watcher.quit:
        return
      }
    }
  }()
}

func (watcher *Watcher) performCheck() {
  watcher.OnCheck("")
}

func (watcher *Watcher) stop() {
  close(watcher.quit)
  log.Println("Stopped")
}
