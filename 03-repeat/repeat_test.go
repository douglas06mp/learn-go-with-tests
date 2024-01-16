package iteration

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeat := Repeat("a", 5)
	expected := strings.Repeat("a", 5)

	if repeat != expected {
		t.Errorf("expected %q but got %q", expected, repeat)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
