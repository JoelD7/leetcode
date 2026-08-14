# Problem
https://leetcode.com/problems/longest-palindromic-substring/

Given a string `s`, return the longest in `s`.


### Example 1:

    Input: s = "babad"
    Output: "bab"
    Explanation: "aba" is also a valid answer.

### Example 2:

    Input: s = "cbbd"
    Output: "bb"


### Constraints:

    1 <= s.length <= 1000
    s consist of only digits and English letters.

# Solution

One of the ways to quickly figure out if a string is a palindrome is *expanding from the middle*, this is, starting at the middle index(or indeces if the string is even) and expanding to the right and the left while the letters pointed by the `right` and `left` pointers match. As long as this condition is maintained, the string is a palindrome.

We do this for every index of `s` and return the string of the maximum length we find.

Note that since we might have even and odd palindromes we must check for both for every index. The difference between checking even and odd palindromes, is that for odds the `right` and `left` pointers start at the same index, while for evens don’t.
 
