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
	ID        int
	Firstname string
	Lastname  string
}

var tpl *template.Template
var db *sql.DB

func main() {
	var err error
	tpl, err = template.ParseGlob("E:\\VScode\\GoLang\\SQL\\templates\\index.html")
	if err != nil {
		panic(err)
	}
	pswd := os.Getenv("MYSQL_PASSWORD")
	db, err = sql.Open("mysql", "root:"+pswd+"@tcp(localhost:3306)/test2")
	if err != nil {
		fmt.Println("Error Connecting database")
		panic(err.Error())
	}
	fmt.Println("SQL Connection Established!")
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Error during db.Ping")
		panic(err.Error())
	}
	fmt.Println("Successfull connection with Database!")

	http.HandleFunc("/search", searchhandler)
	http.ListenAndServe("localhost:8080", nil)
}
func searchhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "index.html", nil)
		return
	}
	r.ParseForm()
	var E Employee
	firstname := r.FormValue("EmployeeName")
	fmt.Println(firstname)
	stmt := "SELECT * FROM customer WHERE firstname = ?;"
	row := db.QueryRow(stmt, firstname)
	err := row.Scan(&E.ID, &E.Firstname, &E.Lastname)
	if err != nil {
		panic(err)
	}

	tpl.ExecuteTemplate(w, "index.html", E)
}
