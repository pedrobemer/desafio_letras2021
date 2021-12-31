package comparator

import "unicode"

// Compare two strings to see if they are the same using insentive mode
func Insensitive(firstWord, secondWord string) bool {

	// If two strings have a different length then they are not the same
	if len(firstWord) != len(secondWord) {
		return false
	}

	for i, letter := range firstWord {
		// if the characters already match then we don't need to
		// alter their case. We can continue to the next rune
		if firstWord[i] == secondWord[i] {
			continue
		}
		if unicode.ToLower(letter) != unicode.ToLower(rune(secondWord[i])) {
			// the lowercase characters do not match so these
			// are considered a mismatch, break and return false
			return false
		}
	}
	// The string length has been traversed without a mismatch
	// therefore the two match
	return true
}
