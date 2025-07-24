package longest_harmonious_subsequence

import (
	"sort"
)

func findLHS(nums []int) int {
	sort.Ints(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}
	updateLargest := func(largest, count int) int {
		//If count == 0, this means all items are equal
		if count > 0 {
			return max(largest, count+1)
		}

		return largest
	}

	var i, newI, diff, count, largest int
	j := i + 1

	for j < len(nums) {
		diff = nums[j] - nums[i]
		if diff == 0 {
			j++
			//I don't increase `count` here because all the items of `nums` might be equal
			continue
		}

		if (diff == 1 || diff == -1) && newI == 0 {
			newI = j
			count = j - i //I'm only increasing `count` when I know there are different items((diff == 1 || diff==-1))
			j++
			continue
		}

		//This means that the first elements cannot be part of the LHS, so we ignore them and start forming our window from the next ones.
		if (diff > 1 || diff < -1) && newI == 0 {
			newI = j
			i = newI
			j = i + 1
			continue
		}

		if diff == 1 || diff == -1 {
			if nums[j-1] != nums[j] {
				newI = j
			}

			if count == 0 {
				count = j - i
			} else {
				count++
			}

			j++
			continue
		}

		if diff > 1 || diff < -1 {
			largest = updateLargest(largest, count)
			count = 0
			i = newI
			newI = j
			j = i + 1
			continue
		}
	}

	largest = updateLargest(largest, count)

	return largest
}
