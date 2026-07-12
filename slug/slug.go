// Package slug biedt een functie om vrije tekst om te zetten naar een URL-vriendelijke slug.
package slug

import (
	"strings"
	"unicode"
)

// Slugify zet s om naar een slug: kleine letters en cijfers blijven, hoofdletters
// worden kleine letters, elke aaneengesloten reeks overige tekens wordt precies
// één koppelteken, en het resultaat heeft geen koppeltekens aan begin of eind.
func Slugify(s string) string {
	var b strings.Builder
	prevDash := false
	for _, r := range s {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(unicode.ToLower(r))
			prevDash = false
		default:
			if !prevDash && b.Len() > 0 {
				b.WriteByte('-')
				prevDash = true
			}
		}
	}
	return strings.TrimSuffix(b.String(), "-")
}
