package main

type Watcher struct {
  interval int
  directory string
  observed func(file File)
  check func(hash string)
}

func (watcher *Watcher) start() {

}
