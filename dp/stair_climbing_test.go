package dp

import (
	"testing"
)

var k = 4
var answer int

func TestRecursion(t *testing.T) {
	n := StairClimbingRecursion(k)
	t.Log("total kind with Recursion: ", n)
}

func BenchmarkRecursion(t *testing.B) {
	for i := 0; i < t.N; i++ {
		StairClimbingRecursion(k)
	}
}

func TestRecursionWithCache(t *testing.T) {
	n := StairClimbingRecursionWithCache(k)
	t.Log("total kind with RecursionWithCache: ", n)
}

func TestStairClimbingForwardPropagate(t *testing.T) {
	n := StairClimbingForwardPropagate(k)
	t.Log("total kind with StairClimbingForwardPropagate: ", n)
}
