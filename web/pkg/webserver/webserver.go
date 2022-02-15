package webserver

import (
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/upload", HandleUpload)
	addr := ":5000"
	log.Printf("Listening on %s\n", addr)
	http.ListenAndServe(addr, nil)
}
