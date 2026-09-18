package util

import "strings"

// IsWarpLicenseKey recognizes the license format currently used by the WARP API.
func IsWarpLicenseKey(value string) bool {
	parts := strings.Split(value, "-")
	if len(parts) != 3 {
		return false
	}
	for _, part := range parts {
		if len(part) != 8 {
			return false
		}
		for _, char := range part {
			if !isASCIIAlphaNumeric(char) {
				return false
			}
		}
	}
	return true
}

func isASCIIAlphaNumeric(char rune) bool {
	return char >= '0' && char <= '9' ||
		char >= 'a' && char <= 'z' ||
		char >= 'A' && char <= 'Z'
}
