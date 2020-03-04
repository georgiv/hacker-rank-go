package sorting

import (
	"sort"
	"testing"
)

func TestCountSwaps(t *testing.T) {
	var tests = []struct {
		a        []int32
		expected string
	}{
		{[]int32{6, 4, 1}, "Array is sorted in 3 swaps.\nFirst Element: 1\nLast Element: 6"},
		{[]int32{1, 2, 3}, "Array is sorted in 0 swaps.\nFirst Element: 1\nLast Element: 3"},
		{[]int32{3, 2, 1}, "Array is sorted in 3 swaps.\nFirst Element: 1\nLast Element: 3"},
	}

	for _, test := range tests {
		if result := countSwaps(test.a); result != test.expected {
			t.Errorf("countSwaps(%v) = %s, \nexpected = \n%s", test.a, result, test.expected)
		}
	}
}

func TestMaximumToys(t *testing.T) {
	var tests = []struct {
		prices   []int32
		money    int32
		expected int32
	}{
		{[]int32{1, 2, 3, 4}, 7, 3},
		{[]int32{1, 12, 5, 111, 200, 1000, 10}, 50, 4},
	}

	for _, test := range tests {
		if result := maximumToys(test.prices, test.money); result != test.expected {
			t.Errorf("maximumToys(%v, %d) = %d, expected = %d", test.prices, test.money, result, test.expected)
		}
	}
}

func TestCompare(t *testing.T) {
	var tests = []struct {
		ps       []Player
		expected []Player
	}{
		{[]Player{{"Ariel", 4}, {"Sebastian", 3}, {"Anna", 4}},
			[]Player{{"Sebastian", 3}, {"Anna", 4}, {"Ariel", 4}}},
	}

	for _, test := range tests {
		psc := make([]Player, len(test.ps))
		copy(psc, test.ps)
		pbs := PlayersByScore(psc)
		sort.Sort(pbs)
		if !playersDeepEqual(psc, test.expected) {
			t.Errorf("unsorted: %v, result: %v, expected: %v", test.ps, psc, test.expected)
		}
	}
}

func TestAcitvityNotifications(t *testing.T) {
	var tests = []struct {
		expenditure []int32
		d           int32
		expected    int32
	}{
		{[]int32{10, 20, 30, 40, 50}, 3, 1},
		{[]int32{2, 3, 4, 2, 3, 6, 8, 4, 5}, 5, 2},
		{[]int32{1, 2, 3, 4, 4}, 4, 0},
		{[]int32{1, 2, 3, 4, 8}, 4, 1},
	}

	for _, test := range tests {
		if result := activityNotifications(test.expenditure, test.d); result != test.expected {
			t.Errorf("activityNotifications(%v, %d) = %d, expected = %d", test.expenditure, test.d, result, test.expected)
		}
	}
}

func playersDeepEqual(a []Player, b []Player) bool {
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

func TestCountInversions(t *testing.T) {
	var tests = []struct {
		arr      []int32
		expected int64
	}{
		{[]int32{2, 4, 1}, 2},
		{[]int32{1, 1, 1, 2, 2}, 0},
		{[]int32{2, 1, 3, 1, 2}, 4},
		{[]int32{480130, 735329, 810013, 140187, 972418, 357056, 334780, 854684, 478980, 34590, 180238, 567090, 935248, 834854, 514996, 950881, 104525, 874212, 160000, 528785, 770712, 201115, 830644, 535327, 668785, 241043, 946633, 645108, 184097, 251762, 97219, 664226, 503442, 907232, 320764, 992211, 264287, 171895, 846895, 259618, 722836, 543484, 343059, 658083, 378338, 858054, 125316, 999214, 248618, 285315, 527999, 535681, 2781, 358642, 71007, 671565, 116036, 17639, 833024, 300132, 785753, 930243, 480710, 805546, 353826, 317825, 797757, 134464, 6072, 161003, 394081, 728907, 704486, 253492, 386990, 599175, 111545, 28657, 598389, 360162, 313971, 642739, 895842, 833103, 1380, 966849, 21019, 633767, 500839, 854043, 450251, 802943, 784285, 930960, 608489, 654462, 765136, 922597, 788925, 771207, 83599, 699357, 500114, 304436, 952848, 403455, 419963, 580745, 432111, 534703, 457258, 262433}, 3083},
		{[]int32{62935, 82200, 877141, 585771, 619073, 183328, 809452, 189197, 41883, 777611, 360495, 295099, 198393, 308583, 537954, 11054, 515803, 403848}, 77},
		{[]int32{9492052, 241944, 5743396, 5758608, 6053545}, 4},
	}

	for _, test := range tests {
		if result := countInversions(test.arr); result != test.expected {
			t.Errorf("countInversions(%v) = %d, expected = %d", test.arr, result, test.expected)
		}
	}
}
