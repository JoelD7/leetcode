package group_anagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	anagramGroup := map[string][]string{}
	var key string

	for _, str := range strs {
		sRunes := []rune(str)

		sort.Slice(sRunes, func(i, j int) bool {
			return sRunes[i] < sRunes[j]
		})

		key = string(sRunes)

		group, ok := anagramGroup[key]
		if ok {
			group = append(group, str)
			anagramGroup[key] = group
			continue
		}

		anagramGroup[key] = []string{str}
	}

	result := make([][]string, 0)

	for _, group := range anagramGroup {
		result = append(result, group)
	}

	return result
}
