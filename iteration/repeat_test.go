package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("print 'aaaaa'", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(repeated, expected, t)
	})
	t.Run("print 'b' specifying N", func(t *testing.T) {
		repeated := Repeat("b", 3)
		expected := "bbb"
		assertCorrectMessage(repeated, expected, t)
	})
}

func assertCorrectMessage(repeated string, expected string, t *testing.T) {
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// go test -bench="."
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("c", 7)
	fmt.Println(repeated)
	//Output: ccccccc
}
