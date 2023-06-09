package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type response struct {
	Msg string `json:"Msg"`
}

type model struct {
	username        string
	ircChatroomName string
	answerField     textinput.Model
	spinner         spinner.Model
	isInputDone     bool
	Resp            string
}

func New() *model {
	answerField := textinput.New()
	answerField.Focus()
	answerField.Width = 20
	s := spinner.New()
	s.Spinner = spinner.Dot
	return &model{answerField: answerField, spinner: s, isInputDone: false}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) postToServer() tea.Cmd {
	return func() tea.Msg {
		jsonBody := []byte(`{"Msg": "Hello from client"}`)
		requestBody := bytes.NewReader(jsonBody)
		resp, err := http.Post("http://localhost:3000/woop", "application/json", requestBody)
		if err != nil {
			return err
		}
		fmt.Println("Request made")

		content, err := ioutil.ReadAll(resp.Body)
		strBody := string(content)
		m.Resp = strBody

		defer resp.Body.Close()
		return strBody
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "d":
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
				m.isInputDone = true
				return m, tea.Batch(
					spinner.Tick,
					m.postToServer(),
				)
			}
		}
	case string:
		m.Resp = msg
		return m, spinner.Tick
	}
	m.answerField, cmd = m.answerField.Update(msg)
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m model) View() string {
	s := ""
	if !m.isInputDone {
		if m.username == "" {
			return fmt.Sprintf("Enter a username \n %s", m.answerField.View())
		}
		if m.ircChatroomName == "" {
			return fmt.Sprintf("Username: %s\nEnter a chatroom name \n %s", m.username, m.answerField.View())
		}
	}
	s += fmt.Sprintf("%s %s connecting to %s\n", m.spinner.View(), m.username, m.ircChatroomName)
	s += "Press d to disconnect"
	if m.Resp != "" {
		return fmt.Sprintf("Response from server: %s\n", m.Resp)
	}
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
