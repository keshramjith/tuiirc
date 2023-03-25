package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type model struct {
	username        string
	ircChatroomName string
	answerField     textinput.Model
}

func New() *model {
	answerField := textinput.New()
	answerField.Focus()
	answerField.Width = 20
	return &model{answerField: answerField}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.username == "" {
				m.username = m.answerField.Value()
				m.answerField.Reset()
				return m, cmd
			}
			if m.ircChatroomName == "" {
				m.ircChatroomName = m.answerField.Value()
				m.answerField.Reset()
				return m, cmd
			}
		}
	}
	m.answerField, cmd = m.answerField.Update(msg)
	return m, cmd
}

func (m model) View() string {
	s := ""
	if m.username == "" {
		return fmt.Sprintf("Enter a username \n %s", m.answerField.View())
	}
	s += fmt.Sprintf("Username: %s\n", m.username)
	if m.ircChatroomName == "" {
		return fmt.Sprintf("Enter a chatroom name \n %s", m.answerField.View())
	}
	s += fmt.Sprintf("Chatroom: %s\n", m.ircChatroomName)
	s += "Press q to quit"
	return s
}

func main() {
	m := New()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
