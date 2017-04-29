package main

import (
  "os"
)

func moveFile(fromPath, toPath string) error {
  err := os.Rename(fromPath, toPath)

  return err
}