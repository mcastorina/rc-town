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
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(black)).
		Background(lipgloss.Color(white)).
		Width(28).
		Height(17)
	var blackBorder = lipgloss.NewStyle().
		Background(lipgloss.Color(black)).
		Width(24).
		Height(10).
		MarginTop(1).
		MarginLeft(2)
	var screen = lipgloss.NewStyle().
		Background(lipgloss.Color(white)).
		Width(20).
		Height(8).
		MarginTop(1).
		MarginLeft(2)
	var blackScreen = lipgloss.NewStyle().
		Background(lipgloss.Color(black)).
		Width(16).
		Height(6).
		MarginTop(1).
		MarginLeft(2)
	var greenBit = lipgloss.NewStyle().
		Background(lipgloss.Color(green)).
		Width(2).
		Height(1).
		MarginTop(1).
		MarginLeft(0).Render("")
	// var blackBit = lipgloss.NewStyle().
	// 	Background(lipgloss.Color(green)).
	// 	Width(2).
	// 	Height(1).
	// 	MarginTop(1).
	// 	MarginLeft(0).Render("")

	greenBits := lipgloss.JoinHorizontal(lipgloss.Top, greenBit, "  ", greenBit, "  ", greenBit)
	otherGreenBits := lipgloss.JoinHorizontal(lipgloss.Top, "  ", greenBit, greenBit, "  ", greenBit, greenBit)
	greenBits = lipgloss.JoinVertical(lipgloss.Top, greenBits, otherGreenBits)

	return style.Render(blackBorder.Render(screen.Render(blackScreen.Render(greenBits))))
}

func main() {
	// flag.Parse()
	runCli()
}

func runCli() {
	model := &model{}
	program := tea.NewProgram(model, teaOptions...)

	exitCode := 0
	if err := program.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "clidle: %s\n", err)
		exitCode = 1
	}
	os.Exit(exitCode)
}
