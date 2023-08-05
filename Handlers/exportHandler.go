package asciiart

import ("os"
"net/http"
"fmt"
"html/template"
"log"
)

func ExportHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("1")
	// if err := r.ParseForm(); err != nil {
	// 	fmt.Fprintf(w, "ParseForm() err: %v", err)
	// 	return
	// }
	// inputString := r.FormValue("input")
	// fmt.Println(inputString)
	// banner := r.FormValue("banner")
	// fmt.Println(banner)
	files := []string{
		"templates/base.html",
		"templates/form.tmpl",
		"templates/output.tmpl",
		"templates/error.tmpl",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal(err)
	}
	
	finaloutput := r.FormValue("download")
	fmt.Println(finaloutput)
	os.WriteFile("outputfile.txt", []byte(finaloutput), 0644)
	t.ExecuteTemplate(w, "base.html", finaloutput)
}