package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedArr := make([]int, 0, len(nums1)+len(nums2))

	if len(mergedArr) == 1 {
		return float64(mergedArr[0])
	}

	var i1, i2 int

	for i1 < len(nums1) && i2 < len(nums2) {
		if nums1[i1] < nums2[i2] {
			mergedArr = append(mergedArr, nums1[i1])
			i1++
		} else {
			mergedArr = append(mergedArr, nums2[i2])
			i2++
		}
	}

	if i1 < len(nums1) {
		mergedArr = append(mergedArr, nums1[i1:]...)
	} else {
		mergedArr = append(mergedArr, nums2[i2:]...)
	}

	mid := len(mergedArr) / 2
	if len(mergedArr)%2 == 0 {
		return float64(mergedArr[mid]+mergedArr[mid-1]) / 2
	}

	return float64(mergedArr[mid])
}
