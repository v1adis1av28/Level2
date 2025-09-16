package tests

import (
	"testing"

	"github.com/v1adis1av28/level2/tasks/task15/internal/tokenizer"
)

// echo "hello world"
// echo 'single quoted'
// echo hello\ world
// echo \"quoted\" -> "quoted"
// echo \| -> |
func TestTokenizer(t *testing.T) {
	testCases := []struct {
		input string
		want  []string
		err   string
	}{
		{`echo "hello world"`, []string{"echo", "hello world"}, ""},
		{`echo hello\ world`, []string{"echo", "hello world"}, ""},
		{`echo \"quoted\"`, []string{"echo", `"quoted"`}, ""},
		{`echo \|`, []string{"echo", "|"}, ""},
		{`echo 'single quoted'`, []string{"echo", "single quoted"}, ""},
		{`echo 'test\''`, []string{"echo", "test'"}, ""},
		{`echo "unclosed`, nil, "unclosed quote"},
		{``, []string{}, "Empty string"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := tokenizer.Tokenize(tc.input)
			if tc.err != "" {
				if err == nil {
					t.Errorf("Expected error %q, got nil", tc.err)
				} else if err.Error() != tc.err {
					t.Errorf("Expected error %q, got %q", tc.err, err.Error())
				}
				return
			}
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if !equalSlices(got, tc.want) {
				t.Errorf("Tokenize(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
