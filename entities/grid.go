package entities

const (
  WIDTH = 40
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

