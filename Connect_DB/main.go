package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	pswd := os.Getenv("MYSQL_PASSWORD")
	db, err := sql.Open("mysql", "root:"+pswd+"@(127.0.0.1:3306)/test2")
	if err != nil {
		fmt.Println("Error validating SQL.Open arguements")
		panic(err.Error())
	}
	fmt.Println("SQL Connection Established")
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}
	insert, err := db.Query("INSERT INTO `test2`.`customer`(`idcustomer`,`firstname`,`lastname`) VALUES ('6','shawn','marsh');")
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
	fmt.Println("Successfull connection with Database!")
}
