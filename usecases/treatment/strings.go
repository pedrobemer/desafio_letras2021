package treatment

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Responsible for removing all accents of a given text and to lower case all
// caracterers from the same text.
func TreatString(text string) (string, error) {

	// Convert caracters from the text to lower case
	textLowerCased := strings.ToLower(text)

	// Remove accents from the treated text
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	treatedString, _, err := transform.String(t, textLowerCased)
	if err != nil {
		return "", err
	}

	return treatedString, nil
}
