package sort

import (
	"fmt"
	"sort"
)

func BaseSort(strs []string) {
	sort.Strings(strs)
	fmt.Println(strs)
}

func ReverseSort(arr []string) {
	sort.Strings(arr)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println(arr)

}

func UniqueSort(arr []string) {
	sort.Strings(arr)

	mp := make(map[string]struct{})
	uniqueSlice := make([]string, 0, len(arr))

	for _, val := range arr {
		if _, ok := mp[val]; !ok {
			mp[val] = struct{}{}
			uniqueSlice = append(uniqueSlice, val)
		}
	}

	fmt.Println(uniqueSlice)
}
