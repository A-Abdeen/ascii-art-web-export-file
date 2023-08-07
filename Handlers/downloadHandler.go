package asciiart

import (
	"asciiart"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	// if r.Method != "POST" {
	// 	ErrorHandler(w, r, http.StatusMethodNotAllowed)
	// 	return
	// } else if r.URL.Path != "/download" {
	// 	ErrorHandler(w, r, http.StatusNotFound)
	// 	return
	// } else {
	inputString := r.FormValue("download")
	asciiart.FileWriter(inputString)
	file, err := os.Open("Ascii-Art.txt")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	fileHeader := make([]byte, 10000)
	file.Read(fileHeader)
	fileContentType := http.DetectContentType(fileHeader)
	// ServeContent uses the name for mime detection
	fmt.Println("1" + fileInfo.Name())
	// tell the browser the returned content should be downloaded
	w.Header().Set("Content-Disposition", "attachment, filename="+fileInfo.Name())
	w.Header().Set("Content-Length", strconv.FormatInt((fileInfo.Size()), 10))
	w.Header().Set("Content-Type", fileContentType)

	// http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
	file.Seek(0, 0)
	io.Copy(w, file)
	fmt.Println("2" + fileInfo.Name())

	fmt.Println("INPUT STRING\n" + inputString)
	defer os.Remove(file.Name()) // clean up
	ArtHandler(w, r)
}

/*
   // func DownloadHandler(w http.ResponseWriter, r *http.Request) {
   // 	r.Body.Read()
   // }

   	func TestHandler() {
   		url := "https://asciiart.com/file.txt"
   		err := Download("local.txt", url)
   		if err != nil {
   			fmt.Printf("An error occured: %v", err)
   			os.Exit(1)
   		}
   	}

   	func Download(filepath, url string) error {
   		resp, err := http.Get(url)
   		if err != nil {
   			return fmt.Errorf("error while making GET request: %w", err)
   		}
   		defer resp.Body.Close()

   		out, err := os.Create(filepath)
   		if err != nil {
   			return fmt.Errorf("error while creating file: %w", err)
   		}
   		defer out.Close()

   		_, err = io.Copy(out, resp.Body)
   		if err != nil {
   			return fmt.Errorf("error while copying data: %w", err)
   		}

   		return nil
*/
