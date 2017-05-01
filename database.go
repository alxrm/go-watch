package main

import (
	"database/sql"
)

type Database struct {
	name   string
	client *sql.DB
}

func makeDB(file string, create string) (*Database, error) {
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

func (wrapper *Database) exec(command string, args ...interface{}) error {
	_, err := wrapper.client.Exec(command, args...)

	if err != nil {
		return err
	}

	return nil
}

func (wrapper *Database) query(command string, args []interface{}, fieldsOf func() []interface{}, entityOf func(fields []interface{}) interface{}) ([]interface{}, error) {
	rows, err := wrapper.client.Query(command, args...)

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

func (wrapper *Database) queryAll(command string, fieldsOf func() []interface{}, fieldsTo func(fields []interface{}) interface{}) ([]interface{}, error) {
	return wrapper.query(command, []interface{}{}, fieldsOf, fieldsTo)
}
