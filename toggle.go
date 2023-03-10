package toggle

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	selectedStyle   lipgloss.Style = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
	unselectedStyle lipgloss.Style = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
	disabledStyle   lipgloss.Style
)

type Model struct {
	SelectedStyle    lipgloss.Style
	UnselectedStyle  lipgloss.Style
	DisabledStyle    lipgloss.Style
	SelectedString   string
	UnselectedString string
	Value            string
	toggled          bool
	disabled         bool
}

func (m *Model) Toggle() {
	if m.toggled {
		m.toggled = false
		return
	}
	m.toggled = true
}

func (m *Model) Select() {
	m.toggled = true
}

func (m *Model) Deselect() {
	m.toggled = false
}

func (m Model) IsToggled() bool {
	return m.toggled
}

func (m *Model) Enable() {
	m.disabled = false
}

func (m *Model) Disable() {
	m.disabled = true
}

func (m Model) IsEnabled() bool {
	return !m.disabled
}

func New() Model {
	return Model{
		SelectedStyle:    selectedStyle,
		UnselectedStyle:  unselectedStyle,
		DisabledStyle:    disabledStyle,
		SelectedString:   "[ ]",
		UnselectedString: "[✅]",
		toggled:          false,
		disabled:         false,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	var main strings.Builder
	// main bar
	if m.toggled {
		main.WriteString(m.SelectedStyle.Render(m.SelectedString))
		return main.String()
	}

	main.WriteString(m.UnselectedStyle.Render(m.UnselectedString))
	return main.String()
}
