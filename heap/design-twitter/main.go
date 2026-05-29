package main

import (
	"time"
)

type Twitter struct {
	m map[int]*User
}

type User struct {
	posts   []Post
	follows map[int]bool
}

type Post struct {
	tweetID   int
	timestamp int64
}

type MinHeap struct {
	arr []Post
}

func (h *MinHeap) size() int {
	return len(h.arr)
}

func (h *MinHeap) insert(val Post) {
	h.arr = append(h.arr, val)
	h.heapifyUp(h.size() - 1)
}

func (h *MinHeap) heapifyUp(index int) {
	parentIdx := parent(index)
	for h.arr[parentIdx].timestamp > h.arr[index].timestamp {
		h.swap(parentIdx, index)
		index = parentIdx
		parentIdx = parent(index)
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func leftChild(index int) int {
	return 2*index + 1
}

func rightChild(index int) int {
	return 2*index + 2
}

func (h *MinHeap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *MinHeap) peek() Post {
	return h.arr[0]
}

func (h *MinHeap) remove() Post {
	val := h.arr[0]
	h.arr[0] = h.arr[h.size()-1]
	h.arr = h.arr[:h.size()-1]
	h.heapifyDown(0)
	return val
}

func (h *MinHeap) heapifyDown(index int) {
	left, right := leftChild(index), rightChild(index)
	minChild := 0

	for left < h.size() {
		if left == h.size()-1 || h.arr[left].timestamp < h.arr[right].timestamp {
			minChild = left
		} else {
			minChild = right
		}

		if h.arr[minChild].timestamp < h.arr[index].timestamp {
			h.swap(minChild, index)
			index = minChild
			left, right = leftChild(index), rightChild(index)
		} else {
			break
		}
	}
}

func Constructor() Twitter {
	return Twitter{
		m: make(map[int]*User),
	}
}

func NewUser() *User {
	return &User{
		posts:   make([]Post, 0),
		follows: make(map[int]bool),
	}
}

func NewPost(tweetID int) Post {
	return Post{
		tweetID:   tweetID,
		timestamp: time.Now().UnixNano(),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	user, ok := this.m[userId]
	if !ok {
		user = NewUser()
	}

	post := NewPost(tweetId)
	user.posts = append(user.posts, post)

	if !ok {
		this.m[userId] = user
	}
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	user, _ := this.m[userId]
	if user == nil {
		return []int{}
	}

	minHeap := &MinHeap{}
	k := 10

	for _, post := range user.posts {
		if minHeap.size() < k || post.timestamp > minHeap.peek().timestamp {
			minHeap.insert(post)

			if minHeap.size() > k {
				minHeap.remove()
			}
		}
	}

	var post Post
	var count int
	var followee *User

	//The inner loop will execute a maximum of k times, which is constant. So the complexity of these nested loops
	//amortizes to O(N), where N = followees.
	for followeeID, _ := range user.follows {
		followee = this.m[followeeID]
		if followee == nil {
			continue
		}

		count = 0
		for i := len(this.m[followeeID].posts) - 1; i >= 0 && count < k; i-- {
			post = this.m[followeeID].posts[i]
			count++

			if minHeap.size() < k || post.timestamp > minHeap.peek().timestamp {
				minHeap.insert(post)

				if minHeap.size() > k {
					minHeap.remove()
				}
			}
		}
	}

	tweets := make([]int, minHeap.size())
	for i := minHeap.size() - 1; i >= 0; i-- {
		tweets[i] = minHeap.remove().tweetID
	}

	return tweets
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	user, ok := this.m[followerId]
	if !ok {
		user = NewUser()
	}
	this.m[followerId] = user
	user.follows[followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	user, ok := this.m[followerId]
	if !ok {
		user = NewUser()
	}
	delete(user.follows, followeeId)
}
