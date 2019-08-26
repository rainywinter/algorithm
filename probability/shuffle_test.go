package probability

import (
	"testing"
)

func TestShuffle(t *testing.T) {
	n := 10
	card := make([]int, n, n)
	for i := 0; i < n; i++ {
		card[i] = i + 1
	}
	for i := 0; i < 4; i++ {
		Shuffle(card)
		t.Log(card)
	}
}
