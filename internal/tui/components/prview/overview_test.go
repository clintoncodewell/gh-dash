package prview

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/stretchr/testify/require"

	"github.com/dlvhdr/gh-dash/v4/internal/data"
)

func TestOverviewRendersChangesLast(t *testing.T) {
	prData := &data.PullRequestData{State: "OPEN"}
	m := newTestModelWithWidth(t, prData, nil, nil, 80)
	m.pr.Data.Enriched.Body = "Summary"

	view := ansi.Strip(m.viewOverviewTab())
	checksIndex := strings.Index(view, "Checks")
	changesIndex := strings.Index(view, "Changes")

	require.NotEqual(t, -1, checksIndex)
	require.NotEqual(t, -1, changesIndex)
	require.Less(t, checksIndex, changesIndex)
}
