package app

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"math/rand"
)

const (
	WIDTH  = 30
	HEIGHT = 20
)

type Grid [][]*Entity

func NewGrid() Grid {
	var g Grid
	g = make([][]*Entity, HEIGHT)
	for i := range g {
		g[i] = make([]*Entity, WIDTH)
	}

	return g
}

func (g *Grid) PlacePlayer(p *Entity) {
	var x, y int
	for true {
		y = rand.Intn(HEIGHT)
		x = rand.Intn(WIDTH)
		if (*g)[y][x] == nil {
			break
		}
	}
	(*g)[y][x] = p
}

func (g Grid) MoveAll() {
	for y, _ := range g {
		for x, _ := range g[y] {
			entity := g[y][x]
			if entity != nil { //&& entity.NextMove != "" {
				log.Infof("Type: %s ID: %s NextMove: %s", entity.Type, entity.ID, entity.NextMove)
				g.Move(entity, y, x)
			}
		}
	}
}

func (g Grid) Move(entity *Entity, y int, x int) {
	switch entity.NextMove {
	case NORTH:
		if y > 0 {
			if g[y-1][x] == nil {
				g[y-1][x] = entity
				g[y][x] = nil
			}
		}
	case SOUTH:
		if y < len(g)-1 {
			if g[y+1][x] == nil {
				g[y+1][x] = entity
				g[y][x] = nil
			}
		}
	case EAST:
		if x < len(g[y])-1 {
			if g[y][x+1] == nil {
				g[y][x+1] = entity
				g[y][x] = nil
			}
		}
	case WEST:
		if x > 0 {
			if g[y][x-1] == nil {
				g[y][x-1] = entity
				g[y][x] = nil
			}
		}
	}
	entity.NextMove = ""
}

func (g Grid) FindEntity(id string) (*Entity, error) {
	for y, _ := range g {
		for x, _ := range g[y] {
			entity := g[y][x]
			if entity != nil {
				if entity.ID == id {
					return entity, nil
				}
			}
		}
	}
	return nil, errors.New(fmt.Sprintf("Failed to find entity '%s'", id))
}
