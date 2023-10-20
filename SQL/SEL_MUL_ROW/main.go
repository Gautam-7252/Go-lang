package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id        int
	Firstname string
	Lastname  string
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
	defer db.Close()
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
		return
	}
	r.ParseForm()
	MinId := r.FormValue("MinId")
	MaxId := r.FormValue("MaxId")
	stmt := "SELECT * FROM customer WHERE idcustomer <= ? AND idcustomer >= ?;"
	row, err := db.Query(stmt, MaxId, MinId)
	if err != nil {
		fmt.Println("Error in Query")
		panic(err)
	}
	defer row.Close()
	var Emp []Employee
	for row.Next() {
		var E Employee
		err = row.Scan(&E.Id, &E.Firstname, &E.Lastname)
		if err != nil {
			panic(err)
		}
		Emp = append(Emp, E)
	}
	tpl.ExecuteTemplate(w, "index2.html", Emp)
}
