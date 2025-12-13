package top_k_frequent_elements

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	freqByValue := map[int]int{}

	//Build frequency map
	for _, num := range nums {
		cur, ok := freqByValue[num]
		if ok {
			freqByValue[num] = cur + 1
			continue
		}

		freqByValue[num] = 1
	}

	type customMap struct {
		value     int
		frequency int
	}

	//Sort frequency map by values
	frequencies := make([]customMap, 0, len(freqByValue))

	for val, freq := range freqByValue {
		frequencies = append(frequencies, customMap{
			value:     val,
			frequency: freq,
		})
	}

	sort.Slice(frequencies, func(i int, j int) bool {
		return frequencies[j].frequency < frequencies[i].frequency
	})

	result := make([]int, 0, k)

	for i := 0; i < k; i++ {
		result = append(result, frequencies[i].value)
	}

	return result
}
