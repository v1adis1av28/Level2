package tests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/v1adis1av28/level2/tasks/task13/internal/cut"
)

type ParseFieldsAns struct {
	Arr []int
	Err bool
}

func TestParseFieldsFlag(t *testing.T) {
	inputArr := []string{"1,3,5", "", "1-3", "1.3@", "4-2"}
	expected := []ParseFieldsAns{ParseFieldsAns{Arr: []int{1, 3, 5}, Err: false},
		ParseFieldsAns{Arr: []int{}, Err: true},
		ParseFieldsAns{Arr: []int{1, 2, 3}, Err: false},
		ParseFieldsAns{Arr: []int{}, Err: true}, ParseFieldsAns{Arr: []int{2, 3, 4}, Err: false}}

	for i, testInput := range inputArr {
		test, err := cut.ParseFieldsFlag(testInput)
		if err != nil && expected[i].Err {
			fmt.Errorf("Error on testin parse fields expected %v input %s", expected[i], inputArr[i])
		}
		if !reflect.DeepEqual(test, expected[i].Arr) {
			fmt.Errorf("Parse fields error expected %v input %s", expected[i], inputArr[i])
		}
	}
}
