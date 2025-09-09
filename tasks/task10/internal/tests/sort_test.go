package tests

import (
	"reflect"
	"testing"

	"github.com/v1adis1av28/level2/tasks/task10/internal/sort"
)

func TestCompareLines(t *testing.T) {
	tests := []struct {
		name     string
		a, b     sort.Line
		keyCol   int
		numeric  bool
		expected bool
	}{
		{
			name:     "lexicographic a < b",
			a:        sort.Line{Columns: []sort.Column{{Number: 1, Line: "a"}}, BaseStr: "a"},
			b:        sort.Line{Columns: []sort.Column{{Number: 1, Line: "b"}}, BaseStr: "b"},
			keyCol:   1,
			numeric:  false,
			expected: true,
		},
		{
			name:     "numeric 2 < 10",
			a:        sort.Line{Columns: []sort.Column{{Number: 1, Line: "2"}}, BaseStr: "x 2"},
			b:        sort.Line{Columns: []sort.Column{{Number: 1, Line: "10"}}, BaseStr: "y 10"},
			keyCol:   1,
			numeric:  true,
			expected: true,
		},
		{
			name:     "numeric 10 > 2 reverse logic",
			a:        sort.Line{Columns: []sort.Column{{Number: 1, Line: "10"}}, BaseStr: "y 10"},
			b:        sort.Line{Columns: []sort.Column{{Number: 1, Line: "2"}}, BaseStr: "x 2"},
			keyCol:   1,
			numeric:  true,
			expected: false,
		},
		{
			name:     "equality a = a lexigraphicly",
			a:        sort.Line{Columns: []sort.Column{{Number: 1, Line: "a"}}, BaseStr: "a"},
			b:        sort.Line{Columns: []sort.Column{{Number: 1, Line: "a"}}, BaseStr: "a"},
			keyCol:   1,
			numeric:  false,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sort.CompareLines(tt.a, tt.b, tt.keyCol, tt.numeric)
			if got != tt.expected {
				t.Errorf("compareLines() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestUniqueValueFlag(t *testing.T) {
	testInput := []sort.Line{
		{Columns: []sort.Column{{Number: 1, Line: "one"}}, BaseStr: "one"},
		{Columns: []sort.Column{{Number: 1, Line: "one"}}, BaseStr: "one"},
		{Columns: []sort.Column{{Number: 1, Line: "two"}}, BaseStr: "two"},
		{Columns: []sort.Column{{Number: 1, Line: "two"}}, BaseStr: "two"},
	}

	expectetOutput := []sort.Line{
		{Columns: []sort.Column{{Number: 1, Line: "one"}}, BaseStr: "one"},
		{Columns: []sort.Column{{Number: 1, Line: "two"}}, BaseStr: "two"},
	}

	uniqueLine := sort.Deduplicate(testInput, 1, true)
	if !reflect.DeepEqual(uniqueLine, expectetOutput) {
		t.Errorf("Unique value sort error input = %v expected = %v", uniqueLine, expectetOutput)
	}
}

func TestReverseOrder(t *testing.T) {
	testInput := []sort.Line{
		{Columns: []sort.Column{{Number: 1, Line: "a"}}, BaseStr: "a"},
		{Columns: []sort.Column{{Number: 1, Line: "b"}}, BaseStr: "b"},
		{Columns: []sort.Column{{Number: 1, Line: "c"}}, BaseStr: "c"},
	}

	expected := []string{"c", "b", "a"}
	test := sort.SortLines(testInput, 1, false, true, false)
	if !reflect.DeepEqual(test, expected) {
		t.Errorf("Wrong reverse order expected = %v get = %v", expected, test)
	}
}
