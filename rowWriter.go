package asciiart

import (
	"fmt"
	"log"
	"os"
)

func RowWriter(fullRowData, output string) {
	outputFile, err := os.OpenFile(output,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("failed openning in method 1")
		log.Println(err)
	}
	defer outputFile.Close()
	if _, err := outputFile.WriteString(fullRowData + "\n"); err != nil {
		fmt.Println("failed writing string in method 1")
		log.Println(err)
	}
}
