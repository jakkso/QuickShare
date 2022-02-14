package webserver

import (
	"log"
	"os"
)

func GetAddress() string {
	host := os.Getenv("QS_HOST")
	if host == "" {
		log.Println("Warning: QS_HOST not set, defaulting to ':'")
	}
	port := os.Getenv("QS_PORT")
	if port == "" {
		log.Println("Warning: QS_PORT not set, defaulting to 5000")
		port = "5000"
	}
	addr := host+":"+port
	return addr
}
