package main

import (
	"asciiart"
	"fmt"
	"log"
	"net/http"
)

// type ArtLoader struct {
// 	Active bool
// 	Art    string
// }

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/asciiart", ArtHandler)
	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func ArtHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	InputString := r.FormValue("input")
	banner := r.FormValue("banner")
	// if len(InputString) > 30 {
	// 	fmt.Fprintf(w, " HTTP Error 400 - Bad Request", http.StatusBadRequest)
	// 	return
	// }
	InputString = asciiart.AsciiArt(InputString, banner)
	// p := ArtLoader{Active: true, Art: InputString}
	// t, _ := template.ParseFiles("asciiart.html")
	// t.Execute(w, p)
	fmt.Fprintf(w, "%s", InputString)

}
