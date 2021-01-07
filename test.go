package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Salary struct {
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

var salaries struct {
	Data []Salary `json:"data"`
}

func employeeSalaries(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	db, err := sql.Open("mysql", "senior:erplysdv@tcp(127.0.0.1:3306)/devops?charset=utf8")
	checkErr(err)

	// query
	employees, err := db.Query("SELECT emp_no, first_name FROM employees LIMIT ?,20", rand.Intn(100000))
	checkErr(err)

	for employees.Next() {
		var empNo string
		var firstName string
		err = employees.Scan(&empNo, &firstName)
		checkErr(err)
		empSalaries, err := db.Query("SELECT max(salary) FROM salaries WHERE emp_no =" + empNo)

		for empSalaries.Next() {
			var empSalary int
			err = empSalaries.Scan(&empSalary)
			checkErr(err)
			salaries.Data = append(salaries.Data, Salary{Name: firstName, Salary: empSalary})
		}
	}

	data, err := json.Marshal(salaries.Data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, string(data))
	elapsed := time.Since(start)
	fmt.Println("Date:", time.Now(), "Response time:", elapsed)
	salaries.Data = nil
	data = nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", employeeSalaries)   // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
