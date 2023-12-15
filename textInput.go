package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type inputModel struct {
	textInput textinput.Model
	err       error
	header    string
	footer    string
}

func initialModel() inputModel {
	data := initialize()
	ti := textinput.New()
	ti.Placeholder = data.Steps[Input].placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return inputModel{
		textInput: ti,
		err:       nil,
		header: data.Steps[Input].header,
		footer: data.Steps[Input].footer,
	}
}

func (m inputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m inputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m inputModel) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		m.header,
		m.textInput.View(),
		m.footer,
	) + "\n"
}

