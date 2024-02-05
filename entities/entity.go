package entities

import (
  "github.com/enescakir/emoji"
  "strings"
  "fmt"
)

const (
  WALL = "wall"
  PLAYER = "player"
)

var (
  pSkins = [10]string {"p1","p2","p3","p4","p5","p6","p7","p8","p9","p10"}
  currentPSkin = 0
)

type Entity struct {
  Skin string
  Type string
}

func CreateNextPlayer() Entity {
  s := pSkins[currentPSkin]
  currentPSkin++
  if currentPSkin >= len(pSkins){
    currentPSkin = 0
  }

  return Entity{
    Type: PLAYER,
    Skin: s,
  }
}

func (e *Entity)Render(sb *strings.Builder){
  if e == nil {
    sb.WriteString("\\/")
  }else{
    switch e.Type{
    case WALL:
      sb.WriteString("W")
    case PLAYER:
     switch e.Skin {
     case pSkins[0]:
       sb.WriteString(fmt.Sprintf("%v",emoji.Mage))
     case pSkins[1]:
       sb.WriteString(fmt.Sprintf("%v",emoji.Elf))
     case pSkins[2]:
       sb.WriteString(fmt.Sprintf("%v",emoji.Vampire))
     case pSkins[3]:
       sb.WriteString(fmt.Sprintf("%v",emoji.Fairy))
     case pSkins[4]:
       sb.WriteString(fmt.Sprintf("%v",emoji.Robot))
     case pSkins[5]:
       sb.WriteString(fmt.Sprintf("%v",emoji.Merperson))
     case pSkins[6]:
       sb.WriteString(fmt.Sprintf("%v",emoji.SantaClaus))
     case pSkins[7]:
       sb.WriteString(fmt.Sprintf("%v",emoji.Cat))
     case pSkins[8]:
       sb.WriteString(fmt.Sprintf("%v",emoji.Genie))
     case pSkins[9]:
       sb.WriteString(fmt.Sprintf("%v",emoji.Ogre))
      default:
        sb.WriteString("?")
      }
    default:
      sb.WriteString("?")
    }
  }
}
