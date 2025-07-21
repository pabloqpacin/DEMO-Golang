package iteration

// https://pkg.go.dev/strings
import "strings"

func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func Count(s, substr string) int {
	return strings.Count(s, substr)
}

func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}
