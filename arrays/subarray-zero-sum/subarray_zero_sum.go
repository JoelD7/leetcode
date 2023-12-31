package subarray_zero_sum

import "fmt"

/*
 Given an array consisting of integers (possibly negative integers). Check if there exists a non-empty subarray such
 that the sum of elements in it is 0.
*/

func findSubarrayZeroSum(array []int) bool {
	for _, val := range array {
		if val == 0 {
			return true
		}
	}

	prefixSum := getPrefixSum(array)
	indicesByVal := make(map[int]int)

	for i, val := range prefixSum {
		_, ok := indicesByVal[val]
		if ok {
			subArray := array[indicesByVal[val]+1 : i+1]
			fmt.Printf("Subarray: %v | Between %d and %d\n", subArray, indicesByVal[val]+1, i)
			return true
		}

		indicesByVal[val] = i
	}

	return false
}

func getPrefixSum(input []int) []int {
	output := make([]int, 0)

	prev := input[0]
	output = append(output, prev)

	for i := 1; i < len(input); i++ {
		prev += input[i]
		output = append(output, prev)
	}

	return output
}

func rangePrefixSum(prefixSum []int, l, r int) int {
	if l == 0 {
		return prefixSum[r]
	}

	return prefixSum[r] - prefixSum[l-1]
}
