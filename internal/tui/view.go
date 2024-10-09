package tui

func (m Model) View() string {
	switch m.currentBubble {
	case "mainmenu":
		return m.mainMenu.View()
		// Add other bubbles like testRunner here
	default:
		return ""
	}
}
