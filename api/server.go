package api

import (
	"log"
	"net/http"
)

func Server() {
	// Following tutorial

	// h1 := func(w http.ResponseWriter, _ *http.Request) {
	// 	io.WriteString(w, "Hello from a HandleFunc #1!\n")
	// }
	// h2 := func(w http.ResponseWriter, _ *http.Request) {
	// 	io.WriteString(w, "Hello from a HandleFunc #2!\n")
	// }

	// http.HandleFunc("/", h1)
	// http.HandleFunc("/ascii-art", h2)

	http.Handle("/", http.FileServer(http.Dir("/Users/abdyn/Codes/ascii-art-web/ascii-art-web/api/ascii-website")))
	http.Handle("/art.html", http.FileServer(http.Dir("/Users/abdyn/Codes/ascii-art-web/ascii-art-web/api/ascii-website")))
	// some := func(w http.ResponseWriter, _ *http.Request) {
	// 	io.WriteString(w, "/Users/abdyn/Codes/ascii-art-web/ascii-art-web/api/ascii-art/index.html")
	// }
	// http.HandleFunc("/ascii-art", some)

	// http.Handle("/ascii-art/", http.StripPrefix("/ascii-art/", http.FileServer(http.Dir("./api"))))

	// TODO Add page for results
	// http.Handle("/ascii-art", http.FileServer(http.Dir("/Users/abdyn/Codes/ascii-art-web/ascii-art-web/api/ascii-website/text")))
	// x := http.FileServer(http.Dir("/Users/abdyn/Codes/ascii-art-web/ascii-art-web/program/shadow.txt"))

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // handler function
	// 	// setting status code or content type
	// 	w.Write([]byte("HELLO")) // should use this to load the HTML file
	// }) // homepage, handle requests
	log.Fatal(http.ListenAndServe(":8080", nil)) // listen for the http connection

}

/*
GET request to load data
Server respondes with a json file with requested data
POST mode or modify data

CLIENT Request[GET] HTML home page
SERVER responds with loading home page
CLIENT sends [POST] string and banner options
SERVER responds with loading /ascii-art page with correct output








*/
