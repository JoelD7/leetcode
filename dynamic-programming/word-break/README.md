# Problem
https://leetcode.com/problems/word-break/

Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.



### Example 1:

    Input: s = "leetcode", wordDict = ["leet","code"]
    Output: true
    Explanation: Return true because "leetcode" can be segmented as "leet code".

### Example 2:

    Input: s = "applepenapple", wordDict = ["apple","pen"]
    Output: true
    Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
    Note that you are allowed to reuse a dictionary word.

### Example 3:

    Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
    Output: false



### Constraints:

    1 <= s.length <= 300
    1 <= wordDict.length <= 1000
    1 <= wordDict[i].length <= 20
    s and wordDict[i] consist of only lowercase English letters.
    All the strings of wordDict are unique.

# Solution
### Rationale

The problem exhibits dynamic programming qualities, as we can solve it by continuously asking “*can this string be formed with the words of `wordDict`?”* with ever reducing smaller versions of the original `s`. In other words, the result can be obtained by combining the solutions to multiple subproblems.

We return `true`  “*if `s` can be segmented into a space-separated sequence of one or more dictionary words*”. The keyword here is *sequence*. In simpler terms, this means that substrings of `s` have to maintain order.  For example, for `s = ccbb` and `wordDict = ["bc", "cb"]` extracting the middle “cb” leaving `s = cb`, is **wrong** because we’d effectively be forming a word that didn’t exist in the original `s`. So we need to extract from the extremes.

### Algorithm

In a nutshell, the algorithm is this: check if `s` has any of the words as a prefix, if it does, check the remainder equally. Do this in a recursive manner until there is no more `s` to check, i.e., is empty. The fact that `s` is empty implicitely means we can form `s` with the words of `wordDict`, because the only condition we have to reduce `s` is that a valid `word` is found.

1. Inside the recursive functoin we first check if `str` is empty.
2. Check is `memo` has a calculated value for `str`.
    1.  `s` can be split in different manners depending the words on `wordDict`. Sometimes we might end up with the same substring of `s` in distinct iterations. To optimize the algorithm and reduce recursive calls, we memoize the already calculated values of substrings of `s` and return them upon producing the same substring several times.
3. Iterate over `wordDict` to test all words
    1. If `str` has `word` a prefix, we made a recursive call with the remainder