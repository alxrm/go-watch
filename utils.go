package main

import (
  "fmt"
  "io/ioutil"
)

func printDirContents(path string) {
  files, _ := ioutil.ReadDir(path)

  for _, f := range files {
    fmt.Println(f.Name())
  }
}
