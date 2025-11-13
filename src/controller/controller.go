package controller

import (
	"goquiz/pages"
	"net/http"
)

func renderPage(w http.ResponseWriter, filename string, data any) {
	err := pages.Temp.ExecuteTemplate(w, filename, data)
	if err != nil {
		http.Error(w, "Erreur rendu template : "+err.Error(), http.StatusInternalServerError)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Message": "Controller Home !",
	}
	renderPage(w, "index.html", data)
}
