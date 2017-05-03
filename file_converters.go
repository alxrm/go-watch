package main

import "strings"

func toFiles(raw []interface{}) []File {
  res := make([]File, len(raw))

  for i, file := range raw {
    res[i] = file.(File)
  }

  return res
}

func fieldsToFile(fields []interface{}) interface{} {
  return File{
    Hash: *(fields[0].(*string)),
    Path: *(fields[1].(*string)),
  }
}

func rawFieldsToFile(fields []string) *File {
  file := &File{}

  for _, fd := range fields {
    arg := strings.Split(fd, "=")

    if len(arg) < 2 {
      continue
    }

    key := strings.ToLower(arg[0])
    value := strings.TrimSpace(arg[1])

    if key == "hash" {
      file.Hash = value
    } else if key == "path" {
      file.Path = value
    }
  }

  return file
}

func fileToRaw(file *File) []interface{} {
  var raw = make([]interface{}, 2)

  raw[0] = file.Hash
  raw[1] = file.Path

  return raw
}

func fileToFields() []interface{} {
  file := File{}
  fields := make([]interface{}, 2)

  fields[0] = &file.Hash
  fields[1] = &file.Path

  return fields
}
