package pgConnector

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"io/ioutil"
	"log"
)

func Connect() (bool, error) {
	fileIn, readErr := ioutil.ReadFile("connString.txt") //conn string in .gitignore contain username and pw in plain text
	if readErr != nil {
		log.Fatal(readErr)
		return false, readErr
	}
	connectionString := string(fileIn)
	db, errdb := sql.Open("postgres", connectionString)
	if errdb != nil {
		log.Fatal(errdb)
		return false, errdb
	}
	_, err := db.Query("SELECT common_name FROM \"public\".\"USERS\"") //check for access to provide test
	if err, ok := err.(*pq.Error); ok {
		fmt.Println("pq error:", err.Code.Name())
		return false, err
	} else {
		return true, nil
	}
	defer db.Close()
	return false, nil
}
