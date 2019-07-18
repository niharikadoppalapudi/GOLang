package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "test123"
	dbname   = "testdb"
)

func main() {
	/*router := mux.NewRouter()
	getMsg := func(res http.ResponseWriter, req *http.Request) {
		//if req.Method == "GET" {
		io.WriteString(res, "GET API")
		/*} else if req.Method == "POST" {
			io.WriteString(res, "POST API")
		}*/

	/*getName := func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "Hi niha!")
	}*/

	//router.HandleFunc("/getmsg", getMsg).Methods("GET")
	//router.HandleFunc("/name", getName).Methods("POST")
	//http.ListenAndServe(":9000", router)

	//psqlInfo := fmt.Sprintf("host=%s port=%d username=%s"+"password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgreSQL", psqlInfo)

	fmt.Println(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected successfully")
}
