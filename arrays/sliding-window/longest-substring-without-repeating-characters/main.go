package main

func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]bool)
	left, right, maxLength := 0, 0, 0
	var isDuplicated bool

	for right < len(s) {
		_, isDuplicated = set[s[right]]
		for isDuplicated {
			delete(set, s[left])
			left++
			_, isDuplicated = set[s[right]]
		}

		set[s[right]] = true
		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
		right++
	}

	return maxLength
}
