package functions

import (
	"regexp"
)

func ConvertAToAn(text string) string {
	aPattern := `\b([aA])\s+([aeiouhAEIOUH]\w+)`
	re := regexp.MustCompile(aPattern)

	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) < 3 {
			return match
		}
		if parts[1] == "a" {
			return "an " + parts[2]
		} else if parts[1] == "A" {
			return "An " + parts[2]
		} else {
			return "" 
		}
	})
}
