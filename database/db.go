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
    createTable()
    return nil
}


func createTable() {
    userTable := `CREATE TABLE IF NOT EXISTS user (
        id INT AUTO_INCREMENT PRIMARY KEY,
        email VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL
    );`

    _, err := DB.Exec(userTable)
    if err != nil {
        panic("User Table Is Not Created: " + err.Error())
    }


    log.Println("User Table Is Created")
    
    eventTable := `CREATE TABLE IF NOT EXISTS events (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        description VARCHAR(255) NOT NULL,
        user_id INT,
        price INT NOT NULL,
        total_count INT DEFAULT 0, -- 'DEFAULT' instead of 'default'
        FOREIGN KEY (user_id) REFERENCES users(id) -- 'users' instead of 'user'
    );`
    
    _, err = DB.Exec(eventTable) 
    if err != nil {
        panic("Event Table Is Not Created: " + err.Error())
    }

    
    log.Println("Events Table Is Created")
}
