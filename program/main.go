package main

import (
	"asciiart"
	"fmt"
	"log"
	"net/http"
	"os"
)

func artHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful \n")
	input := r.FormValue("input")
	x := AsciiArt(input)
	fmt.Fprintf(w, "Input = %s\n", input)
	fmt.Fprintf(w, "%s", x)
}

func main() {

	fileServer := http.FileServer(http.Dir("./static")) // tell server to look into the directory / folder
	// ("/" is called root route)
	http.Handle("/", fileServer)              // Handle tells server to send the route request into the filesServer above
	http.HandleFunc("/ascii-art", artHandler) // Handle a route by calling a function
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func AsciiArt(rawInput string) string {
	// Check if input is correct
	// var rawInput string
	var outputFile string
	isPrint := true
	file := "standard"
	// outputErr := "\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n\n"

	// if len(os.Args) == 2 { // [STRING]
	// 	rawInput = os.Args[1]
	// } else if len(os.Args) == 3 { // [STRING] [BANNER]
	// 	rawInput = os.Args[1]
	// 	file = os.Args[2]
	// } else if len(os.Args) == 4 { // [OPTION] [STRING] [BANNER]
	// 	if os.Args[1][:9] == "--output=" {
	// 		outputFile = os.Args[1][9:]
	// 		isPrint = false
	// 	} else {
	// 		fmt.Print(outputErr)
	// 	}
	// 	rawInput = os.Args[2]
	// 	file = os.Args[3]
	// } else { // ERROR (FOR OUTPUT USE)
	// 	return outputErr
	// 	// return
	// }
	// Banner Loader
	switch {
	case file == "standard":
		file = "standard.txt"
	case file == "shadow":
		file = "shadow.txt"
	case file == "thinkertoy":
		file = "thinkertoy.txt"
	default:
		return "\nAvailable banner formats are: standard, shadow or thinkertoy."
	}
	sourceFile, err := os.ReadFile(file)
	if err != nil {
		return err.Error()
	}

	// Main function: Splitting (split string based on newline position)
	// ∟--> Sub function: Formatting (change input to allow use of newline & qoutation marks)
	splitInput := asciiart.LineSplitter(rawInput, asciiart.InputFormatter)

	// Main function: Printing (printing the row of characters within input string)
	// ∟--> Sub function: Parsing (parsing the data of the 8 rows to print sequentially)

	return asciiart.RowPrinter(splitInput, sourceFile, asciiart.RowParser, outputFile, isPrint)

}
