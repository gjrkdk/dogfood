package wrap

import "testing"

// TestWrap is de specificatie van Wrap(s, width): woorden (gescheiden door
// willekeurige witruimte, inclusief newlines) worden greedy over regels van
// maximaal width tekens verdeeld, gescheiden door één spatie; een woord
// langer dan width wordt hard afgebroken op width; regels worden met "\n"
// verbonden; width <= 0 geeft "".
func TestWrap(t *testing.T) {
	cases := []struct {
		name  string
		in    string
		width int
		want  string
	}{
		{"past op één regel", "hello", 10, "hello"},
		{"greedy vullen", "aa bb cc dd", 5, "aa bb\ncc dd"},
		{"exacte pasvorm", "abcde", 5, "abcde"},
		{"woord exact breedte plus volgend", "abcde fg", 5, "abcde\nfg"},
		{"te lang woord wordt gebroken", "abcdefghij", 4, "abcd\nefgh\nij"},
		{"rest van gebroken woord telt mee", "abcdefg hi", 4, "abcd\nefg\nhi"},
		{"meerdere spaties collapsen", "a   b", 10, "a b"},
		{"newlines zijn ook witruimte", "a\nb c", 10, "a b c"},
		{"randen getrimd", "  hallo  ", 10, "hallo"},
		{"lege invoer", "", 5, ""},
		{"breedte nul", "hallo", 0, ""},
		{"breedte één met lang woord", "abc", 1, "a\nb\nc"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Wrap(c.in, c.width); got != c.want {
				t.Errorf("Wrap(%q, %d) = %q, want %q", c.in, c.width, got, c.want)
			}
		})
	}
}
