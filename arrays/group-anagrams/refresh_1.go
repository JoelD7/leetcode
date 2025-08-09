package group_anagrams

import (
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortStr(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func groupAnagrams(strs []string) [][]string {
	groupMap := make(map[string][]string)
	var str, sorted string

	for _, str = range strs {
		sorted = sortStr(str)
		group, ok := groupMap[sorted]
		if ok {
			groupMap[sorted] = append(group, str)
			continue
		}

		groupMap[sorted] = []string{str}
	}

	result := make([][]string, 0)
	for _, v := range groupMap {
		result = append(result, v)
	}

	return result
}
