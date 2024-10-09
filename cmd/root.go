package cmd

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cwmill3r/shihtzu/internal/config"
	"github.com/cwmill3r/shihtzu/internal/tui"
)

var rootCmd = &cobra.Command{
	Use:     "shihtzu",
	Short:   "shihtzu is a smol testing framework for people and llms",
	Version: "0.0.1",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Call LoadConfig to load and possibly create the config
		if err := config.LoadConfig(); err != nil {
			log.Fatalf("Error loading config: %v", err)
		}

		// If logging is enabled, logs will be output to debug.log.
		if viper.GetBool("settings.enable_logging") {
			f, err := tea.LogToFile("debug.log", "debug")
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			defer func() {
				if err = f.Close(); err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
			}()
		}

		// Initialize TUI
		model := tui.NewModel()
		var opts []tea.ProgramOption

		// Always append alt screen program option
		opts = append(opts, tea.WithAltScreen())

		// If mousewheel is enabled, append it to the program options
		if viper.GetBool("settings.enable_mousewheel") {
			opts = append(opts, tea.WithMouseAllMotion())
		}

		// Initialize new app
		p := tea.NewProgram(model, opts...)
		if _, err := p.Run(); err != nil {
			log.Fatal("Failed to start shihtzu", err)
			os.Exit(1)
		}
	},
}

// Execute executes the root command which starts the application.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
