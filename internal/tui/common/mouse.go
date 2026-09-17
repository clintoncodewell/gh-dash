package common

import (
	tea "charm.land/bubbletea/v2"
	zone "github.com/lrstanley/bubblezone/v2"
)

// MarkMouseZone keeps components usable in isolated tests that do not
// initialize bubblezone's global manager.
func MarkMouseZone(id, content string) string {
	if zone.DefaultManager == nil {
		return content
	}
	return zone.Mark(id, content)
}

func MouseZoneInBounds(id string, msg tea.MouseMsg) bool {
	if zone.DefaultManager == nil {
		return false
	}
	return zone.Get(id).InBounds(msg)
}

func MouseZonePosition(id string, msg tea.MouseMsg) (int, int) {
	if zone.DefaultManager == nil {
		return -1, -1
	}
	return zone.Get(id).Pos(msg)
}
