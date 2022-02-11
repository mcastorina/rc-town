package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	teaOptions = []tea.ProgramOption{tea.WithAltScreen(), tea.WithOutput(os.Stderr)}
	black      = "#231f20"
	green      = "#5aba47"
	white      = "#ffffff"
)

type model struct {
	cursor int
	// color grid
	// TODO: more data than just color
	grid [][]string
}

func (m *model) Init() tea.Cmd {
	return nil
}
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		// The "down" and "j" keys move the cursor down
		case "down", "j":
			m.cursor++
		}
	}
	return m, nil
}

func (m *model) View() string {
	var image string

	for _, row := range m.grid {
		for _, color := range row {
			style := lipgloss.NewStyle().Background(lipgloss.Color(color)).Width(2).Height(1)
			image += style.Render("")
		}
		image += "\n"
	}

	return image
}

func main() {
	// flag.Parse()
	runCli()
}

func runCli() {
	model := &model{
		grid: [][]string{
			{white, white, white, white, white, white, white, white, white, white, white, white, white, white},
			{white, black, black, black, black, black, black, black, black, black, black, black, black, white},
			{white, black, white, white, white, white, white, white, white, white, white, white, black, white},
			{white, black, white, black, black, black, black, black, black, black, black, white, black, white},
			{white, black, white, green, black, green, black, green, black, black, black, white, black, white},
			{white, black, white, black, black, black, black, black, black, black, black, white, black, white},
			{white, black, white, black, green, green, black, green, green, black, black, white, black, white},
			{white, black, white, black, black, black, black, black, black, black, black, white, black, white},
			{white, black, white, black, black, black, black, black, black, black, black, white, black, white},
			{white, black, white, white, white, white, white, white, white, white, white, white, black, white},
			{white, black, black, black, black, black, black, black, black, black, black, black, black, white},
			{white, white, white, white, white, black, black, black, black, white, white, white, white, white},
			{white, white, black, black, black, black, black, black, black, black, black, black, white, white},
			{white, black, black, black, white, black, white, black, white, black, white, black, black, white},
			{white, black, black, white, black, white, black, white, black, white, black, black, black, white},
			{white, black, black, black, black, black, black, black, black, black, black, black, black, white},
			{white, white, white, white, white, white, white, white, white, white, white, white, white, white},
		},
	}

	program := tea.NewProgram(model, teaOptions...)

	exitCode := 0
	if err := program.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "clidle: %s\n", err)
		exitCode = 1
	}
	os.Exit(exitCode)
}
