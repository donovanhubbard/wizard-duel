package tui

import (
  "strings"
	tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/lipgloss"
  "github.com/enescakir/emoji"
	"github.com/donovanhubbard/wizard-duel/entities"
)

const (
  PURPLE = lipgloss.Color("13")
  GREEN  = lipgloss.Color("10")
)

type Model struct{
  Grid entities.Grid
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {

  var sb strings.Builder
  for i, _ := range m.Grid{
    for j, _ := range m.Grid[i]{
      entity := m.Grid[i][j]
      entity.Render(&sb)
    }
    sb.WriteString("\n")
  }

  defaultStyle := lipgloss.NewStyle().
                  BorderStyle(lipgloss.RoundedBorder()).
		              BorderForeground(PURPLE).
                  Foreground(GREEN)


  header := defaultStyle.SetString(emoji.Parse(":mage: Wizard Duel :mage:")).String()

  grid := defaultStyle.
            PaddingLeft(1).
            PaddingRight(1).
            SetString(sb.String()).
            String()

  text := lipgloss.JoinVertical(lipgloss.Center, header, grid)

  return text
}
