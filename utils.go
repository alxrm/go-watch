package main

import (
  "crypto/md5"
  "encoding/hex"
  "io"
  "io/ioutil"
  "os"
  "log"
  "path"
  "strings"
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

func fullPathBy(rootDir, filePath string) string {
  if path.IsAbs(filePath) {
    return filePath
  } else {
    if strings.HasSuffix(rootDir, "/") {
      rootDir = rootDir[:len(rootDir) - 1]
    }

    return rootDir + filePath[1:] // because we don't need the dot like here: `./`
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
