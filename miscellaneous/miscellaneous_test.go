package miscellaneous

import "testing"

func TestFlippingBits(t *testing.T) {
	var tests = []struct {
		n        int64
		expected int64
	}{
		{9, 4294967286},
		{2147483647, 2147483648},
		{1, 4294967294},
		{0, 4294967295},
		{4, 4294967291},
		{123456, 4294843839},
		{0, 4294967295},
		{802743475, 3492223820},
		{35601423, 4259365872},
	}

	for _, test := range tests {
		if result := flippingBits(test.n); result != test.expected {
			t.Errorf("flippingBits(%d) = %d, expected = %d", test.n, result, test.expected)
		}
	}
}

func TestPrimality(t *testing.T) {
	var tests = []struct {
		n        int32
		expected string
	}{
		{12, "Not prime"},
		{5, "Prime"},
		{7, "Prime"},
	}

	for _, test := range tests {
		if result := primality(test.n); result != test.expected {
			t.Errorf("primality(%d) = %s, expected = %s", test.n, result, test.expected)
		}
	}
}

func TestMaxCircle(t *testing.T) {
	var tests = []struct {
		queries  [][]int32
		expected []int32
	}{
		{[][]int32{{1, 2}, {3, 4}, {2, 3}}, []int32{2, 2, 4}},
		{[][]int32{{1, 2}, {1, 3}}, []int32{2, 3}},
		{[][]int32{{1000000000, 23}, {11, 3778}, {7, 47}, {11, 1000000000}}, []int32{2, 2, 2, 4}},
		{[][]int32{{1, 2}, {3, 4}, {1, 3}, {5, 7}, {5, 6}, {7, 4}}, []int32{2, 2, 4, 4, 4, 7}},
		{[][]int32{{6, 4}, {5, 9}, {8, 5}, {4, 1}, {1, 5}, {7, 2}, {4, 2}, {7, 6}}, []int32{2, 2, 3, 3, 6, 6, 8, 8}},
	}

	for _, test := range tests {
		if result := maxCircle(test.queries); !arrayDeepEqual(result, test.expected) {
			t.Errorf("maxCircle(%v) = %v, expected = %v", test.queries, result, test.expected)
		}
	}
}

func TestMaxXor(t *testing.T) {
	var tests = []struct {
		ns       []int32
		queries  []int32
		expected []int32
	}{
		{[]int32{3, 7, 15, 10}, []int32{3}, []int32{12}},
		{[]int32{0, 1, 2}, []int32{3, 7, 2}, []int32{3, 7, 3}},
		{[]int32{5, 1, 7, 4, 3}, []int32{2, 0}, []int32{7, 7}},
		{[]int32{1, 3, 5, 7}, []int32{17, 6}, []int32{22, 7}},
	}

	for _, test := range tests {
		if result := maxXor(test.ns, test.queries); !arrayDeepEqual(result, test.expected) {
			t.Errorf("maxXor(%v, %v) = %v, expected = %v", test.ns, test.queries, result, test.expected)
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
