package app

import (
	"fmt"
	"github.com/enescakir/emoji"
	"github.com/google/uuid"
	"strings"
)

const (
	WALL              = "wall"
	PLAYER            = "player"
	FIREBALL          = "fireball"
	CASTLE            = "castle"
	NORTH             = "north"
	SOUTH             = "south"
	EAST              = "east"
	WEST              = "west"
	PLAYER_MAX_HEALTH = 5
)

var (
	pSkins       = [10]string{"p1", "p2", "p3", "p4", "p5", "p6", "p7", "p8", "p9", "p10"}
	currentPSkin = 0
)

type Entity struct {
	Skin            string
	Type            string
	ID              string
	NextMove        string
	Indestructible  bool
	RemoveOnContact bool
	Health          int
	Damage          int
	IsDead          bool
}

func CreateNextPlayer() *Entity {
	s := pSkins[currentPSkin]
	currentPSkin++
	if currentPSkin >= len(pSkins) {
		currentPSkin = 0
	}
	id := uuid.New().String()
	return &Entity{
		Type:           PLAYER,
		Skin:           s,
		ID:             id,
		Indestructible: false,
		Health:         PLAYER_MAX_HEALTH,
	}
}

func CreateCastle() *Entity {
	id := uuid.New().String()
	return &Entity{
		Type:           CASTLE,
		ID:             id,
		Indestructible: true,
	}
}

func CreateFireball(direction string) *Entity {
	id := uuid.New().String()
	return &Entity{
		Type:            FIREBALL,
		ID:              id,
		NextMove:        direction,
		Indestructible:  false,
		RemoveOnContact: true,
		Health:          1,
		Damage:          1,
	}
}

func (e *Entity) Render(sb *strings.Builder) {
	if e == nil {
		sb.WriteString("  ")
	} else {
		switch e.Type {
		case WALL:
			sb.WriteString("W")
		case PLAYER:
			switch e.Skin {
			case pSkins[0]:
				sb.WriteString(fmt.Sprintf("%v", emoji.Mage))
			case pSkins[1]:
				sb.WriteString(fmt.Sprintf("%v", emoji.Elf))
			case pSkins[2]:
				sb.WriteString(fmt.Sprintf("%v", emoji.Vampire))
			case pSkins[3]:
				sb.WriteString(fmt.Sprintf("%v", emoji.Fairy))
			case pSkins[4]:
				sb.WriteString(fmt.Sprintf("%v", emoji.Robot))
			case pSkins[5]:
				sb.WriteString(fmt.Sprintf("%v", emoji.Merperson))
			case pSkins[6]:
				sb.WriteString(fmt.Sprintf("%v", emoji.SantaClaus))
			case pSkins[7]:
				sb.WriteString(fmt.Sprintf("%v", emoji.Cat))
			case pSkins[8]:
				sb.WriteString(fmt.Sprintf("%v", emoji.Genie))
			case pSkins[9]:
				sb.WriteString(fmt.Sprintf("%v", emoji.Ogre))
			default:
				sb.WriteString("?")
			}
		case FIREBALL:
			sb.WriteString(fmt.Sprintf("%v", emoji.Fire))
		case CASTLE:
			sb.WriteString(fmt.Sprintf("%v", emoji.Castle))
		default:
			sb.WriteString("?")
		}
	}
}
