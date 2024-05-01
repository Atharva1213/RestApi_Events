package database

import (
    "database/sql"
    "log"
	"os"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectToDB() error {
    // Database connection parameters

	DRIVER:=os.Getenv("DRIVERNAME")
	SQLPATH:=os.Getenv("SQLPATH")

    var err error
    DB, err = sql.Open(DRIVER,SQLPATH)
    if err != nil {
        return err
    }

    // Check if the connection is successful
    if err = DB.Ping(); err != nil {
        DB.Close()
        return err
    }

    log.Println("Connected to MySQL database!")
    return nil
}
