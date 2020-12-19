package database

import (
	"database/sql"
	"fmt"

	// Blank import of postgres driver necessary to access postgres database
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Storeable interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	Prepare(query string) (*sql.Stmt, error)
	Begin() (*sql.Tx, error)
}

// Store represents a persistent store
type Store struct {
	*sql.DB
}

var db *sql.DB

var databaseURL string

func init() {
	databaseURL = fmt.Sprintf("%s?sslmode=disable", os.Getenv("DATABASE_URL"))
	var err error
	db, err = sql.Open("postgres", databaseURL)
	if err != nil {
		log.Println("Failed to open connection to postgres", err)
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
}

// NewStore returns a persistent store object. A transaction has automatically been
// started so make sure you call the Close() method to close the database
func NewStore() Store {
	return Store{db}
}

func NewTransaction() (*sql.Tx, error) {
	return db.Begin()
}
