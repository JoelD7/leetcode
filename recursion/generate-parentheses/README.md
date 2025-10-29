# Problem
https://leetcode.com/problems/generate-parentheses/description/

Given `n` pairs of parentheses, write a function to generate all combinations of well-formed parentheses.


### Example 1:

    Input: n = 3
    Output: ["((()))","(()())","(())()","()(())","()()()"]

### Example 2:

    Input: n = 1
    Output: ["()"]



### Constraints:

    1 <= n <= 8

# Solution
We use two integer variables `open` & `closed` to see how many `'('` & `')'` are in the current string.

The solution builds a parentheses string step by step depending on a set of conditions. The intuition here is:

1. **We can only add a closing/right parentheses `)` to the string once we know for sure there is a corresponding opening/left parentheses `(`**. For that reason we first add all the left parentheses we can, and then make sure to check if `closed < open` before adding closing parentheses.
2. **the amount of opening parentheses must be equal to the amount of closing parentheses**. This is why we keep increasing `closed` until it reaches `open`.
3. **the base case is `len(s) == n * 2`**. The problem requires combinations made with `n` **pairs**, once `s` reaches this length it means we have `n` pairs of parentheses in the string.

   Another more clear way to understand the base case is having the condition be `open == closed == n`, which is saying: “when the amount of opening and closing parentheses is equal to `n`, we have a valid combination of parentheses”.

4. **create a two-branched recursive tree**. If we walkthrough the algorithm upto this point using the conditions we defined earlier, we notice that almost every time we have two *valid* choices: add a opening parentheses or a closing one. Every decision taken leads to the same question again: should we add an opening or a closing parentheses? If we draw the sequence of decisions for `n = 2` we arrive at a decision tree like the following:

    ```python
                                (0, 0, '')
                                    |	
                                (1, 0, '(')  
                               /           \
                        (2, 0, '((')      (1, 1, '()')
                           /                 \
                    (2, 1, '(()')           (2, 1, '()(')
                       /                       \
                (2, 2, '(())')                (2, 2, '()()')
                          |	                             |
                res.append('(())')             res.append('()()')
       
    ```

   From this decision tree alone is kind of obvious that we’re dealing with a recursive algorithm with two separate branches, each representing the two choices we can take while building a valid combination.

> **Tip**
> 
> To better understand this solution try to not think too much about how a recursive calls relates to another, but focus on what a *single* call is doing instead. We first add a opening, then a closing parentheses. The subsequent recursive calls will handle what to do with the opening parentheses you just added, the same applies for when you add a closing parentheses.