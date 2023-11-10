# Problem
https://leetcode.com/problems/top-k-frequent-elements/

Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.



Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:

Input: nums = [1], k = 1
Output: [1]



Constraints:

    1 <= nums.length <= 105
    -104 <= nums[i] <= 104
    k is in the range [1, the number of unique elements in the array].
    It is guaranteed that the answer is unique.



Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
# Solution
First, build a frequency map that keeps track of the total times a value appears in `nums`. In this map the keys
are the elements of `nums` and the values are the frequency of those elements. 

Second, since we need to obtain the `k` more frequent values from `nums`, we need to find a way to sort the 
frequency map by its values. Sorting a map by keys in Golang is easy, by values its not. To do this I created
a struct `customMap` with two fields: `value` and `frequency`. This custom type would keep the same data as 
one entry in the frequency map. I then aggregated all the frequency map's data in an `[]customMap`, 
took advantage of the fact that in Go is possible to customize the sorting of an array of type `X`, and sorted
it by `frequency`. 

The use of this custom type array allowed me to sort the frequency map by its values(frequency), and because
each element in `[]customMap` is an object, I could easily obtain the `value` associated with the `frequency`.