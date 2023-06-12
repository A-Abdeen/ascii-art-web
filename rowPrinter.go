package asciiart

import "fmt"

func RowPrinter(splitInput []string, sourceFile []byte, f func([]byte, int) []byte) {
	fullRowData := ""
	for _, singleLine := range splitInput { // to print one line at a time
		if singleLine != "" {
			for i := 2; i < 10; i++ {   // to print 8 lines of each character
				for _, charRune := range singleLine { // to combine the prespective line of each char to the next
					rowLocation := (((int(charRune) - 32) * 9) + i)
					charRowData := f(sourceFile, rowLocation)
					fullRowData = fullRowData + string(charRowData)
				}
				fmt.Println(fullRowData)
				fullRowData = ""
			}
		} else {
			fmt.Print("\n")
		}
	}
}
