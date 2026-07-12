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
