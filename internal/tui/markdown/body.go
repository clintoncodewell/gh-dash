package markdown

import (
	"regexp"
	"strings"

	"github.com/charmbracelet/x/ansi"
)

var leadingSummaryHeading = regexp.MustCompile(
	`(?i)^\s{0,3}(?:#{1,6}[ \t]+)?summary[ \t]*:?[ \t]*(?:\r?\n(?:[=-]{2,}[ \t]*(?:\r?\n+|$))?|$)`,
)

var renderedBulletMarker = regexp.MustCompile(`^\s{2,}•\s`)

// NormalizeBody removes a redundant leading Summary heading. The surrounding
// view already labels this content as a summary.
func NormalizeBody(body string) string {
	body = strings.TrimSpace(body)
	body = leadingSummaryHeading.ReplaceAllString(body, "")
	return strings.TrimSpace(body)
}

// AlignWrappedBullets removes the list block's two-cell leading indent from
// bullet lines. Continuation lines retain that indent, so wrapped text begins
// directly below the first character after the bullet.
func AlignWrappedBullets(rendered string) string {
	lines := strings.Split(rendered, "\n")
	for i, line := range lines {
		plain := ansi.Strip(line)
		if renderedBulletMarker.MatchString(plain) {
			lines[i] = ansi.Cut(line, 2, ansi.StringWidth(line))
		}
	}
	return strings.Join(lines, "\n")
}
