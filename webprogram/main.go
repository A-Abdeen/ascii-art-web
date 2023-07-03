package main

import (
	"asciiart"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/asciiart", ArtHandler)
	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

type Testing struct {
	Title string
	News  string
}

func ArtHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	InputString := r.FormValue("input")
	banner := r.FormValue("banner")
	InputString = asciiart.AsciiArt(InputString, banner)
	p := Testing{Title: InputString, News: "testing"}
	// fmt.Println(p)
	// testing := "test"
	// fmt.Fprintf(w, "%s", InputString)
	// fmt.Println("")
	t, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(t.Execute(w, p))
}
