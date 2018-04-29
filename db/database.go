package db

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func OpenDatabaseConnection() {
   fmt.Println("Initiate database")

   db, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3306)/apidb")

   if err != nil {
     panic(err.Error())
   }

   defer db.Close()

   fmt.Println("Successfully connected to the database")

   insert, err := db.Query("INSERT INTO users VALUES('Elliot', 'eliot@hacker.com')")

   if err != nil {
     panic(err.Error())
   }

   defer insert.Close()

   fmt.Println("Successfully inserted values in database")
}
