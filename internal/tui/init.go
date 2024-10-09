package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Init initializes the TUI with necessary components.
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.mainMenu.Init(),            // Initialize the main menu
		tea.SetWindowTitle("My App"), // Set the window title
		// You can add more initialization commands here if needed
	)
}
