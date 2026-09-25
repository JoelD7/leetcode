# Problem
https://leetcode.com/problems/encode-and-decode-strings/description/

Design an algorithm to encode a list of strings to a string. The encoded string is then sent over the network and is decoded back to the original list of strings.

**Machine 1 (sender)** has the function:

```
String encode(List<String> strs) {
    // ... your code
    return encoded_string;
}
```

**Machine 2 (receiver)** has the function:

    List<String> decode(String encoded_string) {
        // ... your code
        return decoded_strs;
    }

So **Machine 1** does:

    String encoded_string = encode(strs);

    and **Machine 2** does:

    List<String> decoded_strs = decode(encoded_string);

`decoded_strs` in Machine 2 should be the same as the input `strs` in Machine 1.

Implement the encode and decode methods.

### Example 1:

    Input: strs = ["Hello","World"]
    
    Output: ["Hello","World"]

**Explanation**:
    
    Solution solution = new Solution();
    String encoded_string = solution.encode(strs);
    
    // Machine 1 ---encoded_string---> Machine 2
    
    List<String> decoded_strs = solution.decode(encoded_string);


### Example 2:
    
    Input: strs = [""]
    
    Output: [""]


### Constraints:

    0 <= strs.length < 100
    0 <= strs[i].length < 200
    strs[i] contains any possible characters out of 256 valid ASCII characters.

# Solution
### Rationale

The solution is pretty straightforward: combine the strings of `strs` during encoding, and split them during decoding. What makes this problem interesting is *how* to combine the strings. You can’t use a string delimiter because according to the problem’s constraints “`*strs[i]` contains any possible characters out of `256` valid ASCII characters*”, so using commas or any other character is out of the question. What we do instead is using something called **Length-Prefix Encoding**: in the encoded string, each of the strings are prefixed by their length so that when decoding, we can know up until which position of `encoded` a valid string ends and another one begins.

### Algorithm

1. Encode `strs`.
    1. For `input = "lint", "code", "love", "you”`.  The result is: “4#lint4#code4#love3#you”
2. Decode
    1. Use two pointers, `i` and `j`. We first need to find the length of the first string, which is the point where the first “#” appears. Here we have to note an important point: “#” has to be *after* the length because this is what divides the length of the string from the string. If we have valid string like “8ab” and we put the “#” before the length, we’d end up with an encoded string like `#38ab`. Since numbers are allowed to form part of the string, in this case we can’t know for sure whether the string has length 3 or 38. For this reason we put the “#” before, to reduce ambiguity → `3#8ab`.
    2. Extract and parse the length
    3. Use the length and the position of the `j` index to know the limits of the string. Note that the starting position is always 1 place after the “#”.
    4. Continue with the next string, setting `i = end`, because when we do string splitting in Golang, the “end” pointer is not included.