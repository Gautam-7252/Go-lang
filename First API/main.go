package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type todo struct {
	ID        string `json:"id"`
	Work      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{"1", "Build API", false},
	{"2", "Build API", false},
	{"3", "Build API", false},
}

func main() {
	fmt.Println("BUILDING AN API IN GOLANG")
	r := mux.NewRouter()
	r.HandleFunc("/get", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	Finaljson, _ := json.Marshal(todos)
	if json.Valid(Finaljson) {
		fmt.Println("VALID JSON")
		w.Write(Finaljson)
	} else {
		fmt.Println("JSON NOT VALID")
	}
}
