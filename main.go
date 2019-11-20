package main

import (
	"fmt"
	"os"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
	"log"
	"github.com/jmoiron/sqlx"
)

const (
	dsn = "SQLX_POSTGRES_DSN"
)

var schema = `
CREATE TABLE IF NOT EXISTS person (
    first_name text,
    last_name text,
    email text
);

CREATE TABLE IF NOT EXISTS place (
    country text,
    city text NULL,
    telcode integer
)`

type Person struct {
    FirstName string `db:"first_name"`
    LastName  string `db:"last_name"`
    Email     string
}

func main() {
	gotenv.Load()
	pgdsn := os.Getenv(dsn)
	db, err := sqlx.Connect("postgres", pgdsn)

	if err != nil {
        log.Fatalln(err)
	}
	db.MustExec(schema)
	
	tx := db.MustBegin()
    tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
    tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
    tx.Commit()

	people := []Person{}
	db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
	fmt.Println(people)
}
