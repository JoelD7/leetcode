# Problem
https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

Given a string `s`, find the length of the longest substring(A substring is a contiguous non-empty sequence of characters within a string.) without duplicate characters.


### Example 1:

    Input: s = "abcabcbb"
    Output: 3
    Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.

### Example 2:

    Input: s = "bbbbb"
    Output: 1
    Explanation: The answer is "b", with the length of 1.

### Example 3:

    Input: s = "pwwkew"
    Output: 3
    Explanation: The answer is "wke", with the length of 3.
    Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.


### Constraints:

    0 <= s.length <= 5 * 104
    s consists of English letters, digits, symbols and spaces.

# Solution
The solution involves using a sliding window of non-repeating characters. We use a hash map to detect when a character is duplicated and reduce the window until that character is no longer present on it. Conversely, the window will keep on increasing until we find a repeating character.

### Variables

- `left`: left pointer indicates the start of the window. It will only move when we detect a repeated character
- `set`: hash map that allows us to know when a character of the window is repeated
- `right`: right pointer, indicates the end of the window. This is the pointer that allows the expansion of the window. It will move on every iteration until we find a repeated character.

### Implementation

1. Iterate over `s` using `right`
2. If the character that right points to(`s[right]`) is already included in the window, remove it and increase the `left` pointer. Do this process until `s[right]` no longer appears repeated.
    1. Note when we detect that character `s[right]` is repeated, this doesn’t necessarily mean that both `right` and `left` point to the same character(which is what I thought when I first tried to solve this problem). Even though this is what will happen for example 1 where `s = abcabcbb`, it won’t be the case for `s = tmmzuxt`. In the case it **doesn’t** happen, we still need to move `left` until `s[right]` is no longer duplicated, even if that means ignoring other characters that aren’t duplicated. Why? Remember, we’re looking for a the largest **substring**, a sequence of contiguous characters. So we can’t just remove *only* the duplicated character because this might mean breaking the sequence.

       For example, for `s = abcabcbb`, when the window is between 3 and 6 meaning “`abcb`" even though we detected that “b” is duplicated we can’t just deleted the “b” in the middle in order to have “acb” because that would break the sequence. We need to also delete the first “a” even though is not duplicated in the current window. So we end up with → “`cb`”.

3. Add `s[right]` to the hash map.
4. Update the value of `max`.