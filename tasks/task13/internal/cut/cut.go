package cut

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Option struct {
	Fields    string // example 1,3-5 <-- string for parsing
	Delimetr  string // default value "\t"
	Separated bool   //if true show only string that contains separator

	Collumns []int
}

type Line struct {
	FullString  string //полная строка которую на каждой итерации readlIne мы проверяем и разбиваем на массив подстрок(колонки)
	CollumnsArr []string
	IsSeparated bool
}

func ParseOption() (*Option, error) {
	//todo
	//считать строку полей
	//обработать строку чтобы выделить что и как с ней надо сделать
	option := &Option{}
	flag.StringVar(&option.Fields, "f", "", "show the fields")
	flag.StringVar(&option.Delimetr, "d", "\t", "split the string by delimetr")
	flag.BoolVar(&option.Separated, "s", false, "flag that show only strings that contains delimetr")

	flag.Parse()
	//Определить разделитель
	//пройтись по строкам найдя разделители и добавить их в срез
	var err error
	option.Collumns, err = ParseFieldsFlag(option.Fields)
	if err != nil {
		return nil, err
	}
	//Пройти по срезу строк и в результирующий добавить только те что имеют true on separated
	return option, nil
}
func ParseFieldsFlag(spec string) ([]int, error) {
	if spec == "" {
		return nil, fmt.Errorf("fields specification is empty")
	}

	parts := strings.Split(spec, ",")
	var fields []int

	for _, part := range parts {
		if strings.Contains(part, "-") {
			bounds := strings.Split(part, "-")
			if len(bounds) != 2 {
				return nil, fmt.Errorf("invalid range: %s", part)
			}
			start, err := strconv.Atoi(bounds[0])
			if err != nil || start < 1 {
				return nil, fmt.Errorf("invalid start of range: %s", bounds[0])
			}
			end, err := strconv.Atoi(bounds[1])
			if err != nil || end < start {
				return nil, fmt.Errorf("invalid end of range: %s", bounds[1])
			}
			for i := start; i <= end; i++ {
				fields = append(fields, i)
			}
		} else {
			num, err := strconv.Atoi(part)
			if err != nil || num < 1 {
				return nil, fmt.Errorf("invalid field number: %s", part)
			}
			fields = append(fields, num)
		}
	}

	fields = uniqueSorted(fields)

	return fields, nil
}

func uniqueSorted(slice []int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	sort.Ints(list)
	return list
}

func Start(arr []Line, op *Option) {
	for _, line := range arr {
		if op.Separated && !line.IsSeparated {
			continue
		}

		var selected []string
		for _, index := range op.Collumns {
			if index <= len(line.CollumnsArr) {
				selected = append(selected, line.CollumnsArr[index-1])
			}
		}

		if len(selected) > 0 {
			fmt.Println(strings.Join(selected, op.Delimetr))
		}
	}
}
