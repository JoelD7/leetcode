# Problem
https://leetcode.com/problems/remove-outermost-parentheses/

A valid parentheses string is either empty "", "(" + A + ")", or A + B, where A and B are valid parentheses strings, and + represents string concatenation.

    For example, "", "()", "(())()", and "(()(()))" are all valid parentheses strings.

A valid parentheses string s is primitive if it is nonempty, and there does not exist a way to split it into s = A + B, with A and B nonempty valid parentheses strings.

Given a valid parentheses string s, consider its primitive decomposition: s = P1 + P2 + ... + Pk, where Pi are primitive valid parentheses strings.

Return s after removing the outermost parentheses of every primitive string in the primitive decomposition of s.



Example 1:
    
    Input: s = "(()())(())"
    Output: "()()()"
    Explanation:
    The input string is "(()())(())", with primitive decomposition "(()())" + "(())".
    After removing outer parentheses of each part, this is "()()" + "()" = "()()()".

Example 2:

    Input: s = "(()())(())(()(()))"
    Output: "()()()()(())"
    Explanation:
    The input string is "(()())(())(()(()))", with primitive decomposition "(()())" + "(())" + "(()(()))".
    After removing outer parentheses of each part, this is "()()" + "()" + "()(())" = "()()()()(())".

Example 3:

    Input: s = "()()"
    Output: ""
    Explanation:
    The input string is "()()", with primitive decomposition "()" + "()".
    After removing outer parentheses of each part, this is "" + "" = "".



Constraints:

    1 <= s.length <= 105
    s[i] is either '(' or ')'.
    s is a valid parentheses string.

# Solution
I noticed the fact that a valid group of parentheses can be detected when using a stack: push when detecting an opening parentheses and popping when detecting a closing one. When the stack has zero elements on it, it means that all opening parentheses have a matching closing parentheses. Each time the stack reaches zero length, it means a parentheses group has been found. When that happens, add the group to the list of `primitives`, removing the first and last characters of the string which will always be the outermost parentheses.