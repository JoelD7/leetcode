# Problem
https://leetcode.com/problems/valid-parentheses/

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.



Example 1:

    Input: s = "()"
    Output: true

Example 2:

    Input: s = "()[]{}"
    Output: true

Example 3:

    Input: s = "(]"
    Output: false



Constraints:

    1 <= s.length <= 104
    s consists of parentheses only '()[]{}'.

# Solution
- Iterate over `s` to analyze each character `char`
- If `char` is an opening tag, add it to the stack
- If `char` is a closing tag, pop from the stack and that popped character must be the opposing tag to `char`, if it isn't, return false.
- At the end of iteration return true

## Caveats
- If `char` is a closing tag and the stack is empty, that means that `s` is invalid because
there is no opening tag that matches. 
- If after finishing processing `s` the stack still has elements, then `s` is invalid. Each
opening tag should be popped from the stack when `s` contains the corresponding closing tag.