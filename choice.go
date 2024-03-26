package cli

// A simple example that shows how to retrieve a value from a Bubble Tea
// program after the Bubble Tea has exited.

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type choiceModel struct {
	cursor          int
	question        string
	choices         []string
	choice          string
	showQuitMessage bool
}

func (m choiceModel) Init() tea.Cmd {
	return nil
}

func (m choiceModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			m.choice = m.choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		}
	}

	return m, nil
}

func (m choiceModel) View() string {
	s := strings.Builder{}
	s.WriteString(titleStyle.Render(m.question))
	s.WriteString("\n")

	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(m.choices[i])
		s.WriteString("\n")
	}
	if m.showQuitMessage {
		s.WriteString("\n(press q to quit)\n")
	}

	return s.String()
}

func Choice(question string, choices []string, showQuitMessage bool) string {
	p := tea.NewProgram(choiceModel{
		question:        question,
		choices:         choices,
		showQuitMessage: showQuitMessage,
	})
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}
	if m, ok := m.(choiceModel); ok && m.choice != "" {
		return m.choice
	}
	return ""
}
