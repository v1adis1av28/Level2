package handlertests

import (
	"strings"
	"testing"

	"github.com/v1adis1av28/level2/tasks/task15/internal/handler"
)

func TestCommandHandler(t *testing.T) {
	testCases := []struct {
		input   []string
		wantErr bool
		err     string
	}{
		{[]string{"cd", "pwd"}, false, ""},
		{[]string{"CD", "pwd"}, true, "unknown command"},
		{[]string{"dsa", "pwd"}, true, "unknown command"},
		{[]string{"", "pwd"}, true, "unknown command"},
		{[]string{"echo", "pwd"}, false, ""},
		{[]string{"kill", "pwd"}, false, ""},
		{[]string{"ps", "pwd"}, false, ""},
	}
	//todo add cases on different errors
	for _, val := range testCases {
		err := handler.HandleSingleCommand(val.input)
		if err != nil {
			if val.wantErr != true && err != nil {
				t.Errorf("throw error on case that dont suposed to throw case: %v", val.input)
			} else {
				if strings.Compare(val.err, err.Error()) != 0 {
					t.Errorf("Case %v throw wrong error expected : %v", val.input, val.err)
				}
			}
		}

	}
}
