package tests

import (
	"strings"
	"testing"

	"github.com/v1adis1av28/Level2/tasks/task9/internal/zip"
)

type TestStruct struct {
	InputStr    string
	ExpectedStr string
	IsCorrect   bool
}

// В данном тесте ожидаем что ф-я будет возрвращать ошибку + nil
func TestEmptyString(t *testing.T) {
	emptyStr := ""
	exptected := zip.EmptyStrError
	_, err := zip.UnzipString(emptyStr)
	if exptected != err {
		t.Errorf("Call on empty string exptect emptyStringError")
	}
}

func TestIvalidString(t *testing.T) {
	expectedError := zip.InvalidStrError
	invalidStrings := []string{"1234", "\\\\\\\\", "     ", "....,.,."}
	for _, val := range invalidStrings {
		_, err := zip.UnzipString(val)
		if err != expectedError {
			t.Errorf("Expected invalid string error on string %s", val)
		}
	}
}

func TestUnpizpStr(t *testing.T) {
	testCases := []TestStruct{{"a4bc2d5e", "aaaabccddddde", true},
		{"abcd", "abcd", true},
		{"", "", true},
		{"qwe\\4\\5", "qwe45", true},
		{"qwe\\45", "qwe44444", true},
		{"45", "", false},
		{"\\\\", "", false}}
	for _, val := range testCases {
		str, err := zip.UnzipString(val.InputStr)
		if err == nil && !val.IsCorrect {
			t.Errorf("Exptected error on string %s expected error result:%s get: %s", val.InputStr, val.ExpectedStr, str)
		}
		compare := strings.Compare(val.ExpectedStr, str)
		if compare != 0 {
			t.Errorf("Wrong result on compare string %s expect %s get: %s", val.InputStr, val.ExpectedStr, str)
		}
	}
}
