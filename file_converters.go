package main

func toFiles(raw []interface{}) []File {
  res := make([]File, len(raw))

  for i, file := range raw {
    res[i] = file.(File)
  }

  return res
}

func fieldsToFile(fields []interface{}) interface{} {
  return File{
    hash: *(fields[0].(*string)),
    path: *(fields[1].(*string)),
  }
}

func fileToRaw(file *File) []interface{} {
  var raw = make([]interface{}, 2)

  raw[0] = file.hash
  raw[1] = file.path

  return raw
}

func fileToFields() []interface{} {
  file := File{}
  fields := make([]interface{}, 2)

  fields[0] = &file.hash
  fields[1] = &file.path

  return fields
}
