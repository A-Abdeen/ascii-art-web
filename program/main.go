package main

import (
	"asciiart"
	"fmt"
	"os"
)

func main() {
	// Check if input is correct
	if len(os.Args) > 3 {
		fmt.Println("Too many arguments")
		return
	} 
	if len(os.Args) < 2 {
		fmt.Println("Input string missing")
		return
	}
	argNum := 2
	file := os.Args[1] // Read char file & string argument
	if len(os.Args) == 2 {
		file = "standard.txt"
		argNum--
	}
	sourceFile, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	rawInput := os.Args[argNum]

	// Main function: Splitting (split string based on newline position)
	// ∟--> Sub function: Formatting (change input to allow use of newline & qoutation marks)
	splitInput := asciiart.LineSplitter(rawInput, asciiart.InputFormatter)

	// Main function: Printing (printing the row of characters within input string)
	// ∟--> Sub function: Parsing (parsing the data of the 8 rows to print sequentially)
	asciiart.RowPrinter(splitInput, sourceFile, asciiart.RowParser)

}
