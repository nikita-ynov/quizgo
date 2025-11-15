package router

import (
	"goquiz/controller"
	"net/http"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/quiz", controller.Quiz)
	mux.HandleFunc("/score", controller.Score)
	return mux
}
