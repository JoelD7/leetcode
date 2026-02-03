# Problem

Given two strings s1 and s2, return true if s2 contains a permutation(a rearrangement of all the characters of a string) of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.


### Example 1:

    Input: s1 = "ab", s2 = "eidbaooo"
    Output: true
    Explanation: s2 contains one permutation of s1 ("ba").

### Example 2:

    Input: s1 = "ab", s2 = "eidboaoo"
    Output: false


### Constraints:

    1 <= s1.length, s2.length <= 10^4
    s1 and s2 consist of lowercase English letters.

# Solution
### Variables

- `freqS1`: array that stores the frequencies of characters in `s1`
- `freqS2`: array that stores the frequencies of characters in `s2`

### Note on the frequency ~~maps~~ arrays

Note that we’re using arrays instead of maps here, specifically arrays of bytes. Why is that? Using an array of bytes alllows us to easily compare the frequencies of both strings. We create an array with 26 slots, one for each letter of the alphabet, in order. Slot 1(index 0) corresponds to the letter “a” while slot 26(index 25) corresponds to letter “z”.

However, the character “a” in `bytes` is 97, character “b” is 98, “c” is 99, etc., so how do we transform that to an array index? We subtract the byte “a” from each character.

```go
  a, b, c := 'a', 'b', 'c' //these vars are of type "byte"
  fmt.Println(a - 'a') //Output: 0
  fmt.Println(b - 'a') //Output: 1
  fmt.Println(c - 'a') //Output: 2
```

### Implementation

So permutation of a string basically means just the reordering of the letters of the string. That means all the permutations of a string are just the anagrams of the string, with all the letters included, just in another way.
And we have to find if such a substring exists in `s2` which is a permutation of `s1`, which means we have to find a substring in `s2` such that the frequency of the characters in that substring is same as the frequency of the characters in `s1`.

1. The first step is to find the freq of all characters of `s1`.
2. Then we will be starting with a window of size 1 initially in `s2`. That means `start=0`, `end=0`.
3. Along with this, we will also be maintaining the frequency of the characters in the current window.
4. We will be continuing the below steps untill we reach a situation where the end of the window reaches the end of `s2`. That means we will be doing the steps while `end<length` of `s2`.
    1. Increase the frequency of the character which is just now newly included inside the window. That means increase the frequency of `s2[end]`.
    2. Now check if the frequency of the characters in the current window, is same as the frequency of characters int `s1`. But this is only possible if the length of current substring(window) is same as the length of s1. If this is true, then we can directly return true from here.
    3. If the frequency doesnt match, we have to change the window:-
        1. If the length of window in less than the length of `s1`, then we should increase the length of the the window by increasing the end. So `end+=1`.
        2. If not, that means length of window is greater than or equal to the length of `s1`, so we will need to move to a new window. For that we will have to move start to the next character but before that we will have to decrease the frequency of start character from the curr window frequency. That means
            1. Decrese the frequency of `s2[start]`.
            2. Move start to the next element.
            3. Move end to the next element.
5. If the algorithm did not return true for any window, then we will reach here(out of the loop). This will mean that we did not find any such substring in `s2`. So return `false`.
