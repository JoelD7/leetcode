# Problem

https://leetcode.com/problems/koko-eating-bananas/description/

Koko loves to eat bananas. There are `n` piles of bananas, the `i^th` pile has `piles[i]` bananas. The guards have gone
and will come back in `h` hours.

Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas
from that pile. If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas
during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer k such that she can eat all the bananas within h hours.

### Example 1:

    Input: piles = [3,6,7,11], h = 8
    Output: 4

### Example 2:

    Input: piles = [30,11,23,4,20], h = 5
    Output: 30

### Example 3:

    Input: piles = [30,11,23,4,20], h = 6
    Output: 23

### Constraints:

    1 <= piles.length <= 104
    piles.length <= h <= 109
    1 <= piles[i] <= 109

# Solution

Is important to note that adding up all the piles and dividing them by `h` doesn’t yield an accurate result because of
the constraint of the problem: “*If the pile has less than `k` bananas, she eats all of them instead and will not eat
any more bananas during this hour.*” Here is the breakdown of how many hours it takes to eat each pile with a speed of *
*18**:

- **Pile 1 (30 bananas):** 30/18 requires **2 hours** (18 in the first hour, 12 in the second).
    - She can’t eat the 12 bananas from pile 1 and then eat 6 from pile 2 because pile 2 has “less than `k`(18)
      bananas”, so Koko has to wait one hour before getting into pile 2. Therefore 18 is not optimal
- **Pile 2 (11 bananas):** 11/18 requires **1 hour**.
- **Pile 3 (23 bananas):** 23/18 requires **2 hours** (18 in the first hour, 5 in the second).
- **Pile 4 (4 bananas):** 4/18 requires **1 hour**.
- **Pile 5 (20 bananas):** 20/18 requires **2 hours** (18 in the first hour, 2 in the second).

**Total time taken:** 2+1+2+1+2= **8 hours**.

> 💡 In other words, after you decide on `k` you have to commit to it even when a pile is split in two.


---

To get the time Koko would need to eat a pile we do `(pile - 1) / speed + 1`. The condition function returns `true` if
after adding up all those times using a specific `speed` we get a time smaller or equal than `h`.

We use binary search to efficiently find the minimum value that satisfies that condition. The value of `left` we end up
with after the loop exits, i.e., after `right` and `left` meet each other, is our return value. Why? The value of `left`
we return would eventually have been `mid` in some point during the iteration because remember that by default, a
division between integers rounds down, so let’s say if `right = 6` and `left = 5`, the middle value will be `5`(the same
as `left`), then it will be evaluated by the condition, `right` will now be `5` and the loop will exit because right and
left are now equal.

When a `mid` value satisfies the condition, we shrink the search range and make it closer to `left` because after all,
we’re minimizing, we want to check if there is a smaller value than the `mid` we just checked that also satisfies the
condition. Note that if that **never happens**, the final value of `left` will be the last `mid` we evaluated to `true`
because of the `else` statement → `left = mid + 1`, i.e., `left` will catch up to `right` making the loop exit.

If `mid` doesn’t satisfy the condition, there is no point in searching the elements before it so we
make `left = mid + 1`. Doing this back-and-forth will produce the minimum value that satisfies the condition, whether
that happens by `right` getting closer to `left`, or `left` going up to meet `right`.

Here is the solution code:

```python
def minEatingSpeed(piles: List[int], H: int) -> int:
    def feasible(speed) -> bool:
				# This sums the results of doing (pile - 1) / speed + 1 for all piles  
        return sum((pile - 1) / speed + 1 for pile in piles) <= H  # faster

    left, right = 1, max(piles)
    while left < right:
        mid = left  + (right - left) // 2
        if feasible(mid):
            right = mid
        else:
            left = mid + 1
    return left
```