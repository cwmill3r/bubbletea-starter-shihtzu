package mainmenu

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cwmill3r/shihtzu/keys"
)

// Update handles messages and updates the main menu bubble
func (m Model) Update(msg tea.Msg, keyMap keys.KeyMap) (Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keyMap.Up):
			m.MoveUp()
		case key.Matches(msg, keyMap.Down):
			m.MoveDown()
		case key.Matches(msg, keyMap.Select):
			cmds = append(cmds, m.Select())
		}
	}

	return m, tea.Batch(cmds...)
}
