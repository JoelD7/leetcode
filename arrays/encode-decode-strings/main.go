package main

import (
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteString("#")
		sb.WriteString(str)
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}

		length, err := strconv.Atoi(encoded[i:j])
		if err != nil {
			break
		}

		start := j + 1 //one place after the #
		end := start + length
		res = append(res, encoded[start:end])
		i = end
	}

	return res
}

func Constructor() *Solution {
	return &Solution{}
}
