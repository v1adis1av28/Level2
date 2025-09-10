package cut

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Option struct {
	Fields    string // example 1,3-5 <-- string for parsing
	Delimetr  string // default value "\t"
	Separated bool   //if true show only string that contains separator
}

type Line struct {
	Text        string
	IsSeparated bool
}

func ParseOption() (*Option, error) {
	//todo
	//считать строку полей
	//обработать строку чтобы выделить что и как с ней надо сделать
	option := &Option{}
	flag.StringVar(&option.Fields, "f", "", "show the fields")
	flag.StringVar(&option.Delimetr, "d", "\\t", "split the string by delimetr")
	flag.BoolVar(&option.Separated, "s", false, "flag that show only strings that contains delimetr")

	flag.Parse()
	//Определить разделитель
	//пройтись по строкам найдя разделители и добавить их в срез

	//Пройти по срезу строк и в результирующий добавить только те что имеют true on separated
	return option, nil
}

func ParseFieldsFlag(str string) ([]int, error) { // возвращаем массив из индексов строк которые надо вывести
	indexes := make([]int, 0)
	unsuitableCharacters := "./\\!@#$%^&*()_"
	//var prev rune
	var isWaitingRange bool
	numberStr := ""
	for _, val := range str {
		if strings.Contains(unsuitableCharacters, string(val)) {
			return nil, fmt.Errorf("String contain unsuitable character")
		}
		if strings.Compare(string(val), ",") == 0 {
			number, _ := strconv.Atoi(numberStr)
			if isWaitingRange {
				for i := indexes[len(indexes)-1] + 1; i < number; i++ {
					indexes = append(indexes, i)
				}
				//	prev = val
				numberStr = ""
			}
			indexes = append(indexes, number)
			//prev = val
			numberStr = ""
		}
		if strings.Compare(string(val), "-") == 0 {
			isWaitingRange = true
			number, _ := strconv.Atoi(numberStr)
			indexes = append(indexes, number)
			numberStr = ""
		}

		if unicode.IsDigit(val) {
			numberStr += string(val)
		}

	}
	if numberStr != "" {
		number, _ := strconv.Atoi(numberStr)
		indexes = append(indexes, number)
	}
	fmt.Println(len(indexes))
	return indexes, nil
}

func Start(arr []Line, op *Option) {

}
