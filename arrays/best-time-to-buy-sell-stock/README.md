# Problem
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

You are given an array `prices` where `prices[i]` is the price of a given stock on the `ith` day.

You want to maximize your profit by choosing a **single day** to buy one stock and choosing a **different day in the future** to sell that stock.

Return *the maximum profit you can achieve from this transaction*. If you cannot achieve any profit, return `0`.


**Example 1:**

    Input: prices = [7,1,5,3,6,4]
    Output: 5
    Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
    Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

**Example 2:**

    Input: prices = [7,6,4,3,1]
    Output: 0
    Explanation: In this case, no transactions are done and the max profit = 0.


**Constraints**:

    1 <= prices.length <= 105
    0 <= prices[i] <= 104

# Solution
To obtain the maximum profit we should buy the stock at the lowest possible price and sell it at the highest price. We emphasize "possible" because the selling price should be positioned after the buying price in the array.

## Implementation
The buying price(`buyPrice`) should be the lowest possible. For that reason, we iterate over the array updating `buyPrice` only when we find a lowest value than the current one. 

On each iteration we subtract the current element(the potential `sellPrice`) from `buyPrice` and update the value of `profit`, a variable that will hold the maximum profit **so far**. If the result of that operation is larger than the current `profit`, we update it's value. 

After iterating over the whole array, we would have reached the maximum profit. 