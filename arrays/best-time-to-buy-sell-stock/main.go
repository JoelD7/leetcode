package main

// Input: prices = [2,4,1]
// Output: 2
func maxProfit(prices []int) int {
	minPosition := 0
	min := 10000 //Max value according to the problem description

	//Get position of the lowest value
	for i, price := range prices {
		if price < min {
			min = price
			minPosition = i
		}
	}

	max := 0
	//Get largest value
	for i := minPosition + 1; i < len(prices); i++ {
		if prices[i] > max {
			max = prices[i]
		}
	}

	if max-min > 0 {
		return max - min
	}

	return 0
}
