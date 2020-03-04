package arrays

import "testing"

func TestHourglassSum(t *testing.T) {
	var tests = []struct {
		arr      [][]int32
		expected int32
	}{
		{[][]int32{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
		}, 7},
		{[][]int32{
			{-9, -9, -9, 1, 1, 1},
			{0, -9, 0, 4, 3, 2},
			{-9, -9, -9, 1, 2, 3},
			{0, 0, 8, 6, 6, 0},
			{0, 0, 0, -2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		}, 28},
		{[][]int32{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 0, 2, 4, 4, 0},
			{0, 0, 0, 2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		}, 19},
	}

	for _, test := range tests {
		if result := hourglassSum(test.arr); result != test.expected {
			t.Errorf("hourglassSum(%v) = %d, expected = %d", test.arr, result, test.expected)
		}
	}
}

func TestRotLeft(t *testing.T) {
	var tests = []struct {
		arr      []int32
		offset   int32
		expected []int32
	}{
		{[]int32{1, 2, 3, 4, 5}, 2, []int32{3, 4, 5, 1, 2}},
		{[]int32{1, 2, 3, 4, 5}, 4, []int32{5, 1, 2, 3, 4}},
		{[]int32{1, 2, 3, 4, 5}, 5, []int32{1, 2, 3, 4, 5}},
		{[]int32{1, 2, 3, 4, 5}, 8, []int32{4, 5, 1, 2, 3}},
	}

	for _, test := range tests {
		if result := rotLeft(test.arr, test.offset); !arrayDeepEqual(result, test.expected) {
			t.Errorf("rotLeft(%v, %d) = %v, expected = %v", test.arr, test.offset, result, test.expected)
		}
	}
}

func TestMinimumBribes(t *testing.T) {
	var tests = []struct {
		queue    []int32
		expected int32
	}{
		{[]int32{2, 1, 5, 3, 4}, 3},
		{[]int32{2, 5, 1, 3, 4}, -1},
		{[]int32{5, 1, 2, 3, 7, 8, 6, 4}, -1},
		{[]int32{1, 2, 5, 3, 7, 8, 6, 4}, 7},
	}

	for _, test := range tests {
		if result := minimumBribes(test.queue); result != test.expected {
			t.Errorf("minimumBribes(%v) = %d, expected = %d", test.queue, result, test.expected)
		}
		minimumBribes(test.queue)
	}
}

func TestMinimumSwaps(t *testing.T) {
	var tests = []struct {
		arr      []int32
		expected int32
	}{
		{[]int32{7, 1, 3, 2, 4, 5, 6}, 5},
		{[]int32{4, 3, 1, 2}, 3},
		{[]int32{2, 3, 4, 1, 5}, 3},
		{[]int32{1, 3, 5, 2, 4, 6, 7}, 3},
	}

	for _, test := range tests {
		if result := minimumSwaps(test.arr); result != test.expected {
			t.Errorf("minimumSwaps(%v) = %d, expected = %d", test.arr, result, test.expected)
		}
	}
}

func TestArrayManipulation(t *testing.T) {
	var tests = []struct {
		n        int32
		queries  [][]int32
		expected int64
	}{
		{10, [][]int32{
			{1, 5, 3},
			{4, 8, 7},
			{6, 9, 1}}, 10},
		{5, [][]int32{
			{1, 2, 100},
			{2, 5, 100},
			{3, 4, 100}}, 200},
	}

	for _, test := range tests {
		if result := arrayManipulation(test.n, test.queries); result != test.expected {
			t.Errorf("arrayManipulation(%d, %v) = %d, expected = %d", test.n, test.queries, result, test.expected)
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
