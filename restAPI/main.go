package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Employee struct {
	Id        string `json: "id,omitempty"`
	FirstName string `json: "firstname, omitEmpty"`
	LastName  string `json: "lastname, omitEmpty"`
	Age       int    `json: "age, omitEmpty"`
}

var emps []Employee

func GetAllEmps(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(emps)
}

/*func GetEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for emp := range emps {
		if emp.Id == params["id"] {
			json.NewEncoder(w).Encode(emp)
			return
		}
	}
	json.NewEncoder(w).Encode(&emp)
}*/
func main() {
	router := mux.NewRouter()
	emps = append(emps, Employee{"ND074W", "Niharika", "Doppalapudi", 24})
	emps = append(emps, Employee{"VS015D", "Vinay", "Sampally", 40})
	router.HandleFunc("/emp", GetAllEmps).Methods("GET")
	router.HandleFunc("emp/{id}", GetEmployee).Methods("GET")
	//router.HandleFunc("/add", CreateEmployee).Methods("POST")
	//router.HandleFunc("/del", DeleteEmployee).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
