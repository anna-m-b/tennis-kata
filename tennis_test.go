package tennis

import (
	"fmt"

	"testing"
)

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

func TestHasPlayerWonWithFourPoints(t *testing.T) {
	g := NewGame()
	g.PlayerA = 4
	wantA := true
	gotA := g.HasPlayerWon("A")
	if !gotA {
		t.Errorf("incorrect return value from HasPlayerWon: wanted %v, got %v", wantA, gotA)
	}
	wantB := true
	gotB := g.HasPlayerWon("B")
	if gotB {
		t.Errorf("incorrect return value from HasPlayerWon: wanted %v, got %v", wantB, gotB)
	}
}

func TestHasPlayerWonByTwoClearPoints(t *testing.T) {
	g := NewGame()
	g.PlayerA = 4
	g.PlayerB = 3
	want := false
	got := g.HasPlayerWon("A")
	if got {
		t.Errorf("incorrect return value from HasPlayerWon: wanted %v, got %v", want, got)
	}
}
