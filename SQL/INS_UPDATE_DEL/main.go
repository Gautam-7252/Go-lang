package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	Id        int
	Firstname string
	Lastname  string
}

var db *sql.DB
var tpl *template.Template

func main() {
	var err error
	tpl, err = template.ParseGlob("E:\\VScode\\Golang\\SQL\\templates\\*.html")
	if err != nil {
		fmt.Println("Error Parsing Template 1")
		panic(err)
	}
	pswd := os.Getenv("MYSQL_PASSWORD")
	db, err = sql.Open("mysql", "root:"+pswd+"@tcp(localhost:3306)/test2")
	if err != nil {
		fmt.Println("Error Connecting database")
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting the Database")
		panic(err)
	}
	defer db.Close()
	fmt.Println("Successfull Database Connection")
	http.HandleFunc("/search", searchhandler)
	http.HandleFunc("/insert", inserthandler)
	http.ListenAndServe("localhost:8000", nil)
}

func searchhandler(w http.ResponseWriter, r *http.Request) {
	stmt := "SELECT * FROM customer;"
	row, err := db.Query(stmt)
	if err != nil {
		fmt.Println("Error in Query")
		panic(err)
	}
	defer row.Close()
	var Cust []Customer
	for row.Next() {
		var E Customer
		err = row.Scan(&E.Id, &E.Firstname, &E.Lastname)
		if err != nil {
			panic(err)
		}
		Cust = append(Cust, E)
	}
	tpl.ExecuteTemplate(w, "index3.html", Cust)
}

func inserthandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "insert.html", nil)
		return
	}
	r.ParseForm()
	Custid := r.FormValue("CustId")
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	var ins *sql.Stmt
	var err error
	ins, err = db.Prepare("INSERT INTO `test2`.`customer` (`idcustomer`, `firstname`, `lastname`) VALUES ( ?, ?, ?);")
	fmt.Println("Inside Insert")
	if err != nil {
		fmt.Println("Error in Insert Statement")
		panic(err)
	}
	defer ins.Close()
	insert, err := ins.Exec(Custid, fname, lname)
	rowaffected, _ := insert.RowsAffected()
	if err != nil || rowaffected != 1 {
		fmt.Println("Error Inserting row", err)
		tpl.ExecuteTemplate(w, "insert.html", "Please fill all the fields.")
	}
	fmt.Println(fname, lname, Custid)
	fmt.Println("Successfully Inserted in the Database!")
	tpl.ExecuteTemplate(w, "insert.html", "Successfully Inserted")
}
