package tennis

import (
	"fmt"
	
	"testing"
)

// game
// - playerA: "love", "15", "30", "40", "deuce", "advantage"
// - playerB
// play() -randomly move one player one point along
// completeGame() - returns winner

func TestNewGame(t *testing.T) {
	g := NewGame()
	if score[g.PlayerA] != "love" {
		t.Errorf("PlayerA is a cheat: wanted love, got %v", score[g.PlayerA])
	}
	if score[g.PlayerB] != "love" {
		t.Errorf("PlayerB is a cheat: wanted love, got %v", score[g.PlayerB])
	}

}

func TestPlayChangesTheScore(t *testing.T) {
	g := NewGame()
	g.Play()
	if score[g.PlayerA] == "love" && score[g.PlayerB] == "love" {
		t.Errorf("g.Play did not change the score: PlayerA: %s, PlayerB: %s", score[g.PlayerA], score[g.PlayerB])
	}

}

func TestPlayIncreaesOnePlayersScore(t *testing.T) {
	g := NewGame()
	g.Play()
	if score[g.PlayerA] != "15" && score[g.PlayerB] != "15" {
		t.Errorf("g.Play did not increase the score: PlayerA: %s, PlayerB: %s", score[g.PlayerA], score[g.PlayerB])
	}
	if score[g.PlayerA] == "15" && score[g.PlayerB] == "15" {
		t.Errorf("g.Play increased both players' score: PlayerA: %s, PlayerB: %s", score[g.PlayerA], score[g.PlayerB])
	}

	fmt.Println(g.PlayerA, g.PlayerB)

}

func TestPlayChangesScoresToDeuce(t *testing.T) {
	g := NewGame()
	g.PlayerA = 3

}
