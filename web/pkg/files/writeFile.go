package files

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func WriteFile(file multipart.File, filename string) error {
	w, err := os.Create(filename);
	defer w.Close()
	if err != nil {
		log.Println(fmt.Sprintf("Err %e", err))
		return err
	}
	if _, err = io.Copy(w, file); err != nil {
		log.Println(fmt.Sprintf("Err %e", err))
		return err
	}
	return nil
}
