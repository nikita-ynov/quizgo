package main

import (
	"fmt"
	initTemp "goquiz/pages"
	"goquiz/router"
	"net/http"
)

func main() {
	initTemp.Init()

	r := router.New()

	// Sert les fichiers statiques (CSS, JS, images, etc.)
	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Serveur demarré sur http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
