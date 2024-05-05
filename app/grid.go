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

func getNextCoordinate(grid *Grid, entity *Entity, y int, x int)(int,int){
	var nextY, nextX int

	switch entity.NextMove {
	case NORTH:
		nextY = y - 1
		nextX = x
	case SOUTH:
		nextY = y + 1
		nextX = x
	case WEST:
		nextX = x - 1
		nextY = y
	case EAST:
		nextX = x + 1
		nextY = y
	default:
		nextY = y
		nextX = x
	}

  return nextY,nextX
}

func trimNextCoordinate(grid *Grid, y int, x int)(int,int){
  var nextY, nextX int
  if(y < 0){
    nextY = 0
  }else if y >= len(*grid){
    nextY = len(*grid)-1
  } else {
    nextY = y
  }
  if(x < 0){
    nextX = 0
  } else if x >= len((*grid)[nextY]){
    nextX = len((*grid)[nextY]) - 1
  }else {
    nextX = x
  }

  return nextY, nextX
}

func isOutOfBounds(grid *Grid, y int, x int) bool {
  if y < 0 {
    return true
  }
  if y >= len(*grid) {
    return true
  }

  if x < 0 {
    return true
  }else if x >= len((*grid)[y]) {
    return true
  }
  return false
}

func Move(oldGrid *Grid, nextGrid *Grid, y int, x int) {
	entity := (*oldGrid)[y][x]
  if entity == nil {
    return
  }

  if entity.IsDead {
	 (*oldGrid)[y][x] = nil
   return
  }

  if isOutOfBounds(oldGrid,y,x) {
    if(entity.RemoveOnContact){
      (*oldGrid)[y][x] = nil
      return
    }
  }
  nextY, nextX := getNextCoordinate(oldGrid, entity, y, x)
  if isOutOfBounds(oldGrid,nextY,nextX) {
    if(entity.RemoveOnContact){
      (*oldGrid)[y][x] = nil
      return
    }else{
      nextY, nextX = trimNextCoordinate(oldGrid,nextY,nextX)
    }
  }

	nextEntityOldGrid := (*oldGrid)[nextY][nextX]
	nextEntityNextGrid := (*nextGrid)[nextY][nextX]


	if (nextEntityOldGrid == nil && nextEntityNextGrid == nil) || (nextEntityOldGrid != nil && nextEntityOldGrid.NextMove != ""){
		(*nextGrid)[nextY][nextX] = entity
	} else {
		if !entity.RemoveOnContact {
			(*nextGrid)[y][x] = entity
		}
		if entity.Damage > 0 {
			if nextEntityOldGrid != nil && !nextEntityOldGrid.Indestructible {
				oldGrid.DealDamage(nextEntityOldGrid, entity.Damage)
			} else if nextEntityNextGrid != nil && !nextEntityNextGrid.Indestructible {
				nextGrid.DealDamage(nextEntityNextGrid, entity.Damage)
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
	e.IsDead = false
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

func (g *Grid) DealDamage(entity *Entity, amount int) {
	if !entity.Indestructible {
		entity.Health -= amount

		if entity.Health <= 0 {
			entity.IsDead = true
      log.Debug(fmt.Sprintf("Killing entity %s ID=%s",entity.Type, entity.ID))
			_, y, x, err := g.FindEntity(entity.ID)

			if err != nil {
				log.Error(fmt.Sprintf("Tried to kill an entity that could not be found. %s", entity.ID))
			} else {
        log.Debug(fmt.Sprintf("Removing entity at y=%d, x=%d", y, x))
				(*g)[y][x] = nil
			}
		}
	}
}
