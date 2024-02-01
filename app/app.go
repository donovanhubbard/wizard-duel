package app

import (
	tea "github.com/charmbracelet/bubbletea"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/donovanhubbard/wizard-duel/entities"
	"github.com/donovanhubbard/wizard-duel/tui"
)

type App struct {
  grid entities.Grid
  teaPrograms []*tea.Program
}

func NewApp() App{
  return App{
    grid: entities.NewGrid(),
  }
}

func (a *App) ProgramHandler(s ssh.Session) *tea.Program {
	if _, _, active := s.Pty(); !active {
		wish.Fatalln(s, "terminal is not active")
	}

	var model tui.Model
  model.Grid = a.grid
  options := bm.MakeOptions(s)
  options = append(options,tea.WithAltScreen())
	p := tea.NewProgram(model, options...)

  a.teaPrograms = append(a.teaPrograms, p)

	return p
}
