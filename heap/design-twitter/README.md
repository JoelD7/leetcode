# Problem
https://leetcode.com/problems/design-twitter/

Design a simplified version of Twitter where users can post tweets, follow/unfollow another user, and is able to see the 10 most recent tweets in the user's news feed.

Implement the Twitter class:

    Twitter() Initializes your twitter object.
    void postTweet(int userId, int tweetId) Composes a new tweet with ID tweetId by the user userId. Each call to this function will be made with a unique tweetId.
    List<Integer> getNewsFeed(int userId) Retrieves the 10 most recent tweet IDs in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user themself. Tweets must be ordered from most recent to least recent.
    void follow(int followerId, int followeeId) The user with ID followerId started following the user with ID followeeId.
    void unfollow(int followerId, int followeeId) The user with ID followerId started unfollowing the user with ID followeeId.



### Example 1:

    **Input**
    ["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"]
    [[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]
    **Output**
    [null, null, [5], null, null, [6, 5], null, [5]]
    
    **Explanation**
    Twitter twitter = new Twitter();
    twitter.postTweet(1, 5); // User 1 posts a new tweet (id = 5).
    twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5]. return [5]
    twitter.follow(1, 2);    // User 1 follows user 2.
    twitter.postTweet(2, 6); // User 2 posts a new tweet (id = 6).
    twitter.getNewsFeed(1);  // User 1's news feed should return a list with 2 tweet ids -> [6, 5]. Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
    twitter.unfollow(1, 2);  // User 1 unfollows user 2.
    twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5], since user 1 is no longer following user 2.



### Constraints:

    1 <= userId, followerId, followeeId <= 500
    0 <= tweetId <= 104
    All the tweets have unique IDs.
    At most 3 * 104 calls will be made to postTweet, getNewsFeed, follow, and unfollow.
    A user cannot follow himself.

# Solution
- `Twitter` type. It has
    - A hash table `m` that maps `userID` to users.
- `User` type. Each user has:
    - List of `posts` he has made sorted from most to least recent.
    - List of users he follows called `follows`. This is a map so we can add(follow) and remove(unfollow) followees in `O(1)` time.
- `Post` type. Each post has
    - tweetID
    - timestamp
- `PostTweet`.
    - Fetch `userID` from `m`
    - Add the new `Post`(tweetID and timestamp) to the list of user’s posts.
- `Follow`.
    - Fetch the user `followerId` from `m`
    - Add `followeeId` to the user’s `follows` array
- `Unfollow`.
    - Fetch the user `followerId` from `m`
    - Remove `followeeId` to the user’s `follows` array


`GetNewsFeed` is the most complex function out of all, so we’ll dedicate special atention to it, using two approaches.

## First approach - brute force

- `GetNewsFeed`. When a user wants to see his feed, the system:
    - Fetches all the users that this user follows using `follows` and `m`
    - Fetch the last 10 posts of each follow
    - Combine the posts of the user + posts of all the users follows.
    - Sort them
    - Return the first 10

This implementation has a time complexity of $O(n \log n)$ because of the posts's sorting.

## Second approach(most optimal) - min-heap

The brute force solution is actually accepted by Leetcode with a runtime of 4ms. However, `GetNewsFeed` has a time complexity of $O(n \log n)$ which can be optimized to $O(n)$. Let’s see how.

The brute force `GetNewsFeed` , gathers the 10 most recent posts of **all** the followees and sorts them *again* in relation to each other. The function iterates over `n` followes and slices the last 10 items of each. Since a slice is $O(1)$, this part takes $O(N)$. The sorting is $O(\log N)$ so the combined complexity is $O(n \log n)$.

The function can be optimized because every user post is already sorted, so doing another sort is kind of unnecessary.

The better approach is using a min heap with K-largest items. Here, “largest” is evaluated in terms of timestamp.

1. Gather the 10 most recent posts of all the followees and the user’s.
2. Create a min-heap of size `k`, where k = 10. The min-heap will keep the k largest items(posts with the biggest/most recent timestamp). During the iteration process, when we get a post with a larger timestamp than the current root of the min-heap(post with the smallest timestamp among the current k largest ones), we pop the root of the min-heap and add this new post.
3. Return all the posts of the min heap in reversed order, so that the most recent ones appear first.

## Complexity analysis

The complexity of `GetNewsFeed` is determined by the amount of followees a user has.

Let `N` be the number of people the user follows.

1. **Outer Loop:** Runs `N` times (once for each followee).
2. **Inner Loop:** Runs a maximum of 10 times per followee.
    1. Note that even though this is a nested loop, complexity is still $O(1)$ because the amount of times this loop will execute doesn’t increase with the input but is always 10. In other words, **is constant**.
3. **Heap Operations:** Inside that inner loop, you do heap inserts/removes. Because the heap is capped at size 10, these operations take $O(\log10)$ time. Since 10 is a constant, $O(\log10)$ simplifies to constant time: $O(1)$.

When you multiply it all together:

- `N (followees) × 10 (posts) × O(1) (heap ops) = O(10×N)`

In Big O notation, we drop constants, leaving us with a final time complexity of **`O(N)`**.