package main

func maxProfit(prices []int) int {
	var profit, curProfit, price int
	buyPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		price = prices[i]

		if price < buyPrice {
			buyPrice = price
		}

		curProfit = price - buyPrice
		if curProfit > profit {
			profit = curProfit
		}
	}

	return profit
}
