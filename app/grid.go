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
	var nextY, nextX int
	entity := (*oldGrid)[y][x]

	switch entity.NextMove {
	case NORTH:
		if y > 0 {
			nextY = y - 1
		} else {
			nextY = y
		}
		nextX = x
	case SOUTH:
		if y < len(*oldGrid)-1 {
			nextY = y + 1
		} else {
			nextY = y
		}
		nextX = x
	case WEST:
		if x > 0 {
			nextX = x - 1
		}
		nextY = y
	case EAST:
		if x < len((*oldGrid)[y])-1 {
			nextX = x + 1
		} else {
			nextX = x
		}
		nextY = y
	default:
		nextY = y
		nextX = x
	}

	nextEntityOldGrid := (*oldGrid)[nextY][nextX]
	nextEntityNextGrid := (*nextGrid)[nextY][nextX]

	if nextEntityOldGrid == nil && nextEntityNextGrid == nil{
		(*nextGrid)[nextY][nextX] = entity
	} else {
		if !entity.RemoveOnContact {
			(*nextGrid)[y][x] = entity
		}
    if entity.Damage > 0 {
      if nextEntityOldGrid != nil && !nextEntityOldGrid.Indestructible{
			  nextEntityOldGrid.DealDamage(entity.Damage)
      }else if nextEntityNextGrid != nil && !nextEntityNextGrid.Indestructible{
			  nextEntityNextGrid.DealDamage(entity.Damage)
      }
    }
	}

	switch entity.Type {
	case PLAYER:
		entity.NextMove = ""
	}
}

func (g *Grid) SpawnEntity(e *Entity, y int, x int) {
	(*g)[y][x] = e
}

func (g *Grid) FindEntity(id string) (*Entity, int, int, error) {
	for y, _ := range *g {
		for x, _ := range (*g)[y] {
			entity := (*g)[y][x]
			if entity != nil {
				if entity.ID == id {
					return entity, y, x, nil
				}
			}
		}
	}
	return nil, 0, 0, errors.New(fmt.Sprintf("Failed to find entity '%s'", id))
}
