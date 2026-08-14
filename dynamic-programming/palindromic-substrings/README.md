# Problem
https://leetcode.com/problems/palindromic-substrings/

Given a string s, return the number of palindromic substrings in it.

A string is a palindrome when it reads the same backward as forward.

A substring is a contiguous sequence of characters within the string.



### Example 1:
    
    Input: s = "abc"
    Output: 3
    Explanation: Three palindromic strings: "a", "b", "c".

### Example 2:

    Input: s = "aaa"
    Output: 6
    Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".


### Constraints:

    1 <= s.length <= 1000
    s consists of lowercase English letters.

# Solution
One of the ways to quickly figure out if a string is a palindrome is *expanding from the middle*, this is, starting at the middle index(or indeces if the string is even) and expanding to the right and the left while the letters pointed by the `right` and `left` pointers match. As long as this condition is maintained, the string is a palindrome.

We do this for every index of `s` and increase a counter variable for every palindrome we find.

Note that since we might have even and odd palindromes we must check for both for every index. The difference between checking even and odd palindromes, is that for odds the `right` and `left` pointers start at the same index, while for evens don’t.