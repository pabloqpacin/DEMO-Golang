package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repite 5 veces por defecto si repeatTimes es 0", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("repite n veces si repeatTimes es positivo", func(t *testing.T) {
		repeated := Repeat("b", 3)
		expected := "bbb"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

// https://pkg.go.dev/testing#hdr-Benchmarks
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("c", 4)
	fmt.Println(repeated)
	// Output: cccc
}
