package sorting

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

func quickSort(items []int, leftIndex, rightIndex int) []int {
	if len(items) <= 1 {
		return items
	}

	pivotIndex := partition(leftIndex, rightIndex, items)

	if leftIndex < pivotIndex-1 {
		quickSort(items, leftIndex, pivotIndex-1)
	}

	if pivotIndex < rightIndex {
		quickSort(items, pivotIndex, rightIndex)
	}

	return items
}

func partition(leftIndex, rightIndex int, items []int) int {
	l := leftIndex
	r := rightIndex
	pivot := items[(rightIndex+leftIndex)/2]

	for l <= r {
		for items[l] < pivot {
			l++
		}

		for items[r] > pivot {
			r--
		}

		if l <= r {
			temp := items[l]
			items[l] = items[r]
			items[r] = temp

			l++
			r--
		}
	}

	return l
}
