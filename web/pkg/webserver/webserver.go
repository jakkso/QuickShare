package webserver

import (
	"fmt"
	"log"
	"net/http"
)

func Homepage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Homepage\n")
}

func Run() {
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/upload", HandleUpload)
	addr := ":5000"
	log.Printf("Listening on %s\n", addr)
	http.ListenAndServe(addr, nil)
}
