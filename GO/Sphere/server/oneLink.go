package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var myLink struct {
	Redirect []oneLink
}

type oneLink struct {
	ActLink string `json:"active_link"`
	HisLink string `json:"history_link"`
}

func (handle *Handle) OneLink(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/links.html")
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	if r.Method == http.MethodGet {
		id, _ := strconv.Atoi(r.RequestURI[17:])
		fmt.Println("here", id)
		userId := oneLink{}
		rows, _ := handle.DB.Query("SELECT activlink,historylink FROM links WHERE id= $1", id)

		for rows.Next() {
			rows.Scan(
				&userId.ActLink,
				&userId.HisLink,
			)
			myLink.Redirect = append(myLink.Redirect, userId)
		}
		t.Execute(w, myLink)
		return
	} else if r.Method == http.MethodPatch {
		id, _ := strconv.Atoi(r.RequestURI[17:])
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			t.ExecuteTemplate(w, "errors.html", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		infoLink := &oneLink{}

		err = json.Unmarshal(data, &infoLink)
		if err != nil {
			log.Println(err)
			t.ExecuteTemplate(w, "errors.html", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = handle.DB.Exec(`UPDATE links 
		SET activlink = $2,
		historylink = $3
		WHERE id =$1`, id, infoLink.ActLink, infoLink.HisLink)
		if err != nil {
			log.Println(err)
			return
		}
		return
	} else if r.Method == http.MethodDelete {
		fmt.Println("here method delete")
		id, _ := strconv.Atoi(r.RequestURI[17:])
		fmt.Println(id)
		_, err = handle.DB.Exec(`DELETE FROM links 
		WHERE id =$1`, id)
		return
	}

}
