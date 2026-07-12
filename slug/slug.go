// Package slug biedt het omzetten van tekst naar URL-vriendelijke slugs.
package slug

import "strings"

// Slugify zet s om naar een slug: kleine letters en cijfers blijven, hoofdletters
// worden kleine letters, elke aaneengesloten reeks overige tekens wordt precies
// één koppelteken, en het resultaat heeft geen koppeltekens aan begin of eind.
func Slugify(s string) string {
	var b strings.Builder
	prevHyphen := true

	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z' || r >= '0' && r <= '9':
			b.WriteRune(r)
			prevHyphen = false
		case r >= 'A' && r <= 'Z':
			b.WriteRune(r - 'A' + 'a')
			prevHyphen = false
		default:
			if !prevHyphen {
				b.WriteByte('-')
				prevHyphen = true
			}
		}
	}

	return strings.TrimSuffix(b.String(), "-")
}
