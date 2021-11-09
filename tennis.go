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

var score = [6]string{"love", "15", "30", "40", "deuce", "advantage"}

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
	} else {
		g.PlayerB = g.PlayerB + 1
	}

}
