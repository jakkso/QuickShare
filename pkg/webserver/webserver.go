package webserver

import (
	"fmt"
	"net/http"
)

func Homepage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Homepage\n")
}

func Run() {
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/upload", HandleUpload)
	addr := GetAddress()
	fmt.Printf("Listening on %s\n", addr)
	http.ListenAndServe(addr, nil)
}
