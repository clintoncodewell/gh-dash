package footer

import (
	"testing"
	"time"

	zone "github.com/lrstanley/bubblezone/v2"
	"github.com/stretchr/testify/require"

	"github.com/dlvhdr/gh-dash/v4/internal/config"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/context"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/keys"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/theme"
)

func TestViewMarksRefreshAndViewControls(t *testing.T) {
	zone.NewGlobal()
	zone.SetEnabled(true)

	cfg, err := config.ParseConfig(config.Location{
		ConfigFlag:       "../../../config/testdata/test-config.yml",
		SkipGlobalConfig: true,
	})
	require.NoError(t, err)
	ctx := &context.ProgramContext{
		Config:      &cfg,
		ScreenWidth: 160,
		View:        config.PRsView,
	}
	ctx.Theme = theme.ParseTheme(ctx.Config)
	ctx.Styles = context.InitStyles(ctx.Theme)

	m := NewModel(ctx)
	zone.Scan(m.View())

	for _, id := range []string{
		"refresh-all",
		"help",
		"view-notifications",
		"view-prs",
		"view-issues",
	} {
		require.Eventually(t, func() bool {
			return !zone.Get(id).IsZero()
		}, 250*time.Millisecond, time.Millisecond, "expected mouse zone %q", id)
	}
}

func TestViewMarksIssueWorkflowControls(t *testing.T) {
	zone.NewGlobal()
	zone.SetEnabled(true)

	cfg, err := config.ParseConfig(config.Location{
		ConfigFlag:       "../../../config/testdata/test-config.yml",
		SkipGlobalConfig: true,
	})
	require.NoError(t, err)
	cfg.Keybindings.Issues = append(cfg.Keybindings.Issues, config.Keybinding{
		Key:     "J",
		Name:    "launch Herdr agent",
		Command: "launch-agent",
		Footer:  "agent",
	})
	ctx := &context.ProgramContext{
		Config:      &cfg,
		ScreenWidth: 160,
		View:        config.IssuesView,
	}
	ctx.Theme = theme.ParseTheme(ctx.Config)
	ctx.Styles = context.InitStyles(ctx.Theme)

	m := NewModel(ctx)
	m.SetIssueAction("archive")
	zone.Scan(m.View())
	require.Equal(t, "J", m.AgentKey())
	require.Contains(t, m.View(), "＋ n")

	for _, id := range []string{"new-issue", "launch-agent", "archive"} {
		require.Eventually(t, func() bool {
			return !zone.Get(id).IsZero()
		}, 250*time.Millisecond, time.Millisecond, "expected mouse zone %q", id)
	}
}

func TestViewUsesReboundIssueCreateKey(t *testing.T) {
	originalKeys := keys.IssueKeys.Create.Keys()
	originalHelp := keys.IssueKeys.Create.Help()
	keys.IssueKeys.Create.SetKeys("N")
	keys.IssueKeys.Create.SetHelp("N", originalHelp.Desc)
	defer func() {
		keys.IssueKeys.Create.SetKeys(originalKeys...)
		keys.IssueKeys.Create.SetHelp(originalHelp.Key, originalHelp.Desc)
	}()

	cfg, err := config.ParseConfig(config.Location{
		ConfigFlag:       "../../../config/testdata/test-config.yml",
		SkipGlobalConfig: true,
	})
	require.NoError(t, err)
	ctx := &context.ProgramContext{
		Config:      &cfg,
		ScreenWidth: 160,
		View:        config.IssuesView,
	}
	ctx.Theme = theme.ParseTheme(ctx.Config)
	ctx.Styles = context.InitStyles(ctx.Theme)

	view := NewModel(ctx).View()
	require.Contains(t, view, "＋ N")
	require.NotContains(t, view, "＋ n")
}

func TestAgentKeyBeforeConfigLoads(t *testing.T) {
	m := NewModel(&context.ProgramContext{})
	require.Empty(t, m.AgentKey())
}
