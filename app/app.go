package app

import (
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
  tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/donovanhubbard/wizard-duel/entities"
)

type App struct {
  grid *entities.Grid
  teaPrograms []*tea.Program
}

func NewApp() App{
  g := entities.NewGrid()
  return App{
    grid: &g,
  }
}

func (a *App) Send(msg tea.Msg) {
	for _, p := range a.teaPrograms {
		go p.Send(msg)
	}
}

func (a *App) ProgramHandler(s ssh.Session) *tea.Program {
	if _, _, active := s.Pty(); !active {
		wish.Fatalln(s, "terminal is not active")
	}

	var model Model
  model.App = a
  model.Grid = *a.grid
  player := entities.CreateNextPlayer()
  a.grid.PlacePlayer(&player)

  gu := GridUpdateMsg{Grid:*a.grid}
  a.Send(gu)

  options := bm.MakeOptions(s)
  options = append(options,tea.WithAltScreen())
	p := tea.NewProgram(model, options...)

  a.teaPrograms = append(a.teaPrograms, p)

	return p
}
