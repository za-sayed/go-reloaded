package functions

import (
	"fmt"
	"regexp"
	"strings"
)

func FormatQuotation(text string) string {
	quotePattern := `(?i)(\s*)'((?:\\.|'(?:\w)|[^'])*)'`
	re := regexp.MustCompile(quotePattern)

	text = re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) < 2 {
			return match
		}
		quotedText := strings.TrimSpace(parts[2])
		quotedWords := strings.Fields(quotedText)
		index := strings.Index(text, match)
		if index == 0 {
			return fmt.Sprintf("'%s'", strings.Join(quotedWords, " "))
		} else {
			return fmt.Sprintf(" '%s'", strings.Join(quotedWords, " "))
		}

	})
	return text
}
