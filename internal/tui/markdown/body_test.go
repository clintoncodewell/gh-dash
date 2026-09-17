package markdown

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/stretchr/testify/require"
)

func TestNormalizeBodyStripsOnlyLeadingSummaryHeading(t *testing.T) {
	tests := []struct {
		name string
		body string
		want string
	}{
		{name: "markdown heading", body: "## Summary\n\n- fixed it", want: "- fixed it"},
		{name: "plain heading", body: "Summary\n\nDetails", want: "Details"},
		{name: "colon", body: "Summary:\nDetails", want: "Details"},
		{name: "setext h1", body: "Summary\n=======\n\nDetails", want: "Details"},
		{name: "setext h2", body: "Summary\n---\nDetails", want: "Details"},
		{name: "not a heading", body: "Summary of the change\n\nDetails", want: "Summary of the change\n\nDetails"},
		{name: "later heading", body: "Intro\n\n## Summary\nDetails", want: "Intro\n\n## Summary\nDetails"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, NormalizeBody(tt.body))
		})
	}
}

func TestWrappedBulletUsesHangingIndent(t *testing.T) {
	markdownStyle = nil
	markdownStyleSource = ""
	renderer := GetMarkdownRenderer(24, nil)
	rendered, err := renderer.Render("- This bullet has enough words to wrap onto another line")
	require.NoError(t, err)

	plain := strings.Trim(ansi.Strip(AlignWrappedBullets(rendered)), "\r\n")
	lines := make([]string, 0)
	for line := range strings.SplitSeq(plain, "\n") {
		if strings.TrimSpace(line) != "" {
			lines = append(lines, line)
		}
	}
	require.GreaterOrEqual(t, len(lines), 2, plain)
	require.Equal(t, 0, len(lines[0])-len(strings.TrimLeft(lines[0], " ")), "%q", lines)
	require.True(t, strings.HasPrefix(lines[0], "• "), "%q", lines)
	require.Equal(t, 2, len(lines[1])-len(strings.TrimLeft(lines[1], " ")), plain)
}

func TestAlignWrappedBulletsPreservesANSIAndCodeIndentation(t *testing.T) {
	rendered := strings.Join([]string{
		"\x1b[31m  • bullet\x1b[0m",
		"    continuation",
		"\x1b[34m    • nested\x1b[0m",
		"      continuation",
		"\x1b[32m  - run: command\x1b[0m",
		"  + added diff line",
	}, "\n")

	got := AlignWrappedBullets(rendered)
	plain := ansi.Strip(got)
	require.Equal(t, strings.Join([]string{
		"• bullet",
		"    continuation",
		"  • nested",
		"      continuation",
		"  - run: command",
		"  + added diff line",
	}, "\n"), plain)
	require.Contains(t, got, "\x1b[31m")
	require.Contains(t, got, "\x1b[32m")
	require.Contains(t, got, "\x1b[34m")
}

func TestAlignWrappedBulletsPreservesFencedCode(t *testing.T) {
	markdownStyle = nil
	markdownStyleSource = ""
	renderer := GetMarkdownRenderer(40, nil)
	rendered, err := renderer.Render("```yaml\n- run: command\n1. literal code\n```")
	require.NoError(t, err)

	require.Equal(t, rendered, AlignWrappedBullets(rendered))
}
