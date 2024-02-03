package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSecondsToLongTime(t *testing.T) {
	c := require.New(t)

	c.Equal("1 year, 234 days, 18 hours, 43 minutes and 8 seconds", secondsToLongTime(51820988))
	c.Equal("187 days, 1 hour, 38 minutes and 12 seconds", secondsToLongTime(16162692))
	c.Equal("3 minutes", secondsToLongTime(180))
	c.Equal("1 hour, 21 minutes and 36 seconds", secondsToLongTime(4896))
	c.Equal("5 seconds", secondsToLongTime(5))
	c.Equal("1 day, 1 hour, 1 minute and 1 second", secondsToLongTime(90061))
}
