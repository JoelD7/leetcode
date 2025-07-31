# Problem
https://leetcode.com/problems/repeated-dna-sequences/description

The DNA sequence is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T'.

    For example, "ACGAATTCCG" is a DNA sequence.

When studying DNA, it is useful to identify repeated sequences within the DNA.

Given a string `s` that represents a DNA sequence, return all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule. You may return the answer in any order.


### Example 1:

    Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
    Output: ["AAAAACCCCC","CCCCCAAAAA"]

### Example 2:

    Input: s = "AAAAAAAAAAAAA"
    Output: ["AAAAAAAAAA"]


### Constraints:

    1 <= s.length <= 105
    s[i] is either 'A', 'C', 'G', or 'T'.

# Solution
Scan the string `s` and extract every 10-character substring. Use a hash map to count the occurrences of each substring. Add any substring with a count of exactly 2 to your results.


## Variables
- `i`: Index that indicates the start of the window/sequence/substring
- `seq`: The current candidate sequence of `s` we're analyzing. 
- `count`: Counter of the times a sequence has appeared on `s`. We use it to know weather we have already added the sequence to the result set or not. Note that the count for each sequence will never increase beyond 2 because it doesn't have to. The problem asks for sequences that appear **2 times** or more. If we have sequence that appears 2-times, that's it!
- `m`: Hash map used to holds the found sequences.

## Algorithm
Iterate over the string up to the index `i` `len(s)-10`(including). After that point there can't be any other substrings starting at `i` that have 10 characters.

---
If the sequence is already on the map we know it's repeated. We only add it to the result set if it's `count` is 1 because this means this is the first time we have added it.

We increase the value of `count` so that if we find the same sequence later, we don't add it to `result` again and produce a duplicate.
```go
if ok && count == 1 {
    result = append(result, seq)
    m[seq] = count + 1
    continue
}
```
---
This conditional acts as a control mechanism. It prevents that a sequence that has already been added to `result` has it's count reset to 1.
```go
...
if count > 1 {
    continue
}

m[seq] = 1

```
