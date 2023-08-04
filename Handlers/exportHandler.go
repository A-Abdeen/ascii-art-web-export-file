package asciiart

import ("os"
"net/http"
"fmt"
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
	finaloutput := r.FormValue("download")
	fmt.Println(finaloutput)
	os.WriteFile("outputfile", []byte(finaloutput), 0644)
}