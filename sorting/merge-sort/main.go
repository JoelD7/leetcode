package main

func mergeSort(items []int) []int {
	if len(items) <= 1 {
		return items
	}

	mid := len(items) / 2
	left := items[0:mid]
	right := items[mid:]

	left = mergeSort(left)
	right = mergeSort(right)
	items = merge(left, right)

	return items
}

func merge(left, right []int) []int {
	i := 0
	j := 0
	items := make([]int, 0, len(left)+len(right))

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			items = append(items, left[i])
			i++
		} else {
			items = append(items, right[j])
			j++
		}
	}

	for ; i < len(left); i++ {
		items = append(items, left[i])
	}

	for ; j < len(right); j++ {
		items = append(items, right[j])
	}

	return items
}
