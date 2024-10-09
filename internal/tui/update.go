package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.currentBubble {
	case "mainmenu":
		m.mainMenu, cmd = m.mainMenu.Update(msg, m.keyMap)
	}

	return m, cmd
}
