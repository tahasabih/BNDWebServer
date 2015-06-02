package pgConnector

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"io/ioutil"
	"log"
)

func Connect() (*sql.DB, error) {
	connectionString, errRead := readFile()
	if errRead != nil {
		return nil, errRead
	}
	db, errdb := sql.Open("postgres", connectionString)
	if errdb != nil {
		log.Fatal(errdb)
		return db, errdb
	}
	_, err := db.Query("SELECT common_name FROM \"public\".\"USERS\"") //check for access to provide test
	if err, ok := err.(*pq.Error); ok {
		fmt.Println("pq error:", err.Code.Name())
		return db, err
	} else {
		return db, nil
	}
	defer db.Close()
	return db, nil
}

func readFile() (string, error) {
	fileIn, readErr := ioutil.ReadFile("connString.txt") //conn string in .gitignore contain username and pw in plain text
	if readErr != nil {
		log.Fatal(readErr)
		return "not found", readErr
	} else {
		return string(fileIn), nil
	}
}
