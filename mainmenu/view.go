package mainmenu

func (m Model) View() string {
	var s string
	s += m.styles.title.Render("Main Menu") + "\n\n"

	for i, item := range m.menuItems {
		if i == m.cursor {
			s += m.styles.selectedItem.Render(item) + "\n"
		} else {
			s += m.styles.item.Render(item) + "\n"
		}
	}

	return s
}
