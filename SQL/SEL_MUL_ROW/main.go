package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var Employee struct {
	Id        int
	Firstname string
	lastname  string
}
var db *sql.DB
var tpl *template.Template

func main() {
	var err error
	tpl, err = tpl.ParseGlob("E:\\VScode\\Golang\\SQL\\templates\\index2.html")
	if err != nil {
		fmt.Println("Error Parsing Template")
		panic(err)
	}
	pswd := os.Getenv("MYSQL_PASSWORD")
	db, err = sql.Open("mysql", "root:"+pswd+"@tcp(localhost:3306)/test2")
	if err != nil {
		fmt.Println("Error connecting Database")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Error in db.Ping")
		panic(err)
	}
	fmt.Println("Successfull connection with the database")
	http.HandleFunc("/search", search)
	http.ListenAndServe("localhost:8000", nil)
}
func search(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "index2.html", nil)
	}
}
