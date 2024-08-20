package functions

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReplaceBin(text string) string {
	binPattern := `(?i)(\-+)*(\b[01]*)*\s*\({1,}\s*bin\s*\){1,}([0-9a-z]*)`
	re := regexp.MustCompile(binPattern)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if parts[3] != "" {
			startIdx := strings.Index(text, match)
			endIdx := startIdx + len(match) - len(parts[3])
			invalidPartEnd := startIdx + len(match)
			invalidPart := text[startIdx:invalidPartEnd]
			fmt.Printf("Error: Invalid format detected \"%s\" remove \"%s\" to make it valid\n", invalidPart, text[endIdx:invalidPartEnd])
			os.Exit(1)
		}
		if parts[2] == "" || parts[1] != "" {
			fmt.Println("Error: Missing binary number before (bin).")
			os.Exit(1)
		}

		if len(parts) < 2 {
			return match
		}
		var binNumber string
		if parts[2] != "" {
			binNumber = strings.TrimSpace(parts[2])
		} else {
			return match
		}
		decimalValue, err := strconv.ParseInt(binNumber, 2, 64)
		if err != nil {
			fmt.Println("Error: Overflow")
			os.Exit(1)
		}
		return fmt.Sprintf("%d", decimalValue)
	})
}
