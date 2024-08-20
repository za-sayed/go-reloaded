package functions

import (
	"regexp"
	"sort"
)

func ApplyTransformations(text string) string {
	patterns := []struct {
		pattern       string
		transformFunc func(string) string
	}{
		{`(?i)(\-+)*(\b[0-9A-Fa-f]*)*\s*\({1,}\s*hex\s*\)([0-9a-z]*)`, ReplaceHex},
		{`(?i)(\-+)*(\b[01]*)*\s*\({1,}\s*bin\s*\){1,}([0-9a-z]*)`, ReplaceBin},
		{`(?i)(.*?)\s*(\({1,}\s*cap\s*((,)\s*([-+]?\d*)\s*)?\){1,})([0-9a-z]*)`, ReplaceCap},
		{`(?i)(.*?)\s*(\({1,}\s*up\s*((,)\s*([-+]?\d*)\s*)?\){1,})([0-9a-z]*)`, ReplaceUp},
		{`(?i)(.*?)\s*(\({1,}\s*low\s*((,)\s*([-+]?\d*)\s*)?\){1,})([0-9a-z]*)`, ReplaceLow},
	}

	var matches []struct {
		start         int
		end           int
		pattern       string
		transformFunc func(string) string
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern.pattern)
		for _, match := range re.FindAllStringIndex(text, -1) {
			start, end := match[0], match[1]
			matches = append(matches, struct {
				start         int
				end           int
				pattern       string
				transformFunc func(string) string
			}{
				start:         start,
				end:           end,
				pattern:       pattern.pattern,
				transformFunc: pattern.transformFunc,
			})
		}
	}
	sort.SliceStable(matches, func(i, j int) bool {
		if matches[i].start < matches[j].start {
			return true
		} else if matches[i].start == matches[j].start {
			return matches[i].end < matches[j].end
		} else {
			return false
		}
	})
	

	count := 0
	for count < len(matches) {
		matchedText := text[:matches[count].end]
		transformedText := matches[count].transformFunc(matchedText)
		text = transformedText + text[matches[count].end:]
		matches = nil
		for _, pattern := range patterns {
			re := regexp.MustCompile(pattern.pattern)
			for _, match := range re.FindAllStringIndex(text, -1) {
				start, end := match[0], match[1]
				matches = append(matches, struct {
					start         int
					end           int
					pattern       string
					transformFunc func(string) string
				}{
					start:         start,
					end:           end,
					pattern:       pattern.pattern,
					transformFunc: pattern.transformFunc,
				})
			}
		}
		sort.SliceStable(matches, func(i, j int) bool {
			if matches[i].start < matches[j].start {
				return true
			} else if matches[i].start == matches[j].start {
				return matches[i].end < matches[j].end
			} else {
				return false
			}
		})
	}
	return text
}
