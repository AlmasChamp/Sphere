package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var links struct {
	Redirect []link
}

type link struct {
	ActLink string `json:"active_link"`
	HisLink string `json:"history_link"`
}

func (handle *Handle) AllLinks(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/links.html")
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	if r.Method == http.MethodGet {
		body, _ := ioutil.ReadFile("links.json")
		err = json.Unmarshal(body, &links.Redirect)
		if err != nil {
			log.Fatal("Can not unmarshal JSON")
		}
		t.Execute(w, links)
	} else if r.Method == http.MethodPost {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			t.ExecuteTemplate(w, "errors.html", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		infoLink := &link{}

		err = json.Unmarshal(data, infoLink)
		if err != nil {
			log.Println(err)
			t.ExecuteTemplate(w, "errors.html", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println(infoLink.ActLink, infoLink.HisLink)
		_, err = handle.DB.Exec(`INSERT INTO links (activlink, historylink)
		VALUES ($1,$2)`, infoLink.ActLink, infoLink.HisLink)

		if err != nil {
			log.Println(err)
			return
		}

	}

}
