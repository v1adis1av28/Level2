package sort

import (
	"fmt"
	"sort"
	"strings"
)

func BaseSort(strs []string) {
	for i, str := range strs {
		strs[i] = strings.ToLower(str)
	}
	sort.Strings(strs)
	fmt.Println(strs)
}
