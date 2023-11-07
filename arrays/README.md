# Solution
I noticed an interesting pattern: the repeated numbers that appear in the prefix sum array(`prefixSum`) of the input
array (`arr`) indicate the indices of the subarray that sum to 0, with the condition that we have to add 1 to the first idnex.
For example, let's say we have the following input:
```go
arr := []int{4, 2, -3, 1, 6}
prefixSum of arr: [4, 6, 3, 4, 10]
```

Notice that the number 4 appears twice in the prefix sum array: first in at index 0, then at index 3. 
This means that the subarray zero sum of `arr` goes from index 1 to index 3, inclusive: `[2, -3, 1]`.