package dictionaries_and_hashmaps

import (
	"fmt"
	"strings"
	"testing"
)

func TestCheckMagazine(t *testing.T) {
	var tests = []struct {
		magazine []string
		note     []string
		expected string
	}{
		{[]string{"give", "me", "one", "grand", "today", "night"}, []string{"give", "one", "grand", "today"}, "Yes"},
		{[]string{"two", "times", "three", "is", "not", "four"}, []string{"two", "times", "two", "is", "four"}, "No"},
		{[]string{"ive", "got", "a", "lovely", "bunch", "of", "coconuts"}, []string{"ive", "got", "some", "coconuts"}, "No"},
	}

	for _, test := range tests {
		if result := checkMagazine(test.magazine, test.note); result != test.expected {
			t.Errorf("checkMagazine(%v, %v) = %s, expected = %s",
				strings.Trim(fmt.Sprint(test.magazine), "[]"),
				strings.Trim(fmt.Sprint(test.note), "[]"),
				result,
				test.expected)
		}
	}
}

func TestTwoStrings(t *testing.T) {
	var tests = []struct {
		s1       string
		s2       string
		expected string
	}{
		{"a", "art", "YES"},
		{"be", "cat", "NO"},
		{"hello", "world", "YES"},
		{"hi", "world", "NO"},
	}

	for _, test := range tests {
		if result := twoStrings(test.s1, test.s2); result != test.expected {
			t.Errorf("twoStrings(%s, %s) = %s, expected = %s", test.s1, test.s2, result, test.expected)
		}
	}
}

func TestSherlockAndAnagrams(t *testing.T) {
	var tests = []struct {
		s        string
		expected int32
	}{
		{"mom", 2},
		{"abba", 4},
		{"abcd", 0},
		{"ifailuhkqq", 3},
		{"kkkk", 10},
		{"cdcd", 5},
	}

	for _, test := range tests {
		if result := sherlockAndAnagrams(test.s); result != test.expected {
			t.Errorf("sherlockAndAnagrams(%s) = %d, expected = %d", test.s, result, test.expected)
		}
	}
}

func TestCountTriplets(t *testing.T) {
	var tests = []struct {
		arr      []int64
		r        int64
		expected int64
	}{
		{[]int64{1, 4, 16, 64}, 4, 2},
		{[]int64{1, 2, 2, 4}, 2, 2},
		{[]int64{1, 3, 9, 9, 27, 81}, 3, 6},
		{[]int64{1, 5, 5, 25, 125}, 5, 4},
		{[]int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 161700},
		{[]int64{1, 2, 1, 2, 4}, 2, 3},
	}

	for _, test := range tests {
		if result := countTriplets(test.arr, test.r); result != test.expected {
			t.Errorf("countTriplets(%v, %d) = %d, expected = %d", test.arr, test.r, result, test.expected)
		}
	}
}

func TestFreqQuery(t *testing.T) {
	var tests = []struct {
		queries  [][]int32
		expected []int32
	}{
		{[][]int32{{1, 1}, {2, 2}, {3, 2}, {1, 1}, {1, 1}, {2, 1}, {3, 2}}, []int32{0, 1}},
		{[][]int32{{1, 5}, {1, 6}, {3, 2}, {1, 10}, {1, 10}, {1, 6}, {2, 5}, {3, 2}}, []int32{0, 1}},
		{[][]int32{{3, 4}, {2, 1003}, {1, 16}, {3, 1}}, []int32{0, 1}},
		{[][]int32{{1, 3}, {2, 3}, {3, 2}, {1, 4}, {1, 5}, {1, 5}, {1, 4}, {3, 2}, {2, 4}, {3, 2}}, []int32{0, 1, 1}},
	}

	for _, test := range tests {
		if result := freqQuery(test.queries); !arrayDeepEqual(result, test.expected) {
			t.Errorf("freqQuery(%v) = %v, expected = %v", test.queries, result, test.expected)
		}
	}
}

func arrayDeepEqual(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
