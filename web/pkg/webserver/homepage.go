package webserver

import (
	"fmt"
	"net/http"
)

func Homepage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Homepage\n")
}
