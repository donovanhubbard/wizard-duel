package entities

import (
  "strings"
)

const (
  WALL = "wall"
  PLAYER = "player"
)

type Entity struct {
  ID string
  Type string
}

func newWall()*Entity{
  return &Entity{
    Type: WALL,
    ID: "none",
  }
}

func (e *Entity)Render(sb *strings.Builder){
  if e == nil {
    sb.WriteString(".")
  }else{
    switch e.Type{
    case WALL:
      sb.WriteString("W")
    default:
      sb.WriteString("?")
    }
  }
}
