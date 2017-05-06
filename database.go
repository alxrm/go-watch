package main

import (
  "database/sql"
)

type Database struct {
  name   string
  client *sql.DB
}

func newDB(file string, create string) (*Database, error) {
  db, errOpen := sql.Open("sqlite3", file)

  if errOpen != nil {
    return nil, errOpen
  }

  _, errCreate := db.Exec(create)

  if errCreate != nil {
    return nil, errCreate
  }

  return &Database{name: file, client: db}, nil
}

func (d *Database) exec(command string, args ...interface{}) error {
  _, err := d.client.Exec(command, args...)

  if err != nil {
    return err
  }

  return nil
}

func (d *Database) query(command string, args []interface{}, fieldsOf func() []interface{},
  entityOf func(fields []interface{}) interface{}) ([]interface{}, error) {

  rows, err := d.client.Query(command, args...)

  if err != nil {
    return nil, err
  }

  defer rows.Close()

  result := []interface{}{}

  for rows.Next() {
    fields := fieldsOf()

    if err = rows.Scan(fields...); err != nil {
      return nil, err
    }

    result = append(result, entityOf(fields))
  }

  return result, nil
}

func (d *Database) queryAll(command string, fieldsOf func() []interface{}, fieldsTo func(fields []interface{}) interface{}) ([]interface{}, error) {
  return d.query(command, []interface{}{}, fieldsOf, fieldsTo)
}
