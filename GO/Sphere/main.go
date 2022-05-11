package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sphere/repository"
	"sphere/server"

	_ "github.com/lib/pq"
)

type Links struct {
	ActLink string `json:"active_link"`
	HisLink string `json:"history_link"`
}

func main() {
	db, _ := repository.InitDb()

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("close database: %v\n", err)
		}
	}(db)

	if err := repository.CreateTables(db); err != nil {
		log.Fatal(err)
		return
	}

	handle := server.CreateHandle(db)

	data := []Links{}
	body, _ := ioutil.ReadFile("links.json")
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal("Can not unmarshal JSON")
	}
	for i := 1; i < len(data); i++ {
		handle.DB.Exec(`INSERT INTO links (activlink, historylink)
		VALUES ($1,$2)`, data[i].ActLink, data[i].HisLink)
	}

	fmt.Println(data[1].ActLink)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/redirects", handle.AllLinks)
	mux.HandleFunc("/admin/redirects/", handle.OneLink)
	log.Println("Server is Listening..." + "\n" + "http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
