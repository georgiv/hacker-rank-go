package warmup

import "testing"

func TestSockMerchant(t *testing.T) {
	var tests = []struct {
		n        int32
		arr      []int32
		expected int32
	}{
		{7, []int32{1, 2, 1, 2, 1, 3, 2}, 2},
		{9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}, 3},
	}

	for _, test := range tests {
		if result := sockMerchant(test.n, test.arr); result != test.expected {
			t.Errorf("sockMerchant(%d, %v) = %d, expected = %v", test.n, test.arr, result, test.expected)
		}
	}
}

func TestCountingValleys(t *testing.T) {
	var tests = []struct {
		n        int32
		path     string
		expected int32
	}{
		{8, "DDUUUUDDD", 1},
		{8, "DDUUUUDDD", 1},
		{8, "UDDDUDUU", 1},
	}

	for _, test := range tests {
		if result := countingValleys(test.n, test.path); result != test.expected {
			t.Errorf("countingValleys(%d, %s) = %d, expected = %d", test.n, test.path, result, test.expected)
		}
	}
}

func TestJumpingOnClouds(t *testing.T) {
	var tests = []struct {
		clouds   []int32
		expected int32
	}{
		{[]int32{0, 1, 0, 0, 0, 1, 0}, 3},
		{[]int32{0, 0, 1, 0, 0, 1, 0}, 4},
		{[]int32{0, 0, 0, 0, 1, 0}, 3},
	}

	for _, test := range tests {
		if result := jumpingOnClouds(test.clouds); result != test.expected {
			t.Errorf("jumpingOnClouds(%v) = %d, expected = %d", test.clouds, result, test.expected)
		}
	}
}

func TestRepeatedString(t *testing.T) {
	var tests = []struct {
		s        string
		n        int64
		expected int64
	}{
		{"abcac", 10, 4},
		{"aba", 10, 7},
		{"a", 1000000000000, 1000000000000},
	}

	for _, test := range tests {
		if result := repeatedString(test.s, test.n); result != test.expected {
			t.Errorf("repeatedString(%s, %d) = %d, expected = %d", test.s, test.n, result, test.expected)
		}
	}
}
