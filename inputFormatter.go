package asciiart

import "strings"

func InputFormatter(rawInput string) string {
	formattedText := strings.ReplaceAll(rawInput, "\\n", "\n")
	formattedText = strings.ReplaceAll(formattedText, "\"", string('"'))
	formattedText = strings.ReplaceAll(formattedText, "\\'", "'")
	formattedText = strings.ReplaceAll(formattedText, "\\!", string('!')) // at end of line it will not work without the backs
	return formattedText
}
