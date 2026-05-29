package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwitter(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		twitter := Constructor()
		twitter.PostTweet(1, 5)
		assert.Equal(t, []int{5}, twitter.GetNewsFeed(1))
		twitter.Follow(1, 2)
		twitter.PostTweet(2, 6)
		assert.Equal(t, []int{6, 5}, twitter.GetNewsFeed(1))
		twitter.Unfollow(1, 2)
		assert.Equal(t, []int{5}, twitter.GetNewsFeed(1))
	})

	t.Run("Test case 1", func(t *testing.T) {
		tw := Constructor()
		tw.PostTweet(1, 1)
		assert.Equal(t, []int{1}, tw.GetNewsFeed(1))
		tw.Follow(2, 1)
		assert.Equal(t, []int{1}, tw.GetNewsFeed(2))
		tw.Unfollow(2, 1)
		assert.Equal(t, []int{}, tw.GetNewsFeed(2))
	})

	t.Run("Test case 2", func(t *testing.T) {
		tw := Constructor()
		assert.Empty(t, tw.GetNewsFeed(1))
	})

	t.Run("Test case 3", func(t *testing.T) {
		tw := Constructor()
		tw.Follow(1, 5)
		assert.Empty(t, tw.GetNewsFeed(1))
	})
}
