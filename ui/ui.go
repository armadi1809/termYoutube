package ui

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"fmt"
	"log"
	"strings"

	youtubeapi "github.com/armadi1809/termYoutube/youtubeApi"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type model struct {
	textInput     textinput.Model
	youtubeClient *youtubeapi.YoutubeApiClient
	err           error
}

func initialModel(c *youtubeapi.YoutubeApiClient) model {
	ti := textinput.New()
	ti.Placeholder = "Pikachu"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput:     ti,
		youtubeClient: c,
		err:           nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.youtubeClient.Search(strings.ReplaceAll(m.textInput.Value(), " ", "+"))
			return m, nil
		case tea.KeyCtrlC, tea.KeyEsc:
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

func (m model) View() string {
	return fmt.Sprintf(
		"What do you want to listen to?\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func Start(apiClient *youtubeapi.YoutubeApiClient) {
	p := tea.NewProgram(initialModel(apiClient))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
