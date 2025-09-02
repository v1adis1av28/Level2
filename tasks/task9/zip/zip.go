package zip

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Написать функцию Go, осуществляющую примитивную распаковку строки, содержащей повторяющиеся символы/руны.
// Примеры работы функции:
// Вход: "a4bc2d5e"
// Выход: "aaaabccddddde"
// Вход: "abcd"
// Выход: "abcd" (нет цифр — ничего не меняется)
// Вход: "45"
// Выход: "" (некорректная строка, т.к. в строке только цифры — функция должна вернуть ошибку)
// Вход: ""
// Выход: "" (пустая строка -> пустая строка)
// Дополнительное задание
// Поддерживать escape-последовательности вида \:
// Вход: "qwe\4\5"
// Выход: "qwe45" (4 и 5 не трактуются как числа, т.к. экранированы)
// Вход: "qwe\45"
// Выход: "qwe44444" (\4 экранирует 4, поэтому распаковывается только 5)
// Требования к реализации
// Функция должна корректно обрабатывать ошибочные случаи (возвращать ошибку, например, через error), и проходить unit-тесты.
// Код должен быть статически анализируем (vet, golint).

var emptyStr = ""
var EmptyStrError = errors.New("String is empty")
var InvalidStrError = errors.New("Invalid type of string")

func UnzipString(str string) (string, error) {

	if isEmpty(str) {
		return emptyStr, EmptyStrError
	}

	if !isValid(str) {
		return emptyStr, InvalidStrError
	}

	var resultStr strings.Builder
	for i := 0; i < len(str); i++ {
		if (i + 1) < (len(str) - 1) {
			if unicode.IsDigit(rune(str[i+1])) && string(str[i]) != "\n" {
				//если число получаем кол-во
				count, err := strconv.Atoi(string(str[i+1]))
				if err != nil {
					fmt.Println(err)
				}
				resultStr.WriteString(strings.Repeat(string(str[i]), count))
			} else {
				if string(str[i]) == "\n" {
					continue
				} else {
					resultStr.WriteString(string(str[i]))
				}
			}
		} else { // кейс последнего эл-та
			//проверить если предпоследний эл-т это экранирование тогда стоп
			//Проверить если последнее это число и перед стоящий эл-т что
			if unicode.IsDigit(rune(str[i])) {
				count, err := strconv.Atoi(string(str[i]))
				if err != nil {
					fmt.Println(err)
				}
				resultStr.WriteString(strings.Repeat(string(str[i]), count))
			}
			if string(str[i]) != "\n" {
				resultStr.WriteString(string(str[i]))
			}
		}
	}
	return resultStr.String(), nil
}

func isEmpty(s string) bool {
	return len(s) < 1
}

func isValid(s string) bool {
	for ind, val := range s {
		if ind != 0 {
			if unicode.IsLetter(val) && string(s[ind-1]) != "\n" {
				return true
			}
		} else {
			if unicode.IsLetter(val) {
				return true
			}
		}
	}
	return false
}
