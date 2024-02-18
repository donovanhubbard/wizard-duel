package app

import (
	"testing"
)

func TestMoveNoDirection(t *testing.T) {
	app := &App{grid: NewGrid()}
	player := CreateNextPlayer()
	(*app.grid)[1][1] = player
	app.MoveAll()
	targetEntity := (*app.grid)[1][1]
	if targetEntity == nil || player.ID != targetEntity.ID {
		t.Fatalf("Player moved when it shouldn't have")
	}
}

func TestMovePlayerNorth(t *testing.T) {
	app := &App{grid: NewGrid()}
	player := CreateNextPlayer()
	player.NextMove = NORTH
	(*app.grid)[1][1] = player
	app.MoveAll()
	targetEntity := (*app.grid)[0][1]
	if targetEntity == nil || player.ID != targetEntity.ID {
		t.Fatalf("Player did not move where it should have")
	}
}

func TestMovePlayerSouth(t *testing.T) {
	app := &App{grid: NewGrid()}
	player := CreateNextPlayer()
	player.NextMove = SOUTH
	(*app.grid)[1][1] = player
	app.MoveAll()
	targetEntity := (*app.grid)[2][1]
	if targetEntity == nil || player.ID != targetEntity.ID {
		t.Fatalf("Player did not move where it should have")
	}
}

func TestMovePlayerEast(t *testing.T) {
	app := &App{grid: NewGrid()}
	player := CreateNextPlayer()
	player.NextMove = EAST
	(*app.grid)[1][1] = player
	app.MoveAll()
	targetEntity := (*app.grid)[1][2]
	if targetEntity == nil || player.ID != targetEntity.ID {
		t.Fatalf("Player did not move where it should have")
	}
}

func TestMovePlayerWest(t *testing.T) {
	app := &App{grid: NewGrid()}
	player := CreateNextPlayer()
	player.NextMove = WEST
	(*app.grid)[1][1] = player
	app.MoveAll()
	targetEntity := (*app.grid)[1][0]
	if targetEntity == nil || player.ID != targetEntity.ID {
		t.Fatalf("Player did not move where it should have")
	}
}

func TestMovePlayerNorthIntoBoundary(t *testing.T) {
	app := &App{grid: NewGrid()}
	player := CreateNextPlayer()
	player.NextMove = NORTH
	(*app.grid)[0][1] = player
	app.MoveAll()
	targetEntity := (*app.grid)[0][1]
	if targetEntity == nil || player.ID != targetEntity.ID {
		t.Fatalf("Player did not move where it should have")
	}
}

func TestMovePlayerSouthIntoBoundary(t *testing.T) {
	app := &App{grid: NewGrid()}
	player := CreateNextPlayer()
	player.NextMove = SOUTH
	(*app.grid)[len(*app.grid)-1][1] = player
	app.MoveAll()
	targetEntity := (*app.grid)[len(*app.grid)-1][1]
	if targetEntity == nil || player.ID != targetEntity.ID {
		t.Fatalf("Player did not move where it should have")
	}
}
func TestMovePlayerEastIntoBoundary(t *testing.T) {
	app := &App{grid: NewGrid()}
	player := CreateNextPlayer()
	player.NextMove = EAST
	(*app.grid)[1][len((*app.grid)[1])-1] = player
	app.MoveAll()
	targetEntity := (*app.grid)[1][len((*app.grid)[1])-1]
	if targetEntity == nil || player.ID != targetEntity.ID {
		t.Fatalf("Player did not move where it should have")
	}
}

func TestMovePlayerWestIntoBoundary(t *testing.T) {
	app := &App{grid: NewGrid()}
	player := CreateNextPlayer()
	player.NextMove = WEST
	(*app.grid)[1][0] = player
	app.MoveAll()
	targetEntity := (*app.grid)[1][0]
	if targetEntity == nil || player.ID != targetEntity.ID {
		t.Fatalf("Player did not move where it should have")
	}
}

func TestMovePlayerNorthIntoCastle(t *testing.T) {
	app := &App{grid: NewGrid()}
	player := CreateNextPlayer()
	player.NextMove = NORTH
	(*app.grid)[1][1] = player
	castle := CreateCastle()
	(*app.grid)[0][1] = castle
	app.MoveAll()
	targetEntity := (*app.grid)[1][1]
	if targetEntity == nil || player.ID != targetEntity.ID {
		t.Fatalf("Player moved when it should not have")
	}
	targetEntity = (*app.grid)[0][1]
	if targetEntity == nil || castle.ID != targetEntity.ID {
		t.Fatalf("Castle moved when it should not have")
	}
}

func TestMovePlayerSouthIntoCastle(t *testing.T) {
	app := &App{grid: NewGrid()}
	player := CreateNextPlayer()
	player.NextMove = SOUTH
	(*app.grid)[1][1] = player
	castle := CreateCastle()
	(*app.grid)[2][1] = castle
	app.MoveAll()
	targetEntity := (*app.grid)[1][1]
	if targetEntity == nil || player.ID != targetEntity.ID {
		t.Fatalf("Player moved when it should not have")
	}
	targetEntity = (*app.grid)[2][1]
	if targetEntity == nil || castle.ID != targetEntity.ID {
		t.Fatalf("Castle moved when it should not have")
	}
}

func TestMovePlayerEastIntoCastle(t *testing.T) {
	app := &App{grid: NewGrid()}
	player := CreateNextPlayer()
	player.NextMove = EAST
	(*app.grid)[1][1] = player
	castle := CreateCastle()
	(*app.grid)[1][2] = castle
	app.MoveAll()
	targetEntity := (*app.grid)[1][1]
	if targetEntity == nil || player.ID != targetEntity.ID {
		t.Fatalf("Player moved when it should not have")
	}
	targetEntity = (*app.grid)[1][2]
	if targetEntity == nil || castle.ID != targetEntity.ID {
		t.Fatalf("Castle moved when it should not have")
	}
}

func TestMovePlayerWestIntoCastle(t *testing.T) {
	app := &App{grid: NewGrid()}
	player := CreateNextPlayer()
	player.NextMove = WEST
	(*app.grid)[1][1] = player
	castle := CreateCastle()
	(*app.grid)[1][0] = castle
	app.MoveAll()
	targetEntity := (*app.grid)[1][1]
	if targetEntity == nil || player.ID != targetEntity.ID {
		t.Fatalf("Player moved when it should not have")
	}
	targetEntity = (*app.grid)[1][0]
	if targetEntity == nil || castle.ID != targetEntity.ID {
		t.Fatalf("Castle moved when it should not have")
	}
}

func TestMoveFireballNorthIntoBoundary(t *testing.T) {
	app := &App{grid: NewGrid()}
	fireball := CreateFireball(NORTH)
	(*app.grid)[0][1] = fireball
	app.MoveAll()
	entity, _, _, err := app.grid.FindEntity(fireball.ID)
	if entity != nil || err == nil {
		t.Fatalf("Fireball did not disapear")
	}
}

func TestMoveFireballSouthIntoBoundary(t *testing.T) {
	app := &App{grid: NewGrid()}
	fireball := CreateFireball(SOUTH)
	(*app.grid)[len(*app.grid)-1][1] = fireball
	app.MoveAll()
	entity, _, _, err := app.grid.FindEntity(fireball.ID)
	if entity != nil || err == nil {
		t.Fatalf("Fireball did not disapear")
	}
}

func TestMoveFireballEastIntoBoundary(t *testing.T) {
	app := &App{grid: NewGrid()}
	fireball := CreateFireball(EAST)
	(*app.grid)[1][len((*app.grid)[1])-1] = fireball
	app.MoveAll()
	entity, _, _, err := app.grid.FindEntity(fireball.ID)
	if entity != nil || err == nil {
		t.Fatalf("Fireball did not disapear")
	}
}

func TestMoveFireballWestIntoBoundary(t *testing.T) {
	app := &App{grid: NewGrid()}
	fireball := CreateFireball(WEST)
	(*app.grid)[1][0] = fireball
	app.MoveAll()
	entity, _, _, err := app.grid.FindEntity(fireball.ID)
	if entity != nil || err == nil {
		t.Fatalf("Fireball did not disapear")
	}
}

func TestMoveFireballNorthIntoCastle(t *testing.T) {
	app := &App{grid: NewGrid()}
	fireball := CreateFireball(NORTH)
	(*app.grid)[1][1] = fireball
	castle := CreateCastle()
	(*app.grid)[0][1] = castle
	app.MoveAll()
	entity, _, _, err := app.grid.FindEntity(fireball.ID)
	if entity != nil || err == nil {
		t.Fatalf("Fireball did not disapear")
	}
}

func TestMoveFireballSouthIntoCastle(t *testing.T) {
	app := &App{grid: NewGrid()}
	fireball := CreateFireball(SOUTH)
	(*app.grid)[1][1] = fireball
	castle := CreateCastle()
	(*app.grid)[2][1] = castle
	app.MoveAll()
	entity, _, _, err := app.grid.FindEntity(fireball.ID)
	if entity != nil || err == nil {
		t.Fatalf("Fireball did not disapear")
	}
}

func TestMoveFireballEastIntoCastle(t *testing.T) {
	app := &App{grid: NewGrid()}
	fireball := CreateFireball(EAST)
	(*app.grid)[1][1] = fireball
	castle := CreateCastle()
	(*app.grid)[1][2] = castle
	app.MoveAll()
	entity, _, _, err := app.grid.FindEntity(fireball.ID)
	if entity != nil || err == nil {
		t.Fatalf("Fireball did not disapear")
	}
}

func TestMoveFireballWestIntoCastle(t *testing.T) {
	app := &App{grid: NewGrid()}
	fireball := CreateFireball(WEST)
	(*app.grid)[1][1] = fireball
	castle := CreateCastle()
	(*app.grid)[1][0] = castle
	app.MoveAll()
	entity, _, _, err := app.grid.FindEntity(fireball.ID)
	if entity != nil || err == nil {
		t.Fatalf("Fireball did not disapear")
	}
}
