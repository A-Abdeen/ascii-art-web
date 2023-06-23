package main

import (
	"asciiart"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	/*
		TODO Tutorial 1
		http.HandleFunc("/", SimpleServer)
		http.ListenAndServe(":8080", nil)
	*/

	// TODO Tutorial 2
	fileServer := http.FileServer(http.Dir("./static")) // tell server to look into the directory / folder

	http.Handle("/", fileServer)            // ("/" is called root route) this handle tells server to send the route request into the filesServer above
	http.HandleFunc("/form", formHandler)   // Handle a route by calling a function
	http.HandleFunc("/hello", helloHandler) // Handle route by calling a function

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	// Check if input is correct
	var rawInput string
	var outputFile string
	isPrint := true
	file := "standard"
	outputErr := "\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n\n"

	if len(os.Args) == 2 { // [STRING]
		rawInput = os.Args[1]
	} else if len(os.Args) == 3 { // [STRING] [BANNER]
		rawInput = os.Args[1]
		file = os.Args[2]
	} else if len(os.Args) == 4 { // [OPTION] [STRING] [BANNER]
		if os.Args[1][:9] == "--output=" {
			outputFile = os.Args[1][9:]
			isPrint = false
		} else {
			fmt.Print(outputErr)
		}
		rawInput = os.Args[2]
		file = os.Args[3]
	} else { // ERROR (FOR OUTPUT USE)
		fmt.Print(outputErr)
		return
	}
	// Banner Loader
	switch {
	case file == "standard":
		file = "standard.txt"
	case file == "shadow":
		file = "shadow.txt"
	case file == "thinkertoy":
		file = "thinkertoy.txt"
	default:
		fmt.Println("\nAvailable banner formats are: standard, shadow or thinkertoy.")
		return
	}
	sourceFile, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Main function: Splitting (split string based on newline position)
	// ∟--> Sub function: Formatting (change input to allow use of newline & qoutation marks)
	splitInput := asciiart.LineSplitter(rawInput, asciiart.InputFormatter)

	// Main function: Printing (printing the row of characters within input string)
	// ∟--> Sub function: Parsing (parsing the data of the 8 rows to print sequentially)
	asciiart.RowPrinter(splitInput, sourceFile, asciiart.RowParser, outputFile, isPrint)

}

/*
TODO tutorial 1
func SimpleServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}
*/

// TODO tutorial 2

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/hello" {
	// 	http.Error(w, "404 not found", http.StatusNotFound)
	// 	return
	// }
	if r.Method != "GET" { // We don't want the user to post anything
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

}
