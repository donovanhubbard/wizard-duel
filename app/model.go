package app

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/enescakir/emoji"
	"strings"
)

const (
	PURPLE = lipgloss.Color("13")
	GREEN  = lipgloss.Color("10")
)

type Model struct {
	Grid Grid
	App  *App
	ID   string
	PC   *Entity
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
		case "a":
			msg := PlanMoveMsg{
				Direction: WEST,
				ID:        m.ID,
			}
			m.App.Send(msg)
		case "w":
			msg := PlanMoveMsg{
				Direction: NORTH,
				ID:        m.ID,
			}
			m.App.Send(msg)
		case "d":
			msg := PlanMoveMsg{
				Direction: EAST,
				ID:        m.ID,
			}
			m.App.Send(msg)
		case "s":
			msg := PlanMoveMsg{
				Direction: SOUTH,
				ID:        m.ID,
			}
			m.App.Send(msg)
		case "up":
			msg := CastSpellMsg{
				ID:        m.ID,
				Type:      FIREBALL,
				Direction: NORTH,
			}
			m.App.Send(msg)
		case "down":
			msg := CastSpellMsg{
				ID:        m.ID,
				Type:      FIREBALL,
				Direction: SOUTH,
			}
			m.App.Send(msg)
		case "right":
			msg := CastSpellMsg{
				ID:        m.ID,
				Type:      FIREBALL,
				Direction: EAST,
			}
			m.App.Send(msg)
		case "left":
			msg := CastSpellMsg{
				ID:        m.ID,
				Type:      FIREBALL,
				Direction: WEST,
			}
			m.App.Send(msg)
		}
	case GridUpdateMsg:
		m.Grid = msg.Grid
	}
	return m, nil
}

func (m Model) View() string {

	var sb strings.Builder
	for i, _ := range m.Grid {
		for j, _ := range m.Grid[i] {
			entity := m.Grid[i][j]
			entity.Render(&sb)
		}
		sb.WriteString("\n")
	}

	defaultStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(PURPLE).
		Foreground(GREEN)

	title := defaultStyle.SetString(emoji.Parse(":mage: Wizard Duel :mage:")).String()
	health := defaultStyle.Align(lipgloss.Right).SetString(fmt.Sprintf("%v %d", emoji.RedHeart, m.PC.Health)).String()
	header := lipgloss.JoinHorizontal(lipgloss.Center, title, health)

	grid := defaultStyle.
		PaddingLeft(1).
		PaddingRight(1).
		SetString(sb.String()).
		String()

	move := defaultStyle.SetString("Move: WASD").String()
	shoot := defaultStyle.SetString(fmt.Sprintf("%v: Arrows", emoji.Fire)).String()
	footer := lipgloss.JoinHorizontal(lipgloss.Center, move, shoot)

	text := lipgloss.JoinVertical(lipgloss.Center, header, grid, footer)

	return text
}
