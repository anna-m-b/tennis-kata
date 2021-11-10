package tennis

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	PlayerA int
	PlayerB int
}

var score = [6]string{"love", "15", "30", "40", "advantage", "win"}

func NewGame() *Game {
	return &Game{
		PlayerA: 0,
		PlayerB: 0,
	}
}

func (g *Game) Play() {
	rand.Seed(time.Now().UnixNano())
	scorer := rand.Intn(2)
	fmt.Print(scorer)
	if scorer == 0 {
		g.PlayerA = g.PlayerA + 1
		g.HasPlayerWon("A")
	} else {
		g.PlayerB = g.PlayerB + 1
		g.HasPlayerWon("B")
	}
}

func (g *Game) HasPlayerWon(p string) bool {
	var result bool

	if p == "A" {
		result = (g.PlayerA >= 4) && (g.PlayerA-g.PlayerB >= 2)
	} else if p == "B" {
		result = (g.PlayerB >= 4) && (g.PlayerB-g.PlayerA >= 2)
	}

	return result
}

// log score

// should player score reduce

// if they had 2 and other has >=2
// if they had advantage

// isItDeuce
// the other player has 3 and now so does this one
