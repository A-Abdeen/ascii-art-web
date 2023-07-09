package asciiart

import (
	"os"
)

func AsciiArt(InputString string, banner string) string {
	switch {
	case banner == "standard":
		banner = "standard.txt"
	case banner == "shadow":
		banner = "shadow.txt"
	case banner == "thinkertoy":
		banner = "thinkertoy.txt"
	default:
		return "\nAvailable banner formats are: standard, shadow or thinkertoy."
	}
	sourceFile, err := os.ReadFile(banner)
	if err != nil {
		return "err: " + err.Error()
	}
	// Main function: Splitting (split string based on newline position)
	// ∟--> Sub function: Formatting (change input to allow use of newline & qoutation marks)
	splitInput := LineSplitter(InputString, InputFormatter)

	// Main function: Printing (printing the row of characters within input string)
	// ∟--> Sub function: Parsing (parsing the data of the 8 rows to print sequentially)
	return RowPrinter(splitInput, sourceFile, RowParser)

}
