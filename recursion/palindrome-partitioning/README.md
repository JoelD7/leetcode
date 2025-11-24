# Problem
https://leetcode.com/problems/palindrome-partitioning

Given a string `s`, partition `s` such that every of the partition is a
palindrome . Return all possible palindrome partitioning of `s`.


### Example 1:

    Input: s = "aab"
    Output: [["a","a","b"],["aa","b"]]

### Example 2:
    
    Input: s = "a"
    Output: [["a"]]


### Constraints:

    1 <= s.length <= 16
    s contains only lowercase English letters.

 
# Solution
### Variables

- `comb`: current combination of palindrome substrings of `s`
- `s`: This is the input, however, it gets reduced on every recursive call so in the context of the backtracking function is a substring of the original `s`.
- `res`: list of palindrome substrings.

### Intuition

The way to obtain “*all possible palindromes of `s`*” is by taking a character of `s` starting with the first one, adding it to `comb` if it’s a palindrome and do the same with the **rest** of the string, recursively. Doing that continously until we can’t no more(`s` is empty), and then repeat the process again by starting with the first two characters of `s` to see if there are other possible ways to get palindrome substrings. You can see clearly here that this is a backtracking problem: we explore every possible path recursively, and then backtrack out of those paths to try other ones.

### Algorithm

1. **Base case**: `s == ""`, or in other words, after there are no more parts of the input to explore. On this point we can add a copy of `comb` to the list of results.
2. Iterate over `s` starting at 1 not 0 because we’re using the index of the loop to get a *substring* of `s`, not a character and to get the first substring we must do `s[:1]`. If the substring is palindrome then…
    1. Add it to `comb`
    2. Call `backtrack` recursively with the **rest** of `s` as input. We do this because we don’t want to generate substrings with repeated characters.
    3. We backtrack by removing the current substring of `comb`, which will allows to consider other possible ways to form substrings of `s` that might be palindrome. For example, for `s = “aab”`, if we already considered the first `a`(s[:1]) on it’s own, we remove it from `comb` and add `aa`(s[:2]) to the combination list in the next iteration. If we hand’t backtracked, i.e. removed s[:1] out of `comb`, on the next iteration we would be considering the first `a` twice: `a` and `aa`.