package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {

	for k := range input {
		result = append(result, input[k])
	}

	sort.Strings(result)

	return
}
