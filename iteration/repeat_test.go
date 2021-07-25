package iteration_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/11me/go-tests/iteration"
)

func TestRepeat(t *testing.T) {
	repeated := iteration.Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q, but got %q\n", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("a", 5)
	}
}

// Native repeat function almost 2.5 faster
func BenchmarkNativeRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat("a", 5)
	}
}

func ExampleRepeat() {
	// Repeat function takes two arguments
	// the first argument is a charachter to repeat
	// and the seocnd one is how many times it should be repeated.

	fmt.Println(iteration.Repeat("b", 10))

	// Output: bbbbbbbbbb
}
