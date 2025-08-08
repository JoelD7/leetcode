# Problem
https://leetcode.com/problems/ransom-note/description/


Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.


### Example 1:

    Input: ransomNote = "a", magazine = "b"
    Output: false

### Example 2:

    Input: ransomNote = "aa", magazine = "ab"
    Output: false

### Example 3:

    Input: ransomNote = "aa", magazine = "aab"
    Output: true


### Constraints:

    1 <= ransomNote.length, magazine.length <= 105
    ransomNote and magazine consist of lowercase English letters.

# Solution
1. Count each character in magazine and store the result in a map(`magazineFreq`). 

2. For each character in `ransomNote`, retrieve it from `maganize` by reducing its frequency by one, signalling we have used one of the available characters.
   - When a count reaches zero or if there is no character from `ransomNote` in `maganize`, return false inmediately as this means that either there aren't enough characters in `magazine` to build the `ransomNote`, or that one of the characters of `ransomNote` isn't in the `magazine` at all.