# Problem
https://leetcode.com/problems/coin-change/description/

You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.



### Example 1:

    Input: coins = [1,2,5], amount = 11
    Output: 3
    Explanation: 11 = 5 + 5 + 1

### Example 2:
    
    Input: coins = [2], amount = 3
    Output: -1

### Example 3:

    Input: coins = [1], amount = 0
    Output: 0


### Constraints:

    1 <= coins.length <= 12
    1 <= coins[i] <= 231 - 1
    0 <= amount <= 104

# Solution
### Rationale

This is an optimization problem so we can use DP to solve it.

The “easiest” way to build our solution is figuring out the minimum amount of coins needed to make an amount 0, then 1, 2, 3,…, up to `amount`. Is a bottom-up approach. It’s a lot easier to know the minimum amount of coins need to make up to `amount` when we already know the minimum for `amount - 1`.

So we start a loop from `0...amount`, calculating the min amount of coins needed for every amount in that range. At the end we just return the last calculated value.

### Variables

- `dp`: This is our DP array. It holds the min amount of coins needed to make up to `i` amount, where `i` is the index of the array.
- `MaxInt`: Constant representing the starting value of every index of our `dp` array. It is initialized as `amount + 1` to indicate an “invalid” value, since the max amount for this problem is `amount`.
- `curAmount`: current amount being calculated in the loop.
- `rem`: remainder of `curAmount - coin`. We use this amount to know how many more coins we need to complete `curAmount` after taking a `coin`.

### Algorithm

1. Initialize `dp` and set `dp[0] = 0`. Why? No amount of coins are required to reach an amount of 0.
2. Iterate from `1...amount`(inclusive)
    1. For every amount `curAmount`, you will iterate over all the coins trying to find the min amount of coins needed to get to `curAmount`
        1. We ignore every `coin` that produces a negative `rem`, because a negative amount is invalid
        2. Then we calculate the min amount of coins need to reach `curAmount`, which is the minimum between current value of `dp[curAmount]` or `dp[rem]+1`.

           **Why do we need to acount for `dp[curAmount]` in the minimum calculation if this is the value we’re looking for?** Because we test for every coin in `coins`, not just the first or last one. This means that after trying with a coin, `dp[curAmount]` will hold the minimum value to reach `curAmount` *so far*. We might discover a new minimum with another coin later. We will know the actual minimum *after* trying with every possible coin.

           **What does `dp[rem]+1` mean?** `dp[rem]` represents the minimum amount of coins need to reach the remainder of `curAmount` *after* subtracting the current `coin`. The `+1` indicates the `coin` we just used. Let’s use an example:

           Imagine your `coins = [1, 2, 5]` and you are trying to make an `amount` of **7** (`i = 7`).

           Let's say you decide to test the **5-cent coin** (`c = 5`).

            - You take one 5-cent coin and put it on the table. **(This is the `+ 1`)**.
            - You originally needed 7. You just used 5. You still need to make up **2** (`curAmount - coin`, or `7 - 5`).
            - You look at your previously calculated records to see the cheapest way to make 2. Your table says `dp[2]` takes exactly **1** coin (a single 2-cent coin).
            - Your total coins used = the coins to make 2 (`dp[2]`) plus the 5-cent coin you put on the table (`+ 1`).
            - Total: `1 + 1 = 2` coins.
3. After doing all that, we return `-1` if `dp[amount]` is still `MaxInt` because that means we couldn’t find a minimum. Otherwise return `dp[amount]`.