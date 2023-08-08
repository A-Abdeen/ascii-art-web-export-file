package asciiart

import (
	"net/http"
	"fmt"
)

func ExportHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	output := r.FormValue("output")
	format := r.FormValue("format")
	// Set the appropriate content type based on the format
	var contentType string
	var fileExtension string
	if r.Method != "POST" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
		} else if r.URL.Path != "/download" {
			ErrorHandler(w, r, http.StatusNotFound)
			return
			} else {
				// // Add more cases for other formats as needed
				switch format {
		case "txt":
			contentType = "text/plain"
			fileExtension = "txt"
		case "doc":
			contentType = "application/msword"
			fileExtension = "doc"
		default:
			ErrorHandler(w, r, http.StatusBadRequest)
		}
		// Set the HTTP headers for downloading the file
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=ascii-art.%s", fileExtension))
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(output)))
		// Write the ASCII art to the response writer
		_, err := w.Write([]byte(output))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
