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

// MustMake is Slugify, maar raakt in paniek als het resultaat leeg is. Een
// lege slug duidt op een programmeerfout bij een niet-lege, niet-triviale
// invoer (bijvoorbeeld een vaste tekst in de code), dus gebruik MustMake
// alleen wanneer een lege slug nooit legitiem kan zijn.
func MustMake(s string) string {
	slug := Slugify(s)
	if slug == "" {
		panic("slug: MustMake(" + s + ") geeft een lege slug")
	}

	return slug
}

// SlugifyN is Slugify, maar het resultaat wordt afgekapt op maximaal maxLen
// tekens. Er wordt niet midden in een woord gekapt: als de limiet in een
// woord valt, wordt teruggekapt tot het laatste koppelteken ervoor. Het
// resultaat eindigt nooit op een koppelteken. maxLen <= 0 geeft "".
func SlugifyN(s string, maxLen int) string {
	return Truncate(Slugify(s), maxLen)
}

// Truncate kapt een slug s af op maximaal max tekens zonder een woord
// doormidden te snijden: als de limiet in een woord valt, wordt teruggekapt
// tot het laatste koppelteken ervoor, en een eventueel achterblijvend
// koppelteken aan het eind wordt verwijderd. Is s korter dan of gelijk aan
// max, dan wordt s ongewijzigd teruggegeven. max <= 0 of een lege s geeft "".
func Truncate(s string, max int) string {
	if max <= 0 || s == "" {
		return ""
	}
	if len(s) <= max {
		return s
	}

	cut := s[:max]
	if s[max] != '-' {
		if i := strings.LastIndexByte(cut, '-'); i >= 0 {
			cut = cut[:i]
		}
	}

	return strings.TrimSuffix(cut, "-")
}

// IsSlug controleert of s een geldige slug is: alleen kleine letters, cijfers
// en enkele koppeltekens, niet beginnend of eindigend op een koppelteken, en
// zonder dubbele koppeltekens. Een lege string is geen geldige slug.
func IsSlug(s string) bool {
	if s == "" || s[0] == '-' || s[len(s)-1] == '-' {
		return false
	}

	prevHyphen := false
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z' || r >= '0' && r <= '9':
			prevHyphen = false
		case r == '-':
			if prevHyphen {
				return false
			}
			prevHyphen = true
		default:
			return false
		}
	}

	return true
}
