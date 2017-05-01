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
		id:   *(fields[0].(*int)),
		hash: *(fields[1].(*string)),
		path: *(fields[2].(*string)),
	}
}

func fileToRaw(file *File) []interface{} {
	var raw = make([]interface{}, 3)

	raw[0] = file.id
	raw[1] = file.hash
	raw[2] = file.path

	return raw
}

func fileToFields() []interface{} {
	file := File{}
	fields := make([]interface{}, 3)

	fields[0] = &file.id
	fields[1] = &file.hash
	fields[2] = &file.path

	return fields
}
