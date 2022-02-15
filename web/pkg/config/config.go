package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var Config config
var MimeJsonFile = "/app/allowed-mimetypes.json"

type config struct {
	UploadDir string
	ShareDir  string
	MimeTypes map[string]bool
}

type mimetypeImport struct {
	Allowed []string `json:"allowedMimeTypes"`
}

func init() {
	Config.UploadDir = "/app/upload"
	Config.ShareDir = "/app/share"
	Config.MimeTypes = loadMimeTypes()
}

func loadMimeTypes() map[string]bool {
	var cfg mimetypeImport
	mimes := make(map[string]bool)
	file, err := os.Open(MimeJsonFile)
	defer file.Close()
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to read MIMETypes config file: %e", err))
	}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&cfg); err != nil {
		log.Fatal(fmt.Sprintf("Failed to decode MIMETypes config file: %e", err))
	}
	for _, item := range cfg.Allowed {
		mimes[item] = true
	}
	return mimes
}
