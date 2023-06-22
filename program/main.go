package main

import (
	"asciiart"
	"fmt"
	"os"
)

func main() {
	/*
		TODO Tutorial 1
		http.HandleFunc("/", SimpleServer)
		http.ListenAndServe(":8080", nil)
	*/
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
