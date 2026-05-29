package main

import (
	"sort"
	"time"
)

type Twitter struct {
	m map[int]User
}

type User struct {
	posts   []Post
	follows map[int]bool
}

type Post struct {
	tweetID   int
	timestamp int64
}

func Constructor() Twitter {
	return Twitter{
		m: make(map[int]User),
	}
}

func NewUser() User {
	return User{
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

	this.m[userId] = user
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	tweets := make([]int, 0)
	user, _ := this.m[userId]
	allPosts := user.posts

	for followee, _ := range user.follows {
		if len(this.m[followee].posts) < 10 {
			allPosts = append(allPosts, this.m[followee].posts...)
			continue
		}
		allPosts = append(allPosts, this.m[followee].posts[len(this.m[followee].posts)-10:]...)
	}

	sort.Slice(allPosts, func(i, j int) bool {
		return allPosts[i].timestamp > allPosts[j].timestamp
	})

	for i := 0; i < 10 && i < len(allPosts); i++ {
		tweets = append(tweets, allPosts[i].tweetID)
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
