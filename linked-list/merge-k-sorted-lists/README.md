# Problem
https://leetcode.com/problems/merge-k-sorted-lists/

You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.



### Example 1:

    Input: lists = [[1,4,5],[1,3,4],[2,6]]
    Output: [1,1,2,3,4,4,5,6]
    Explanation: The linked-lists are:
    [
        1->4->5,
        1->3->4,
        2->6
    ]
    merging them into one sorted linked list:
    1->1->2->3->4->4->5->6

### Example 2:

    Input: lists = []
    Output: []

### Example 3:

    Input: lists = [[]]
    Output: []


### Constraints:

    k == lists.length
    0 <= k <= 104
    0 <= lists[i].length <= 500
    -104 <= lists[i][j] <= 104
    lists[i] is sorted in ascending order.
    The sum of lists[i].length will not exceed 104.

# Solution
This problem has the peculiarity of having a brute force implementation that has a 0ms runtime, beating 100% of other solutions. Here it is:

- *Brute force*

  Quite straightforward. We first combine all the elements of all the lists into an array, sort the array, and convert back the items to a linked list form. $O(n log n)$ solution.

    ```go
    type ListNode struct {
    	Val  int
    	Next *ListNode
    }
    
    func mergeKLists(lists []*ListNode) *ListNode {
    	arr := make([]int, 0)
    
    	var cur *ListNode
    	for _, list := range lists {
    		cur = list
    		for cur != nil {
    			arr = append(arr, cur.Val)
    			cur = cur.Next
    		}
    	}
    
    	sort.Ints(arr)
    
    	var result *ListNode
    	cur = result
    	for i, item := range arr {
    		if cur != nil {
    			cur.Val = item
    		} else {
    			result = &ListNode{
    				Val: item,
    			}
    			cur = result
    		}
    
    		if i < len(arr)-1 {
    			cur.Next = &ListNode{}
    		}
    		cur = cur.Next
    	}
    
    	return result
    }
    ```


---

We can obviously do better than that.

The problem can be solved by iteratively taking two lists out of `lists`, merging them together using the same solution as the [Merge Two Sorted Lists](../linked-list/merge-two-sorted-lists/README.md) problem, appending the result to the end of `lists`, and removing those two input lists from `lists`.

On each iteration we’ll have a shorter and shorter `lists` array.

```go
func mergeKLists(lists []*ListNode) *ListNode {
	//Edge case
	if len(lists) == 0 {
		return nil
	}

	//Edge case
	if len(lists) == 1 {
		return lists[0]
	}

	var mergedList *ListNode

	for len(lists) > 1 {
		mergedList = mergeTwoLists(lists[0], lists[1])

		if len(lists) > 2 {
  		//Remove the first two lists because they're already merged in "mergedList"
			lists = lists[2:]
			lists = append(lists, mergedList)
		} else {
			break
		}
	}

	return mergedList
}
```