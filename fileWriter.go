package asciiart

import (
	"fmt"
	"log"
	"os"
)

func FileWriter(art string) {
	// file, err := os.CreateTemp("/Users/abdyn/Codes/ascii-art-web-exportfile/ascii-art-web-exportfile/webprogram", "AsciiArt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // defer os.Remove(file.Name()) // clean up

	// if _, err := file.Write([]byte(art)); err != nil {
	// 	log.Fatal(err)
	// }
	// if err := file.Close(); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(file)
	// return file

	outputFile, err := os.OpenFile("Ascii-Art.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer outputFile.Close()
	// Clear output file before next run
	if err := os.Truncate("Ascii-Art.txt", 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}
	if _, err := outputFile.WriteString(art); err != nil {
		log.Println(err)
	} else {
		fmt.Println("------> Done. Go to output file <------")
	}
}
