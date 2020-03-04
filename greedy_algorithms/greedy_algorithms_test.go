package greedy_algorithms

import "testing"

func TestMinimumAbsoluteDifference(t *testing.T) {
	var tests = []struct {
		ns       []int32
		expected int32
	}{
		{[]int32{-2, 2, 4}, 2},
		{[]int32{3, -7, 0}, 3},
		{[]int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75}, 1},
		{[]int32{1, -3, 71, 68, 17}, 3},
	}

	for _, test := range tests {
		if result := minimumAbsoluteDifference(test.ns); result != test.expected {
			t.Errorf("minimumAbsoluteDifference(%v) = %d, expected = %d", test.ns, result, test.expected)
		}
	}
}

func TestLuckBalance(t *testing.T) {
	var tests = []struct {
		k        int32
		contests [][]int32
		expected int32
	}{
		{2, [][]int32{{5, 1}, {1, 1}, {4, 0}}, 10},
		{1, [][]int32{{5, 1}, {1, 1}, {4, 0}}, 8},
		{3, [][]int32{{5, 1}, {2, 1}, {1, 1}, {8, 1}, {10, 0}, {5, 0}}, 29},
	}

	for _, test := range tests {
		if result := luckBalance(test.k, test.contests); result != test.expected {
			t.Errorf("luckBalance(%d, %v) = %d, expected = %d", test.k, test.contests, result, test.expected)
		}
	}
}

func TestGetMinimumCost(t *testing.T) {
	var tests = []struct {
		buyers   int32
		prices   []int32
		expected int32
	}{
		{3, []int32{1, 2, 3, 4}, 11},
		{3, []int32{2, 5, 6}, 13},
		{2, []int32{2, 5, 6}, 15},
		{3, []int32{1, 3, 5, 7, 9}, 29},
	}

	for _, test := range tests {
		if result := getMinimumCost(test.buyers, test.prices); result != test.expected {
			t.Errorf("getMinimumCost(%d, %v) = %d, expected = %d", test.buyers, test.prices, result, test.expected)
		}
	}
}

func TestMaxMin(t *testing.T) {
	var tests = []struct {
		k        int32
		ns       []int32
		expected int32
	}{
		{2, []int32{1, 4, 7, 2}, 1},
		{3, []int32{10, 100, 300, 200, 1000, 20, 30}, 20},
		{4, []int32{1, 2, 3, 4, 10, 20, 30, 40, 100, 200}, 3},
		{2, []int32{1, 2, 1, 2, 1}, 0},
		{3, []int32{100, 200, 300, 350, 400, 401, 402}, 2},
	}

	for _, test := range tests {
		if result := maxMin(test.k, test.ns); result != test.expected {
			t.Errorf("maxMin(%d, %v) = %d, expected = %d", test.k, test.ns, result, test.expected)
		}
	}
}

func TestReverseShuffleMerge(t *testing.T) {
	var tests = []struct {
		s        string
		expected string
	}{
		{"abab", "ab"},
		{"eggegg", "egg"},
		{"abcdefgabcdefg", "agfedcb"},
		{"aeiouuoiea", "aeiou"},
	}

	for _, test := range tests {
		if result := reverseShuffleMerge(test.s); result != test.expected {
			t.Errorf("reverseShuffleMerge(%s) = %s, expected = %s", test.s, result, test.expected)
		}
	}
}
