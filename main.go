package main

import (
	"bytes"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "Matt"
	password = "123"
	dbname   = "gnema"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user,
		password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	resetDB(db, dbname)
	if err != nil {
		panic(err)
	}
	db.Close()

	createPhoneNumbersTable(db)
}

func createPhoneNumbersTable(db *sql.DB) error {
	statement := `
	  CREATE TABLE IF NOT EXISTS phone_numbers (
		id SERIAL, 
		value VARCHAR(255)
	)`
	_, err := db.Exec(statement)
	return err
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return createDB(db, name)
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}
	return nil
}

func normalize(phone string) string {
	var buf bytes.Buffer
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}
