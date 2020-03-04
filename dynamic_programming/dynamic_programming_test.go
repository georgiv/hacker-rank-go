package dynamic_programming

import "testing"

func TestMaxSubsetSum(t *testing.T) {
	var tests = []struct {
		ns       []int32
		expected int32
	}{
		{[]int32{-2, 1, 3, -4, 5}, 8},
		{[]int32{3, 7, 4, 6, 5}, 13},
		{[]int32{2, 1, 5, 8, 4}, 11},
		{[]int32{3, 5, -7, 8, 10}, 15},
	}

	for _, test := range tests {
		if result := maxSubsetSum(test.ns); result != test.expected {
			t.Errorf("maxSubsetSum(%v) = %d, expected = %d", test.ns, result, test.expected)
		}
	}
}

func TestAbbreviation(t *testing.T) {
	var tests = []struct {
		a        string
		b        string
		expected string
	}{
		{"AbcDE", "ABDE", "YES"},
		{"AbcDE", "AFDE", "NO"},
		{"daBcd", "ABC", "YES"},
		{"Pi", "P", "YES"},
		{"AfPZN", "APZNC", "NO"},
		{"LDJAN", "LJJM", "NO"},
		{"sYOCa", "YOCN", "NO"},
		{"bBccC", "BBC", "YES"},
	}

	for _, test := range tests {
		if result := abbreviation(test.a, test.b); result != test.expected {
			t.Errorf("abbreviation(%s, %s) = %s, expected = %s", test.a, test.b, result, test.expected)
		}
	}
}

func TestCandies(t *testing.T) {
	var tests = []struct {
		students int32
		ratings  []int32
		expected int64
	}{
		{6, []int32{4, 6, 4, 5, 6, 2}, 10},
		{3, []int32{1, 2, 2}, 4},
		{10, []int32{2, 4, 2, 6, 1, 7, 8, 9, 2, 1}, 19},
		{8, []int32{2, 4, 3, 5, 2, 6, 4, 5}, 12},
		{6, []int32{1, 2, 4, 3, 2, 1}, 8},
	}

	for _, test := range tests {
		if result := candies(test.students, test.ratings); result != test.expected {
			t.Errorf("candies(%d, %v) = %d, expected = %d", test.students, test.ratings, result, test.expected)
		}
	}
}

func TestDecibinaryNumbers(t *testing.T) {
	var tests = []struct {
		x        int64
		expected int64
	}{
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 10},
		{10, 100},
		{8, 12},
		{23, 23},
		{19, 102},
		{16, 14},
		{26, 111},
		{7, 4},
		{6, 11},
		{19, 102},
		{25, 103},
		{6, 11},
		{8, 12},
		{20, 110},
		{10, 100},
		{27, 8},
		{24, 31},
		{30, 32},
		{11, 5},
	}

	for _, test := range tests {
		if result := decibinaryNumbers(test.x); result != test.expected {
			t.Errorf("decibinaryNumbers(%d) = %d, expected = %d", test.x, result, test.expected)
		}
	}
}
