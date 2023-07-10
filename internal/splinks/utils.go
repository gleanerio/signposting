package splinks

import "strings"

func isAbsoluteURL(url string) bool {
	// Check if the URL starts with "http://" or "https://".
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return true
	}

	// Check if the URL starts with "//".
	if strings.HasPrefix(url, "//") {
		return true
	}

	return false
}

func isRelativeURL(url string) bool {
	// Check if the URL starts with a forward slash.
	if strings.HasPrefix(url, "/") {
		return true
	}

	return false
}

// StringInSlice checks if a string exists in a slice of strings.
func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		//		if v == str {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}
