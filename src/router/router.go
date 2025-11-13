package router

import (
	"goquiz/controller"
	"net/http"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.Home)
	mux.HandleFunc("/quiz", controller.Quiz)
	return mux
}
