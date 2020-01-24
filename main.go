package main

import (
	"fmt"
	"os"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
	"log"
	"math/rand"
	"time"
	"strconv"
	"context"
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
func getPerson(c chan Person, db *sqlx.DB) {
	person := Person{}
	s1 := rand.NewSource(time.Now().UnixNano())
    rand := rand.New(s1)
	offset := rand.Intn(10)
	db.Get(&person, "SELECT * FROM person ORDER BY first_name ASC OFFSET " + strconv.Itoa(offset))
	c <- person
}


func main() {
	gotenv.Load()
	pgdsn := os.Getenv(dsn)
	rand.Seed(86)
	db, err := sqlx.Connect("postgres", pgdsn)

	if err != nil {
        log.Fatalln(err)
	}
	db.MustExec(schema)
	db.SetMaxOpenConns(3)
	tx := db.MustBegin()
    tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
    tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
    tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Frank", "Tompson", "tompson@gmail.net")
    tx.Commit()

	c := make(chan Person)
	for i:=0; i < 100; i++ {
		go getPerson(c, db)
	}
	ctx := context.WithValue(context.Background(), "key", "Go")
	fmt.Printf("%+v %+v\n", db.Stats(), db.DB.PingContext(ctx))
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Print(<-c)
}
