package functions

import (
	"fmt"
	"regexp"
	"strings"
)

func FormatPunctuation(text string) string {
	punctPattern := `(?i)(\w*)\s*([.,!?;:]+(?:\s+[.,!?;:]+)*)\s*(\w*)`
	re := regexp.MustCompile(punctPattern)

	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) < 2 {
			return match
		}
		wordBefore := strings.TrimSpace(parts[1])
		wordAfter := strings.TrimSpace(parts[3])
		punct := parts[2]

		punct = strings.ReplaceAll(punct, " ", "")

		if wordAfter != "" {
			return fmt.Sprintf("%s%s %s", wordBefore, punct, wordAfter)
		} else {
			return fmt.Sprintf("%s%s", wordBefore, punct)
		}
	})
}
