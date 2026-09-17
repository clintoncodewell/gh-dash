package carousel

import (
	"fmt"
	"testing"
	"time"

	zone "github.com/lrstanley/bubblezone/v2"
)

func TestWithZonePrefixMarksVisibleItemsByOriginalIndex(t *testing.T) {
	zone.NewGlobal()
	zone.SetEnabled(true)

	m := New(
		WithItems([]string{"one", "two", "three"}),
		WithWidth(80),
		WithHeight(1),
		WithZonePrefix("tab"),
	)
	zone.Scan(m.View())

	for index := range m.Items() {
		id := fmt.Sprintf("tab-%d", index)
		if !waitForZone(id) {
			t.Errorf("expected mouse zone %q", id)
		}
	}
}

func waitForZone(id string) bool {
	deadline := time.Now().Add(250 * time.Millisecond)
	for time.Now().Before(deadline) {
		if !zone.Get(id).IsZero() {
			return true
		}
		time.Sleep(time.Millisecond)
	}
	return false
}
