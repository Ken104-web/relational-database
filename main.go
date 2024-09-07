package main

import (
	"database/sql"
	"fmt"
	"log"
    "os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main()  {
    // capture connection properties

    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "recordings",
    }
    // Aquire a database handle
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil{
        log.Fatal(err)
    }
    // to confirm the connection of the database works
    pingErr := db.Ping()
    if pingErr != nil{
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
    
}
