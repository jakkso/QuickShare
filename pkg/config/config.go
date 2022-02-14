package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var UploadDir string
var ShareDir string
var MimeTypes map[string]bool

type mimetypeImport struct {
	Allowed []string	`json:"allowedMimeTypes"`
}

func init() {
	loadDir("QS_UPLOAD_DIR", &UploadDir)
	loadDir("QS_SHARE_DIR", &ShareDir)
	loadMimeTypes()
}

func loadDir(name string, dirName *string) {
	dir := os.Getenv(name)
	if dir == "" {
		log.Fatal(fmt.Sprintf("%s not set", name))
	} else {
		dirName = &dir
	}
}

func loadMimeTypes() {
	var cfg mimetypeImport
	mimes := make(map[string]bool)
	filename := os.Getenv("QS_MIMETYPES")
	if filename == "" {
		log.Fatal("QS_MIMETYPES unset")
	}
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&cfg); err != nil {
		fmt.Println("Failed to decode MIMETypes file.")
		log.Fatal(err)
	}
	for _, item := range cfg.Allowed {
		mimes[item] = true
	}
	MimeTypes = mimes
}