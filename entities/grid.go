package entities

import (
  "math/rand"
)

const (
  WIDTH = 30
  HEIGHT = 20
)

type Grid [][] *Entity

func NewGrid() Grid{
  var g Grid
  g = make([][]*Entity,HEIGHT)
  for i := range g {
    g[i] = make([]*Entity,WIDTH)
  }

  return g
}

func (g *Grid) PlacePlayer(p *Entity){
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

