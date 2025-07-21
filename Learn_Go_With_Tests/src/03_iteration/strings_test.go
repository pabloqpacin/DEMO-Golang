package iteration

import (
	"strings"
	"testing"
)

func TestContains(t *testing.T) {
	contains := strings.Contains("test", "es")
	if !contains {
		t.Errorf("expected %t but got %t", true, contains)
	}
}

func TestCount(t *testing.T) {
	count := strings.Count("test", "t")
	if count != 2 {
		t.Errorf("expected %d but got %d", 2, count)
	}
}

func TestHasPrefix(t *testing.T) {
	hasPrefix := strings.HasPrefix("test", "te")
	if !hasPrefix {
		t.Errorf("expected %t but got %t", true, hasPrefix)
	}
}
