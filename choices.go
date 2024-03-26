package cli

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type choicesModel struct {
	question        string
	choices         []string
	cursor          int
	selected        map[int]struct{}
	showQuitMessage bool
}

func initialChoicesModel(question string, choices []string, showQuitMessage bool) choicesModel {
	return choicesModel{
		question:        question,
		choices:         choices,
		selected:        make(map[int]struct{}),
		showQuitMessage: showQuitMessage,
	}
}

func (m choicesModel) Init() tea.Cmd {
	return nil
}

func (m choicesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			os.Exit(1)
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "enter":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m choicesModel) View() string {
	s := titleStyle.Render(m.question)
	s += "\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += helpStyle.Render("Use space to select...") // Hello, kitty.

	if m.showQuitMessage {
		s += "\nPress q to quit.\n"
	}
	return s
}

func Choices(question string, choices []string, showQuitMessage bool) ([]int, []string) {
	p := tea.NewProgram(initialChoicesModel(question, choices, showQuitMessage))

	state, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	selected := state.(choicesModel).selected
	selectedItems := make([]string, 0, len(selected))
	var selectedIndexes []int
	for index := range selected {
		selectedIndexes = append(selectedIndexes, index)
		selectedItems = append(selectedItems, state.(choicesModel).choices[index])
	}
	return selectedIndexes, selectedItems
}
