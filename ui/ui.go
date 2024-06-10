package ui

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"fmt"
	"log"
	"strings"

	youtubeapi "github.com/armadi1809/termYoutube/youtubeApi"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

var columns = []table.Column{
	{Title: "Title", Width: 60},
	{Title: "Description", Width: 60},
}

type model struct {
	textInput     textinput.Model
	youtubeClient *youtubeapi.YoutubeApiClient
	table         table.Model
	err           error
}

func initialModel(c *youtubeapi.YoutubeApiClient) model {
	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(false),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)
	ti := textinput.New()
	ti.Placeholder = "Song of the year"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput:     ti,
		youtubeClient: c,
		err:           nil,
		table:         t,
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
			searchRes := m.youtubeClient.Search(strings.ReplaceAll(m.textInput.Value(), " ", "+"))
			newRows := []table.Row{}
			for _, item := range searchRes.Items {
				newRows = append(newRows, table.Row{item.Snippet.Title, item.Snippet.Description})
			}
			m.table.SetRows(newRows)
			if m.table.Focused() {
				m.table.Blur()
				m.textInput.Focus()
			} else {
				m.table.Focus()
				m.textInput.Blur()
			}
			return m, nil
		case tea.KeyEsc:
			if m.table.Focused() {
				m.table.Blur()
				m.textInput.Focus()
			} else {
				m.table.Focus()
				m.textInput.Blur()
			}

		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}
	if m.table.Focused() {
		m.table, cmd = m.table.Update(msg)
	} else {
		m.textInput, cmd = m.textInput.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"\n%s\n%s\n\n%s\n\n%s\n",
		"What do you want to listen to?",
		m.textInput.View(),
		"(esc to quit)",
		baseStyle.Render(m.table.View()),
	)
}

func Start(apiClient *youtubeapi.YoutubeApiClient) {
	p := tea.NewProgram(initialModel(apiClient))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
