package main

import (
	"flag"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
	"os"

	"github.com/subosito/gotenv"
)

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

func main ()  {
	gotenv.Load()
	dsn := flag.String("dsn", os.Getenv("DB_DSN"), "database dsn")
	flag.Parse()

	pool, err := sql.Open("postgres", *dsn)
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}
	defer pool.Close()

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	rows, err := pool.Query("select first_name, age from person limit 1")
	if err != nil {
		log.Fatal("Error getting rows", err)
	}

	person := Person{}

	if rows.Next() {
		rows.Scan(person)
	}
	fmt.Println(person)
}
