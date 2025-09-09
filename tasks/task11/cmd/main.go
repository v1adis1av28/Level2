package main

import (
	"fmt"
	"sort"
	"strings"
)

//task
//Напишите функцию, которая находит все множества анаграмм по заданному словарю.

// Требования
// На вход подается срез строк (слов на русском языке в Unicode).
// На выходе: map-множество -> список, где ключом является первое встреченное слово множества,
// а значением — срез из всех слов, принадлежащих этому множеству анаграмм, отсортированных по возрастанию.
// Множества из одного слова не должны выводиться (т.е. если нет анаграмм, слово игнорируется).
// Все слова нужно привести к нижнему регистру.
// Пример:

// Вход: ["пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"]
// Результат (ключи в примере могут быть в другом порядке):
// – "пятак": ["пятак", "пятка", "тяпка"]
// – "листок": ["листок", "слиток", "столик"]
// Слово «стол» отсутствует в результатах, так как не имеет анаграмм.
// Для решения задачи потребуется умение работать со строками, сортировать
// и использовать структуры данных (map).

// Оценим эффективность: решение должно работать за линейно-логарифмическое время относительно
// количества слов (допустимо n * m log m, где m — средняя длина слова для сортировки букв).

func main() {
	input := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	resMap := findAnagrams(input) // срез отсортированных по возрастанию
	//if len(resMap[str]) == 0 // then skip
	//all adding word should be down cast to lower case

	fmt.Println(resMap)
}

func findAnagrams(arr []string) map[string][]string {
	res := make(map[string][]string)
	mp := make(map[string][]string)
	for _, val := range arr {
		lower := strings.ToLower(val)
		sortKey := sortKey(lower)
		mp[sortKey] = append(mp[sortKey], lower)
	}

	for _, val := range mp {
		if len(val) <= 1 {
			continue
		}
		sort.Strings(val)
		res[val[0]] = val
	}
	return res
}

func sortKey(str string) string {
	runes := []rune(strings.ToLower(str))
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	srted := string(runes)
	return srted
}
