package main

import (
	"asciiart"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Art struct {
	Output string
}

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
	inputString := r.FormValue("input")
	banner := r.FormValue("banner")
	output := Art{Output: asciiart.AsciiArt(inputString, banner)}
	t, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, output)
}
