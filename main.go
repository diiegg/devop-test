package main
import (
  "database/sql"
  _"github.com/go-sql-driver/mysql"
)
func main() {
  db, err := sql.Open("mysql", "admin:password@tcp(127.0.0.1:3306)/devops")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()
  insert, err := db.Query("INSERT INTO posts(title) VALUES('My post')")
  if err != nil {
    panic(err.Error())
  }
  defer insert.Close()
}
