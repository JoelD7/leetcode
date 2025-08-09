package group_anagrams

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	sortResults := func(result [][]string) {
		// Sort strings within each group
		for _, group := range result {
			sort.Strings(group)
		}
		// Sort the groups themselves based on the first element
		sort.Slice(result, func(i, j int) bool {
			if len(result[i]) == 0 || len(result[j]) == 0 {
				return len(result[i]) < len(result[j])
			}
			return result[i][0] < result[j][0]
		})
	}

	t.Run("Example 1: Multiple groups", func(t *testing.T) {
		input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
		want := [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}
		got := groupAnagrams(input)

		sortResults(got)
		sortResults(want)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("groupAnagrams(%v) = %v; want %v", input, got, want)
		}
	})

	t.Run("Example 2: Single empty string", func(t *testing.T) {
		input := []string{""}
		want := [][]string{{""}}
		got := groupAnagrams(input)

		sortResults(got)
		sortResults(want)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("groupAnagrams(%v) = %v; want %v", input, got, want)
		}
	})

	t.Run("Example 3: Single character string", func(t *testing.T) {
		input := []string{"a"}
		want := [][]string{{"a"}}
		got := groupAnagrams(input)

		sortResults(got)
		sortResults(want)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("groupAnagrams(%v) = %v; want %v", input, got, want)
		}
	})

	t.Run("No input strings", func(t *testing.T) {
		input := []string{}
		want := [][]string{}
		got := groupAnagrams(input)

		// No need to sort empty slices, but we'll do it for consistency
		sortResults(got)
		sortResults(want)

		assert.Equal(t, want, got)
	})

	t.Run("Strings with different cases (should not be anagrams)", func(t *testing.T) {
		input := []string{"Abc", "bca"}
		want := [][]string{{"Abc"}, {"bca"}}
		got := groupAnagrams(input)

		sortResults(got)
		sortResults(want)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("groupAnagrams(%v) = %v; want %v", input, got, want)
		}
	})

	t.Run("All unique strings", func(t *testing.T) {
		input := []string{"hello", "world", "go"}
		want := [][]string{{"go"}, {"hello"}, {"world"}}
		got := groupAnagrams(input)

		sortResults(got)
		sortResults(want)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("groupAnagrams(%v) = %v; want %v", input, got, want)
		}
	})

	t.Run("[cab, tin, pew, duh, may, ill, buy, bar, max, doc]", func(t *testing.T) {
		input := []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}
		want := [][]string{{"max"}, {"buy"}, {"doc"}, {"may"}, {"ill"}, {"duh"}, {"tin"}, {"bar"}, {"pew"}, {"cab"}}
		got := groupAnagrams(input)

		sortResults(got)
		sortResults(want)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v\n want %v", got, want)
		}
	})
}
