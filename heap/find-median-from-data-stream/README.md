# Problem
https://leetcode.com/problems/find-median-from-data-stream/description/

The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value, and the median is the mean of the two middle values.

    For example, for arr = [2,3,4], the median is 3.
    For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.

Implement the MedianFinder class:

    MedianFinder() initializes the MedianFinder object.
    void addNum(int num) adds the integer num from the data stream to the data structure.
    double findMedian() returns the median of all elements so far. Answers within 10-5 of the actual answer will be accepted.



### Example 1:

    Input
    ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
    [[], [1], [2], [], [3], []]
    Output
    [null, null, null, 1.5, null, 2.0]
    
    Explanation
    MedianFinder medianFinder = new MedianFinder();
    medianFinder.addNum(1);    // arr = [1]
    medianFinder.addNum(2);    // arr = [1, 2]
    medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
    medianFinder.addNum(3);    // arr[1, 2, 3]
    medianFinder.findMedian(); // return 2.0



### Constraints:

    -105 <= num <= 105
    There will be at least one element in the data structure before calling findMedian.
    At most 5 * 104 calls will be made to addNum and findMedian.

# Solution
### Intuition

To get the median we only care about the middle or two middle values in a sorted integer sequence. If we split a sorted sequence in two halves, the last item on the lowest half will be in the middle and the first item of the highest half will also be in the middle. If both halves are of the same size, those two elements will conform the median. If one of them is larger, the corresponding extreme element will be the median.

The data structure that best allow us to model this behaviour is a heap. There we only care about the max or min item of a sequence, and it allows us to fetch it in $O(1)$ time.

Store the lower half of elements in a max-heap and the higher half of elements in a min-heap. The top of the max-heap would have the largest number across the lower half, and the top of the min-heap would have the lowest number across the higher half. **Either one or two of these top of heaps is the median, *so we must always keep both heaps balanced so that this is always true***.

To accommodate uneven number of elements, we choose one of the heaps to always be the largest one if required. Here we choose max-heap. By organizing elements in this way, fetching the median is done in $O(1)$ time because the median would either be the top of max-heap, or a division between the top of both heaps. Here is an example:

!image.png

Using heaps is valid here because we only care about the relative ordering of the middle elements. It doesn’t matter that the other ones are unsorted.

### `AddNum` function

1. Add `num` to the max-heap
2. Remove the top of the max-heap and add it to the min-heap. We do this to maintain both heaps of the same size or, making the max-heap just one element larger.
    1. If doing this makes the min-heap be the larger heap, move the top of the min-heap to the max-heap. We do this because of two reasons:
        1. Min-heap can’t be larger than max-heap, therefore we need to move one element of it to the max-heap
        2. We move the lowest number of the higher half to maintain the property of having the lowest half on the max-heap, and the highest half on the min-heap.

### `FindMedian` function

If the max-heap is the largest, return the top. Else, get the average of max-heap and min-heap’s tops.