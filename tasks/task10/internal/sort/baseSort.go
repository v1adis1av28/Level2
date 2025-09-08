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
