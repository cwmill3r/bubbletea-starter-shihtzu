package mainmenu

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model holds the state for the main menu
type Model struct {
	cursor    int
	menuItems []string
	styles    styles
}

type styles struct {
	title        lipgloss.Style
	item         lipgloss.Style
	selectedItem lipgloss.Style
}

// MoveUp moves the cursor up
func (m *Model) MoveUp() {
	if m.cursor > 0 {
		m.cursor--
	}
}

// MoveDown moves the cursor down
func (m *Model) MoveDown() {
	if m.cursor < len(m.menuItems)-1 {
		m.cursor++
	}
}

// Select handles the selection of a menu item
func (m *Model) Select() tea.Cmd {
	// Logic for menu item selection, e.g., switch to another bubble
	return nil
}

// New creates a new main menu model
func New() Model {
	return Model{
		menuItems: []string{"Start Test", "Settings", "Exit"},
		styles: styles{
			title:        lipgloss.NewStyle().Bold(true).Underline(true),
			item:         lipgloss.NewStyle(),
			selectedItem: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205")),
		},
	}
}
