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
Use a sliding window with two pointers: `l`(buying price) and `r`(selling price). The loop will stop once we can’t find a new selling price.

While `l` keeps pointing to a smaller value than `r`, we update the current `maxProfit` if necessary. On each iteration we keep enlarging the window by increasing the value of `r`, to see if we can find a higher selling price that allow us to get a better profit. Once we find a smaller selling price(`r`) than the buying price, `l` will point to the current value of `r` and `r` will increase by 1. `l` should be now equal to `r` because we have already calculated the profit of all the values between `l` and `r`, so there is no point in increasing `l` just by 1(a much in-depth explanation about this logic bellow).

The key insight is always maintaining the left pointer(buying price) lower than the right pointer(selling price).

## Why do `left = right` instead of `left = left + 1`?

While solving the problem one of the things I struggled to understand was why increasing `left` by 1 didn’t worked. Let’s analyze example 1 to explain this…

```go
For, 
prices = [7, 1, 5, 3, 6, 4]
left = 0 -> points to 7
right = 1 -> points to 1
```

According to the solution’s logic, since `prices[right] < prices[left]` we should move `left` to the position of `right`, effectively ignoring a buying price of 7 “forever”. This is what I couldn’t understand. What if there is a another much higher price of 18 in the future, then our profit would be 11 instead of 5(the expected output of example 1), a much better profit. Well, the fact that we just found a lower price than `left`(the value of `right`)  means that we will obviously get a higher profit with this new value rather than the current value of `left` because we just found a lower buying price. Following the hypothetical future price of 18, our **profit would be 17(18 - 1) instead of 11(18 - 7)**. **A lower buying price = higher profit OBVIOUSLY!**