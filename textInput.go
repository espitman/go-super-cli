package cli

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
)

type (
	textInputErrMsg error
)

type textInputModel struct {
	textInput textinput.Model
	question  string
	showQuit  bool
	err       error
}

func initialTextInputModel(question string, placeholder string, showQuit bool) textInputModel {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return textInputModel{
		textInput: ti,
		question:  question,
		showQuit:  showQuit,
		err:       nil,
	}
}

func (m textInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m textInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {

		case tea.KeyEnter:
			return m, tea.Quit

		case tea.KeyCtrlC, tea.KeyEsc:
			os.Exit(1)
		}

	case textInputErrMsg:
		m.err = msg
		return m, nil
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m textInputModel) View() string {
	quiteMsg := ""
	if m.showQuit {
		quiteMsg = "(esc to quit)"
	}

	q := titleStyle.Render(m.question)
	return fmt.Sprintf(
		q+"\n%s\n%s",
		m.textInput.View(),
		quiteMsg,
	)
}

func TextInput(question string, placeholder string, showQuit bool) string {
	p := tea.NewProgram(initialTextInputModel(question, placeholder, showQuit))
	m, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
	return m.(textInputModel).textInput.Value()
}
