package main

import (
  "crypto/md5"
  "encoding/hex"
  "io"
  "io/ioutil"
  "os"
  "log"
)

func printDirContents(path string) {
  files, err := ioutil.ReadDir(path)

  if err != nil {
    log.Fatal(err)
  }

  for _, f := range files {
    log.Println(f.Name())
  }
}

func md5sum(filePath string) (string, error) {
  file, errOpen := os.Open(filePath)

  var result string

  if errOpen != nil {
    return result, errOpen
  }

  defer file.Close()

  hash := md5.New()
  _, errCopy := io.Copy(hash, file)

  if errCopy != nil {
    return result, errCopy
  }

  result = hex.EncodeToString(hash.Sum(nil))

  return result, nil
}
