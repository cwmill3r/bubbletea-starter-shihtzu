package mainmenu

// View renders the UI of the application.
func (m Model) View() string {
	return docStyle.Render(m.list.View())
}
