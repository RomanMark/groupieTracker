package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "Error 404: PAGE NOT FOUND.", http.StatusBadRequest)
		return
	}

	if renderTmpl(w, "templates/index.html", API.Artists) != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
	}

}

func artistPage(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	pageId := r.FormValue("id")

	if pageId == "" {
		path := strings.Split(r.URL.Path, "/")
		pageId = path[2]
	}
	index, err := strconv.Atoi(pageId)
	if err != nil {
		http.Error(w, "Error 404: PAGE NOT FOUND.", http.StatusBadRequest)
		return
	}

	if renderTmpl(w, "templates/artist.html", API.Artists[index-1]) != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
	}
}

func renderTmpl(w http.ResponseWriter, tmplName string, data any) error {

	tmpl, err := template.ParseFiles(tmplName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
