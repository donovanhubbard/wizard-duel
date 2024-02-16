package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	"time"
)

type App struct {
	grid        *Grid
	teaPrograms []*tea.Program
}

func NewApp() App {
	g := NewGrid()
	return App{
		grid: &g,
	}
}

func (a *App) Start() {
	log.Info("Game Engine starting")
	go func() {
		for true {
			time.Sleep(1 * time.Second)
			log.Info("tick")
			a.grid.MoveAll()
			a.Send(GridUpdateMsg{*a.grid})
		}
	}()
}

func (a *App) Send(msg tea.Msg) {
	switch msg := msg.(type) {
	case GridUpdateMsg:
		for _, p := range a.teaPrograms {
			go p.Send(msg)
		}
	case PlanMoveMsg:
		log.Infof("ID: %s, Direction: %s", msg.ID, msg.Direction)

		entity, err := a.grid.FindEntity(msg.ID)

		if err != nil {
			log.Errorf(err.Error())
		} else {
			entity.NextMove = msg.Direction
		}
	}
}

func (a *App) ProgramHandler(s ssh.Session) *tea.Program {
	if _, _, active := s.Pty(); !active {
		wish.Fatalln(s, "terminal is not active")
	}

	var model Model
	model.App = a
	model.Grid = *a.grid
	player := CreateNextPlayer()
	a.grid.PlacePlayer(&player)
	model.ID = player.ID

	gu := GridUpdateMsg{Grid: *a.grid}
	a.Send(gu)

	options := bm.MakeOptions(s)
	options = append(options, tea.WithAltScreen())
	p := tea.NewProgram(model, options...)

	a.teaPrograms = append(a.teaPrograms, p)

	return p
}
