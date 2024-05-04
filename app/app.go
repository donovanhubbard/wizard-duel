package app

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	"time"
  "errors"
)

const (
	TICK_DURATION_MS = 100
)

type App struct {
	grid          *Grid
	teaPrograms   []*tea.Program
	pendingSpells map[string]CastSpellMsg
  players       []*Entity
}

func NewApp() App {
	g := NewGrid()
	(*g).PlaceCastles()
	return App{
		grid: g,
	}
}

func (a *App) Start() {
	log.Info("Game Engine starting")
	go func() {
		for true {
			time.Sleep(TICK_DURATION_MS * time.Millisecond)
			a.MoveAll()
			a.CastSpells()
			a.Send(GridUpdateMsg{*a.grid})
		}
	}()
}

func (a *App) FindPlayer(id string) (error,*Entity){
  for _, p := range a.players {
    if p.ID == id {
      return nil, p
    }
  }
  return errors.New(fmt.Sprintf("Could not find player id=%s",id)),nil
}

func (a *App) MoveAll() {
	oldGrid := a.grid
	nextGrid := NewGrid()
	for y, _ := range *oldGrid {
		for x, _ := range (*oldGrid)[y] {
			entity := (*oldGrid)[y][x]
			if entity != nil { //&& entity.NextMove != "" {
				Move(oldGrid, nextGrid, y, x)
			}
		}
	}
	a.grid = nextGrid
}

func (a *App) CastSpells() {
	for id, spellMsg := range a.pendingSpells {
		log.Info(fmt.Sprintf("Player %s is casting spell %s in diretion %s", id, spellMsg.Type, spellMsg.Direction))
		_, y, x, err := a.grid.FindEntity(id)

		log.Info(fmt.Sprintf("spellMsg.Type=%s spellMsg.Direction=%s", spellMsg.Type, spellMsg.Direction))
		log.Info(err)

		if err == nil {
			switch spellMsg.Type {
			case FIREBALL:
				switch spellMsg.Direction {
				case NORTH:
					if y > 0 {
						e := CreateFireball(NORTH)
						a.grid.SpawnEntity(e, y-1, x)
					}
				case SOUTH:
					if y < len(*a.grid)-1 {
						e := CreateFireball(SOUTH)
						a.grid.SpawnEntity(e, y+1, x)
					}
				case EAST:
					if x < len((*a.grid)[y])-1 {
						e := CreateFireball(EAST)
						a.grid.SpawnEntity(e, y, x+1)
					}
				case WEST:
					if x > 0 {
						e := CreateFireball(WEST)
						a.grid.SpawnEntity(e, y, x-1)
					}
				}
			}
		}
	}

	a.pendingSpells = make(map[string]CastSpellMsg)
}

func (a *App) Send(msg tea.Msg) {
	switch msg := msg.(type) {
	case GridUpdateMsg:
		for _, p := range a.teaPrograms {
			go p.Send(msg)
		}
	case PlanMoveMsg:
		entity, _, _, err := a.grid.FindEntity(msg.ID)
		if err != nil {
			log.Errorf(err.Error())
		} else {
			entity.NextMove = msg.Direction
		}
	case CastSpellMsg:
		a.pendingSpells[msg.ID] = msg
	case TryRespawnMsg:
		log.Debug(fmt.Sprintf("Trying to respawn %s", msg.ID))
    err, player := a.FindPlayer(msg.ID)
		if err != nil {
			log.Errorf(err.Error())
		} else {
			if player.IsDead {
				a.grid.PlacePlayer(player)
        player.IsDead = false
			}
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
	a.grid.PlacePlayer(player)
  a.players = append(a.players, player)
	model.ID = player.ID
	model.PC = player

	gu := GridUpdateMsg{Grid: *a.grid}
	a.Send(gu)

	options := bm.MakeOptions(s)
	options = append(options, tea.WithAltScreen())
	p := tea.NewProgram(model, options...)

	a.teaPrograms = append(a.teaPrograms, p)

	return p
}
