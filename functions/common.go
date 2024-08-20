package functions

import "strconv"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetNumber(parts []string) int {
	if len(parts) > 5 && parts[5] != "" {
		number, err := strconv.Atoi(parts[5])
		if err == nil {
			return number
		}
	}
	return 1
}
