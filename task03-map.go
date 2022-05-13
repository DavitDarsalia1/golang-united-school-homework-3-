package homework

import (
	"fmt"
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, 0)
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for i, _ := range keys {
		fmt.Println(input[i])
		result = append(result, input[i+1])
	}

	return

}
