// Package wrap biedt tekstopmaak door woorden greedy over regels te verdelen.
package wrap

import "strings"

// Wrap verdeelt de woorden in s greedy over regels van maximaal width tekens,
// gescheiden door "\n". Woorden langer dan width worden hard afgebroken.
func Wrap(s string, width int) string {
	if width <= 0 {
		return ""
	}

	fields := strings.Fields(s)
	if len(fields) == 0 {
		return ""
	}

	var lines []string
	var cur strings.Builder
	curLen := 0

	flush := func() {
		if curLen > 0 {
			lines = append(lines, cur.String())
			cur.Reset()
			curLen = 0
		}
	}

	for _, w := range fields {
		for len(w) > 0 {
			if curLen == 0 {
				if len(w) <= width {
					cur.WriteString(w)
					curLen = len(w)
					w = ""
				} else {
					cur.WriteString(w[:width])
					w = w[width:]
					curLen = width
					flush()
				}
			} else if avail := width - curLen - 1; avail >= len(w) {
				cur.WriteByte(' ')
				cur.WriteString(w)
				curLen += 1 + len(w)
				w = ""
			} else {
				flush()
			}
		}
	}
	flush()

	return strings.Join(lines, "\n")
}
