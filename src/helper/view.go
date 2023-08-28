package helper

import (
	"strings"
	
	"github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
)

var NameList []string
var checkBox = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("171"))

type Model struct {
	Cursor int
	Choice string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			m.Choice = NameList[m.Cursor]
			return m, tea.Quit

		case "down", "j", "s":
			m.Cursor++
			if m.Cursor >= len(NameList) {
				m.Cursor = 0
			}

		case "up", "k", "w":
			m.Cursor--
			if m.Cursor < 0 {
				m.Cursor = len(NameList) - 1
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	s := strings.Builder{}
	for i:=0; i<len(NameList); i++ {
		if m.Cursor == i {
			s.WriteString(checkBox.Render("[x] "))
		} else {
			s.WriteString("[ ] ")
		}
		s.WriteString(NameList[i])
		s.WriteString("\n")
	}
	s.WriteString("\nq: exit\n")
	return s.String()
}