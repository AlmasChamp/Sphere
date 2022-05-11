package repository

import (
	"database/sql"
	"log"
)

func CreateTables(db *sql.DB) error {

	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS links (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,        
		activlink VARCHAR NOT NULL,       
		historylink VARCHAR NOT NULL
	  );
	`)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Table Successfully Create!")
	return nil
}
