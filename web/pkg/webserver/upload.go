package webserver

import (
	"QuickShare/pkg/config"
	"QuickShare/pkg/files"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path"
)

func HandleUpload(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fmt.Fprintf(w, "Invalid method %s", req.Method)
		return
	}
	err := FileUpload(req)
	if err != nil {
		fmt.Fprintf(w, "500 Error!")
		fmt.Printf("Error: %e\n", err)
		return
	}
	fmt.Fprintf(w, "Upload successful")
	log.Println("File uploaded")
}

func FileUpload(req *http.Request) error {
	req.ParseMultipartForm(32 << 20)
	file, header, err := req.FormFile("file")
	defer file.Close()
	if err != nil {
		fmt.Println(fmt.Sprintf("Err %e", err))
		return err
	}
	buff := make([]byte, 512) // Used to get MIME type
	if _, err := file.Read(buff); err != nil {
		fmt.Println(fmt.Sprintf("Err %e", err))
		return err
	}
	contentType := http.DetectContentType(buff)
	if _, err := file.Seek(0, 0); err != nil {
		return errors.New("data seek error")
	}
	// Stick content type into a map at program start
	if config.Config.MimeTypes[contentType] {
		filename, err := files.GenerateFilename(config.Config.UploadDir, path.Ext(header.Filename))
		if err != nil {
			return err
		}
		return files.WriteFile(file, filename)
	} else {
		return errors.New(fmt.Sprintf("invalid mimetype %s", contentType))
	}
}
