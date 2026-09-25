package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodec(t *testing.T) {
	codec := Constructor()

	t.Run("basic strings with standard words", func(t *testing.T) {
		input := []string{"lint", "code", "love", "you"}
		encoded := codec.Encode(input)
		decoded := codec.Decode(encoded)
		assert.Equal(t, input, decoded)
	})

	t.Run("strings containing special characters and spaces", func(t *testing.T) {
		input := []string{"we", "say", ":", "yes", "!@#$%^&*()"}
		encoded := codec.Encode(input)
		decoded := codec.Decode(encoded)
		assert.Equal(t, input, decoded)
	})

	t.Run("slice with empty string", func(t *testing.T) {
		input := []string{""}
		encoded := codec.Encode(input)
		decoded := codec.Decode(encoded)
		assert.Equal(t, input, decoded)
	})

	t.Run("empty string slice", func(t *testing.T) {
		input := []string{}
		encoded := codec.Encode(input)
		decoded := codec.Decode(encoded)
		assert.Equal(t, input, decoded)
	})

	t.Run("slice containing empty strings", func(t *testing.T) {
		input := []string{"", "", ""}
		encoded := codec.Encode(input)
		decoded := codec.Decode(encoded)
		assert.Equal(t, input, decoded)
	})

	t.Run("strings containing length delimiters inside the text", func(t *testing.T) {
		input := []string{"4#code", "1#", "hello#world"}
		encoded := codec.Encode(input)
		decoded := codec.Decode(encoded)
		assert.Equal(t, input, decoded)
	})
}
