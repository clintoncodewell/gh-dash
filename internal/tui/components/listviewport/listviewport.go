package listviewport

import (
	"time"

	"charm.land/bubbles/v2/viewport"
	"charm.land/lipgloss/v2"

	"github.com/dlvhdr/gh-dash/v4/internal/tui/constants"
	"github.com/dlvhdr/gh-dash/v4/internal/tui/context"
	"github.com/dlvhdr/gh-dash/v4/internal/utils"
)

type Model struct {
	ctx             context.ProgramContext
	viewport        viewport.Model
	topBoundId      int
	bottomBoundId   int
	currId          int
	ListItemHeight  int
	NumCurrentItems int
	NumTotalItems   int
	LastUpdated     time.Time
	CreatedAt       time.Time
	ItemTypeLabel   string
}

func NewModel(
	ctx context.ProgramContext,
	dimensions constants.Dimensions,
	lastUpdated time.Time,
	createdAt time.Time,
	itemTypeLabel string,
	numItems, listItemHeight int,
) Model {
	model := Model{
		ctx:             ctx,
		NumCurrentItems: numItems,
		ListItemHeight:  listItemHeight,
		currId:          0,
		viewport: viewport.New(
			viewport.WithWidth(dimensions.Width),
			viewport.WithHeight(dimensions.Height),
		),
		topBoundId:    0,
		ItemTypeLabel: itemTypeLabel,
		LastUpdated:   lastUpdated,
		CreatedAt:     createdAt,
	}
	model.bottomBoundId = utils.Min(
		model.NumCurrentItems-1,
		model.getNumPrsPerPage()-1,
	)
	return model
}

func (m *Model) SetNumItems(numItems int) {
	m.NumCurrentItems = numItems
	m.bottomBoundId = utils.Min(m.NumCurrentItems-1, m.getNumPrsPerPage()-1)
}

func (m *Model) SetTotalItems(total int) {
	m.NumTotalItems = total
}

func (m *Model) SetItemHeight(height int) {
	m.ListItemHeight = height
}

func (m *Model) SyncViewPort(content string) {
	m.viewport.SetContent(content)
}

func (m *Model) getNumPrsPerPage() int {
	if m.ListItemHeight == 0 {
		return 0
	}
	return m.viewport.Height() / m.ListItemHeight
}

func (m *Model) ResetCurrItem() {
	m.currId = 0
	m.viewport.GotoTop()
}

func (m *Model) GetCurrItem() int {
	return m.currId
}

func (m *Model) SetCurrItem(item int) int {
	if m.NumCurrentItems == 0 {
		m.currId = 0
		return m.currId
	}

	m.currId = min(max(item, 0), m.NumCurrentItems-1)
	itemsPerPage := max(m.getNumPrsPerPage(), 1)
	if m.currId < m.topBoundId {
		m.topBoundId = m.currId
	} else if m.currId > m.bottomBoundId {
		m.topBoundId = m.currId - itemsPerPage + 1
	}
	m.topBoundId = max(m.topBoundId, 0)
	m.bottomBoundId = min(m.topBoundId+itemsPerPage-1, m.NumCurrentItems-1)
	m.viewport.SetYOffset(m.topBoundId * m.ListItemHeight)
	return m.currId
}

func (m *Model) ItemAtOffset(offset int) int {
	if offset < 0 || offset >= m.viewport.Height() || m.ListItemHeight <= 0 {
		return -1
	}
	item := m.topBoundId + offset/m.ListItemHeight
	if item < 0 || item >= m.NumCurrentItems {
		return -1
	}
	return item
}

func (m *Model) NextItem() int {
	atBottomOfViewport := m.currId >= m.bottomBoundId
	if atBottomOfViewport {
		m.topBoundId += 1
		m.bottomBoundId += 1
		m.viewport.ScrollDown(m.ListItemHeight)
	}

	newId := utils.Min(m.currId+1, m.NumCurrentItems-1)
	newId = utils.Max(newId, 0)
	m.currId = newId
	return m.currId
}

func (m *Model) PrevItem() int {
	if m.currId > 0 && m.currId <= m.topBoundId {
		m.topBoundId -= 1
		m.bottomBoundId -= 1
		m.viewport.ScrollUp(m.ListItemHeight)
	}

	m.currId = utils.Max(m.currId-1, 0)
	return m.currId
}

func (m *Model) FirstItem() int {
	m.currId = 0
	m.viewport.GotoTop()
	return m.currId
}

func (m *Model) LastItem() int {
	m.currId = m.NumCurrentItems - 1
	m.viewport.GotoBottom()
	return m.currId
}

func (m *Model) SetDimensions(dimensions constants.Dimensions) {
	m.viewport.SetHeight(max(0, dimensions.Height))
	m.viewport.SetWidth(max(0, dimensions.Width))
}

func (m *Model) View() string {
	viewport := m.viewport.View()
	return lipgloss.NewStyle().
		Width(m.viewport.Width()).
		MaxWidth(m.viewport.Width()).
		Render(
			viewport,
		)
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = *ctx
}
