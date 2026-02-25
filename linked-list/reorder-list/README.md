# Problem
https://leetcode.com/problems/reorder-list/

![img.png](img.png)![img_1.png](img_1.png)

### Constraints:

    The number of nodes in the list is in the range [1, 5 * 104].
    1 <= Node.val <= 1000

# Solution
For list `[1,2,3,4,5,6,7]` we need to return `[1,7,2,6,3,5,4]`. We can note, that it is actually two lists `[1,2,3,4]` and `[7,6,5]`, where elements are interchange. So, in a nutshell, the solution to this problem is splitting the input in two lists: one with the first half of the elements, and another with the second half of the elements but in reversed order, and then merging both of those lists in intertwining order.

1. **Find the middle of or list** - be careful, it needs to work properly both for even and for odd number of nodes. For this we can use `slow/fast` iterators trick, where `slow` moves with speed `1` and `fast` moves with speed `2`. Then when `fast` reches the end, `slow` will be in the middle, as we need.
2. **Reverse the second part of linked list**. The second part of the list starts at `slow.next`, because remember: `slow` is now at the middle of the list thanks to what we did in the previous step.
    1. **Do not forget to use `slow.next = nil`, to clearly separate the two lists**. Setting `slow.next = nil` is crucial because it cleanly cuts the original list into two completely separate linked lists. If you skip this step, the last node of your first half (`slow`) will still point to the first node of the original second half, and this will produce weird behaviours.
3. **Finally, we merge the two lists in a zigzag manner**. We take one node from the first half, then one from the reversed second half, and so on, until all nodes are merged into a single, reordered list.