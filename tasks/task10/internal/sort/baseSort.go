package sort

import (
	"sort"
	"strconv"
)

type Column struct {
	Number int
	Line   string
}

type Line struct {
	Columns []Column
	BaseStr string
}

func GetColumnValue(line Line, colIndex int) string {
	if colIndex <= 0 || colIndex > len(line.Columns) {
		return ""
	}
	return line.Columns[colIndex-1].Line
}

func CompareLines(a, b Line, keyCol int, numeric bool) bool {
	valA := GetColumnValue(a, keyCol)
	valB := GetColumnValue(b, keyCol)

	if numeric {
		numA, errA := strconv.ParseFloat(valA, 64)
		numB, errB := strconv.ParseFloat(valB, 64)

		if errA == nil && errB == nil {
			return numA < numB
		}
	}
	return valA < valB
}

func Deduplicate(lines []Line, keyCol int, numeric bool) []Line {
	if len(lines) == 0 {
		return lines
	}

	result := []Line{lines[0]}

	for i := 1; i < len(lines); i++ {
		last := result[len(result)-1]
		curr := lines[i]

		if numeric {
			valLast := GetColumnValue(last, keyCol)
			valCurr := GetColumnValue(curr, keyCol)

			numLast, errLast := strconv.ParseFloat(valLast, 64)
			numCurr, errCurr := strconv.ParseFloat(valCurr, 64)

			if errLast == nil && errCurr == nil {
				if numLast == numCurr {
					continue
				}
			} else {
				if valLast == valCurr {
					continue
				}
			}
		} else {
			if GetColumnValue(last, keyCol) == GetColumnValue(curr, keyCol) {
				continue
			}
		}

		result = append(result, curr)
	}

	return result
}

func SortLines(lines []Line, keyColumn int, numeric, reverse, unique bool) []string {
	sorted := make([]Line, len(lines))
	copy(sorted, lines)

	sort.Slice(sorted, func(i, j int) bool {
		shouldIgoBeforeJ := CompareLines(sorted[i], sorted[j], keyColumn, numeric)
		if reverse {
			return !shouldIgoBeforeJ
		}
		return shouldIgoBeforeJ
	})

	if unique {
		sorted = Deduplicate(sorted, keyColumn, numeric)
	}

	result := make([]string, len(sorted))
	for i, line := range sorted {
		result[i] = line.BaseStr
	}

	return result
}
