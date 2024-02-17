package app

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	WIDTH  = 30
	HEIGHT = 20
)

type Grid [][]*Entity

func NewGrid() *Grid {
	var g *Grid
	grid := make(Grid, HEIGHT)
	g = &grid
	for i := range *g {
		(*g)[i] = make([]*Entity, WIDTH)
	}

	return g
}

func (g *Grid) PlaceCastles() {
	g.SpawnEntity(CreateCastle(), 4, 4)
	g.SpawnEntity(CreateCastle(), 14, 4)
	g.SpawnEntity(CreateCastle(), 4, 24)
	g.SpawnEntity(CreateCastle(), 14, 24)
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

func Move(oldGrid *Grid, nextGrid *Grid, y int, x int) {
	entity := (*oldGrid)[y][x]
	switch entity.NextMove {
	case NORTH:
		if y > 0 {
			if (*oldGrid)[y-1][x] == nil {
				(*nextGrid)[y-1][x] = entity
			} else if !entity.RemoveOnContact {
				(*nextGrid)[y][x] = entity
			}
		} else if !entity.RemoveOnContact {
			(*nextGrid)[y][x] = entity
		}
	case SOUTH:
		if y < len(*oldGrid)-1 {
			if (*oldGrid)[y+1][x] == nil {
				(*nextGrid)[y+1][x] = entity
			} else if !entity.RemoveOnContact {
				(*nextGrid)[y][x] = entity
			}
		} else if !entity.RemoveOnContact {
			(*nextGrid)[y][x] = entity
		}
	case EAST:
		if x < len((*oldGrid)[y])-1 {
			if (*oldGrid)[y][x+1] == nil {
				(*nextGrid)[y][x+1] = entity
			} else if !entity.RemoveOnContact {
				(*nextGrid)[y][x] = entity
			}
		} else if !entity.RemoveOnContact {
			(*nextGrid)[y][x] = entity
		}
	case WEST:
		if x > 0 {
			if (*oldGrid)[y][x-1] == nil {
				(*nextGrid)[y][x-1] = entity
			} else if !entity.RemoveOnContact {
				(*nextGrid)[y][x] = entity
			}
		} else if !entity.RemoveOnContact {
			(*nextGrid)[y][x] = entity
		}
	default:
		(*nextGrid)[y][x] = entity
	}
	switch entity.Type {
	case PLAYER:
		entity.NextMove = ""
	}
}

func (g *Grid) SpawnEntity(e Entity, y int, x int) {
	(*g)[y][x] = &e
}

func (g Grid) FindEntity(id string) (*Entity, int, int, error) {
	for y, _ := range g {
		for x, _ := range g[y] {
			entity := g[y][x]
			if entity != nil {
				if entity.ID == id {
					return entity, y, x, nil
				}
			}
		}
	}
	return nil, 0, 0, errors.New(fmt.Sprintf("Failed to find entity '%s'", id))
}
