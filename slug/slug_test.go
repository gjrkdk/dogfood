package slug

import "testing"

// TestSlugify is de specificatie van Slugify: kleine letters en cijfers
// blijven, hoofdletters worden kleine letters, elke aaneengesloten reeks
// overige tekens wordt precies één koppelteken, en het resultaat heeft
// geen koppeltekens aan begin of eind.
func TestSlugify(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"kleine letters blijven", "hello", "hello"},
		{"hoofdletters worden klein", "Hello", "hello"},
		{"spatie wordt koppelteken", "hello world", "hello-world"},
		{"meerdere spaties collapsen", "a   b", "a-b"},
		{"leestekens collapsen mee", "Go 1.24 is uit!", "go-1-24-is-uit"},
		{"underscores zijn ook scheiding", "UPPER_case_09", "upper-case-09"},
		{"randen worden getrimd", "  spaties  rondom  ", "spaties-rondom"},
		{"lege invoer", "", ""},
		{"alleen troep", "!!!", ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Slugify(c.in); got != c.want {
				t.Errorf("Slugify(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}

// TestSlugifyN is de specificatie van SlugifyN: het resultaat van Slugify,
// afgekapt op maximaal maxLen tekens zonder midden in een woord te breken
// (er wordt teruggekapt op het laatste koppelteken vóór de limiet), zonder
// koppelteken aan het eind. maxLen <= 0 geeft "".
func TestSlugifyN(t *testing.T) {
	cases := []struct {
		name   string
		in     string
		maxLen int
		want   string
	}{
		{"past ruim binnen de limiet", "hello", 10, "hello"},
		{"past precies binnen de limiet", "hello world", 11, "hello-world"},
		{"limiet valt precies op een koppelteken", "hello world foo", 11, "hello-world"},
		{"limiet valt midden in een woord, kapt terug tot koppelteken", "hello world foo", 8, "hello"},
		{"limiet net na een koppelteken kapt terug", "hello world", 6, "hello"},
		{"geen koppelteken vóór de limiet, hakt hard af", "helloworld", 5, "hello"},
		{"maxLen nul geeft lege string", "hello world", 0, ""},
		{"maxLen negatief geeft lege string", "hello world", -5, ""},
		{"lege invoer geeft lege string", "", 5, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := SlugifyN(c.in, c.maxLen); got != c.want {
				t.Errorf("SlugifyN(%q, %d) = %q, want %q", c.in, c.maxLen, got, c.want)
			}
		})
	}
}
