package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

var Config config

type config struct {
	UploadDir string
	ShareDir  string
	MimeTypes map[string]bool
}

type mimetypeImport struct {
	Allowed []string `json:"allowedMimeTypes"`
}

func init() {
	shouldLoadEnvFile()
	Config.UploadDir = getDirOrFail("QS_UPLOAD_DIR")
	Config.ShareDir = getDirOrFail("QS_SHARE_DIR")
	Config.MimeTypes = loadMimeTypes()
}

func shouldLoadEnvFile() {
	shouldLoad := false
	for _, item := range []string{"QS_UPLOAD_DIR", "QS_SHARE_DIR", "QS_MIME_TYPES", "QS_PORT", "QS_HOST"} {
		if os.Getenv(item) == "" {
			log.Printf("Warning %s not set!\n", item)
			shouldLoad = true
		}
	}
	if shouldLoad {
		LoadEnvFile()
	}
}

func LoadEnvFile() {
	log.Println("Attempting to load .env file")
	cwd, err := os.Getwd()
	if err != nil {
		log.Printf("Error reading working dir: %e", err)
		return
	}
	env := path.Join(cwd, ".env")
	file, err := os.Open(env)
	defer file.Close()
	if err != nil {
		log.Printf("Error opening env file %s: %e", env, err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.SplitN(line, "=", 2)
		if string(line[0]) != "#" {
			os.Setenv(split[0], split[1])
		}
	}
}

func getDirOrFail(name string) string {
	dir := os.Getenv(name)
	if dir == "" {
		log.Fatal(fmt.Sprintf("%s not set", name))
	}
	return dir
}

func loadMimeTypes() map[string]bool {
	var cfg mimetypeImport
	mimes := make(map[string]bool)
	filename := os.Getenv("QS_MIME_TYPES")
	if filename == "" {
		log.Fatal("Fatal: QS_MIME_TYPES unset")
	}
	file, err := os.Open(filename)
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
