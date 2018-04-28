package db

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func OpenDatabaseConnection() {
   fmt.Print("Initiate database")

   db, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3306)/testdb")

   if err != nil {
     panic(err.Error( ))
   }

   defer db.Close()

   fmt.Print("Successfully connected to the database")
}