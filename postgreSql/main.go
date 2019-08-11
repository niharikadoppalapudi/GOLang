package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "niharikadoppalapudi"
	password = ""
	dbname   = "demo"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port =%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Printf("Coonection unsuccess!")
	}
	rows, err := db.Query("SELECT * FROM demo")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Println(rows)
}
