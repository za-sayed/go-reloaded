package functions

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ReplaceCap(text string) string {
	capPattern := `(?i)(.*?)\s*(\({1,}\s*cap\s*((,)\s*([-+]?\d*)\s*)?\){1,})([0-9a-z]*)`
	re := regexp.MustCompile(capPattern)

	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if parts[6] != "" {
			startIdx := strings.Index(text, match)
			endIdx := startIdx + len(match) - len(parts[6])
			invalidPartEnd := startIdx + len(match)
			startCapIdx := strings.Index(match, parts[2])
			invalidPart := text[startCapIdx:invalidPartEnd]
			fmt.Printf("Error: Invalid format detected \"%s\" remove \"%s\" to make it valid\n", invalidPart, text[endIdx:invalidPartEnd])
			os.Exit(1)
		}
		if parts[5] == "" && parts[4] == "," {
			fmt.Println("Error: Missing number of words for (up).")
			os.Exit(1)
		}

		if len(parts) < 2 {
			return match
		}
		numWords := GetNumber(parts)
		if numWords < 1 {
			fmt.Println("Error: The number of words specified for (cap) must be at least 1.")
			os.Exit(1)
		}
		words := strings.Fields(parts[1])
		if numWords > len(words) {
			fmt.Printf("Error: The number of words specified for (cap) exceeds the available words.\n")
			os.Exit(1)
		}
		startIdx := Max(0, len(words)-numWords)
		for i := startIdx; i < len(words); i++ {
			words[i] = strings.Title(words[i])
		}
		text := strings.Join(words, " ")
		return text
	})
}
