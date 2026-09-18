package tui

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/dlvhdr/gh-dash/v4/internal/data"
)

type issueCreateTestRow struct {
	repo string
}

func (r issueCreateTestRow) GetRepoNameWithOwner() string { return r.repo }
func (r issueCreateTestRow) GetTitle() string             { return "" }
func (r issueCreateTestRow) GetNumber() int               { return 0 }
func (r issueCreateTestRow) GetUrl() string               { return "" }
func (r issueCreateTestRow) GetUpdatedAt() time.Time      { return time.Time{} }

func TestResolveIssueCreateRepo(t *testing.T) {
	tests := []struct {
		name         string
		row          data.RowData
		filters      string
		fallbackRepo string
		want         string
		wantErr      string
	}{
		{
			name:         "selected issue takes priority",
			row:          issueCreateTestRow{repo: "selected/repo"},
			filters:      "repo:filtered/repo is:open",
			fallbackRepo: "fallback/repo",
			want:         "selected/repo",
		},
		{
			name:         "single section repository",
			filters:      "is:open repo:filtered/repo label:bug",
			fallbackRepo: "fallback/repo",
			want:         "filtered/repo",
		},
		{
			name:         "quoted section repository",
			filters:      `is:open repo:"filtered/repo"`,
			fallbackRepo: "fallback/repo",
			want:         "filtered/repo",
		},
		{
			name:         "repository filters are deduplicated case insensitively",
			filters:      "repo:Owner/Repo repo:owner/repo",
			fallbackRepo: "fallback/repo",
			want:         "Owner/Repo",
		},
		{
			name:         "typed nil row falls through to section repository",
			row:          (*data.IssueData)(nil),
			filters:      "repo:filtered/repo",
			fallbackRepo: "fallback/repo",
			want:         "filtered/repo",
		},
		{
			name:         "launch repository fallback",
			fallbackRepo: "fallback/repo",
			want:         "fallback/repo",
		},
		{
			name:    "multiple section repositories are ambiguous",
			filters: "repo:first/repo repo:second/repo",
			wantErr: "multiple repositories",
		},
		{
			name:    "missing repository gives instructions",
			wantErr: "add a repo:owner/name filter",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resolveIssueCreateRepo(tt.row, tt.filters, tt.fallbackRepo)
			if tt.wantErr != "" {
				require.ErrorContains(t, err, tt.wantErr)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestNewIssueCommand(t *testing.T) {
	cmd := newIssueCommand("clintoncodewell/advisewell")

	require.Equal(t, []string{
		"gh", "issue", "create", "--repo", "clintoncodewell/advisewell",
	}, cmd.Args)
}

type issueCreateTestExitError int

func (e issueCreateTestExitError) Error() string { return "exit" }
func (e issueCreateTestExitError) ExitCode() int { return int(e) }

func TestIssueCreateWasCancelled(t *testing.T) {
	require.True(t, issueCreateWasCancelled(issueCreateTestExitError(2)))
	require.True(t, issueCreateWasCancelled(
		fmt.Errorf("wrapped: %w", issueCreateTestExitError(2)),
	))
	require.False(t, issueCreateWasCancelled(issueCreateTestExitError(1)))
	require.False(t, issueCreateWasCancelled(errors.New("plain error")))
}
