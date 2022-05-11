// package repository

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// )

// func InitDb() *sql.DB {

// 	const (
// 		host = "localhost"
// 		// host     = "db"
// 		port     = "5432"
// 		user     = "postgres"
// 		password = "12345"
// 		dbname   = "postgres"
// 		sslMode  = "disable"
// 	)

// 	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
// 		"password=%s dbname=%s sslmode=%s",
// 		host, port, user, password, dbname, sslMode)

// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Println("Successfully connected!")

// 	return db
// }

package repository

import (
	"database/sql"

	//Driver for sqlite3
	_ "github.com/mattn/go-sqlite3"
)

//DbInit ..
func InitDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./forum.db?_foreign_keys=on")
	if err != nil {
		return nil, err
	}
	return db, nil
}
