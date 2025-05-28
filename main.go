package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	cursorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	checkedStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10"))
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	adding   bool
	input    string
}

func initialModel() model {
	return model{
		choices:  []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if m.adding {
			switch msg.Type {
			case tea.KeyEnter:
				trimmed := strings.TrimSpace(m.input)
				if trimmed != "" {
					m.choices = append(m.choices, trimmed)
					m.cursor = len(m.choices) - 1
				}
				m.input = ""
				m.adding = false
				return m, nil
			case tea.KeyEsc:
				m.input = ""
				m.adding = false
				return m, nil
			case tea.KeyBackspace:
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}
			default:
				m.input += msg.String()
			}
			return m, nil
		}

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "a":
			m.adding = true
			m.input = ""
		case "d":
			if len(m.selected) > 0 {
				newChoices := []string{}
				for i, choice := range m.choices {
					if _, selected := m.selected[i]; !selected {
						newChoices = append(newChoices, choice)
					}
				}
				m.choices = newChoices
				m.selected = make(map[int]struct{}) // clear selection
				if m.cursor >= len(m.choices) {
					m.cursor = len(m.choices) - 1
					if m.cursor < 0 {
						m.cursor = 0
					}
				}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.adding {
		return fmt.Sprintf("Add an item: %s\n\n(Enter to confirm, Esc to cancel)", m.input)
	}

	s := "What should we buy at the market?\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = cursorStyle.Render(">")
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = checkedStyle.Render("x")
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\n↑/↓: move • space/enter: toggle • a: add • d: delete • q: quit\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v\n", err)
		os.Exit(1)
	}

	m := finalModel.(model)

	fmt.Println("\nYou selected:")
	for i := range m.selected {
		if i >= 0 && i < len(m.choices) {
			fmt.Println("-", m.choices[i])
		}
	}
}
