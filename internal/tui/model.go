package tui

import (
	"github.com/cwmill3r/shihtzu/keys"
	"github.com/cwmill3r/shihtzu/mainmenu"
)

// Model holds the global TUI state
type Model struct {
	mainMenu      mainmenu.Model // Bubble for main menu
	currentBubble string         // Tracks the current active bubble (e.g., "mainmenu")
	keyMap        keys.KeyMap
}

// NewModel creates and returns a new Model instance
func NewModel() Model {
	return Model{
		mainMenu:      mainmenu.New(), // Initialize the main menu model
		currentBubble: "mainmenu",     // Set the active bubble to main menu initially
		keyMap:        keys.DefaultKeyMap(),
	}
}
