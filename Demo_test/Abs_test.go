package Demo_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestAbs(t *testing.T) {
	got := math.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}
func TestSum(t *testing.T) {
	x := 1
	y := 1
	var got int
	got = x + y
	caseName := t.Name()
	caseName = "Test Sum 1 + 1"
	fmt.Println(caseName)
	if got != 2 {
		t.Errorf("x + y = %d; want 2", got)
	}
}

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}
