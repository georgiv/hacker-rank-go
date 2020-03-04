package stacks_and_queues

import (
	"strings"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	var tests = []struct {
		s        string
		expected string
	}{
		{"{[()]}", "YES"},
		{"{[(])}", "NO"},
		{"{{[[(())]]}}", "YES"},
		{"[{", "NO"},
	}

	for _, test := range tests {
		if result := isBalanced(test.s); result != test.expected {
			t.Errorf("isBalanced(%s) = %s, expected = %s", test.s, result, test.expected)
		}
	}
}

func TestIntQueue(t *testing.T) {
	var tests = []struct {
		queries  [][]int
		expected []int
	}{
		{[][]int{{1, 42}, {2}, {1, 14}, {3}, {1, 28},
			{3}, {1, 60}, {1, 78}, {2}, {2}},
			[]int{14, 14}},
		{[][]int{{1, 1}, {2}, {1, 2}, {1, 3}, {1, 4},
			{1, 5}, {2}, {1, 6}, {1, 7}, {1, 8},
			{1, 9}, {3}, {1, 10}, {1, 11}, {1, 12},
			{1, 13}, {1, 14}, {3}, {1, 15}, {3},
			{1, 16}, {1, 17}, {1, 18}, {1, 19}, {3},
			{1, 20}, {3}, {3}, {2}, {1, 21},
			{2}, {1, 22}, {1, 23}, {1, 24}, {1, 25},
			{2}, {3}},
			[]int{3, 3, 3, 3, 3, 3, 6}},
	}

	for _, test := range tests {
		result := []int{}

		queue := NewIntQueueStackImpl()

		for _, q := range test.queries {
			switch q[0] {
			case 1:
				queue.Enqueue(q[1])
			case 2:
				queue.Dequeue()
			case 3:
				e, _ := queue.Peek()
				result = append(result, e)
			}
		}

		if !arrayIntDeepEqual(result, test.expected) {
			t.Errorf("Inconsistent IntQueue\nqueries: %v\nresult: %v\nexpected: %v\n", test.queries, result, test.expected)
		}
	}
}

func TestLargestRectangle(t *testing.T) {
	var tests = []struct {
		heights  []int32
		expected int64
	}{
		{[]int32{3, 2, 3}, 6},
		{[]int32{1, 2, 3, 4, 5}, 9},
	}

	for _, test := range tests {
		if result := largestRectangle(test.heights); result != test.expected {
			t.Errorf("largestRectangle(%v) = %d, expected = %d", test.heights, result, test.expected)
		}
	}
}

func TestRiddle(t *testing.T) {
	var tests = []struct {
		arr      []int64
		expected []int64
	}{
		{[]int64{6, 3, 5, 1, 12}, []int64{12, 3, 3, 1, 1}},
		{[]int64{2, 6, 1, 12}, []int64{12, 2, 1, 1}},
		{[]int64{1, 2, 3, 5, 1, 13, 3}, []int64{13, 3, 2, 1, 1, 1, 1}},
		{[]int64{3, 5, 4, 7, 6, 2}, []int64{7, 6, 4, 4, 3, 2}},
		{[]int64{789168277, 694294362, 532144299, 20472621, 316665904,
			59654039, 685958445, 925819184, 371690486, 285650353,
			522515445, 624800694, 396417773, 467681822, 964079876,
			355847868, 424895284, 50621903, 728094833, 535436067,
			221600465, 832169804, 641711594, 518285605, 235027997,
			904664230, 223080251, 337085579, 5125280, 448775176,
			831453463, 550142629, 822686012, 555190916, 911857735,
			144603739, 751265137, 274554418, 450666269, 984349810,
			716998518, 949717950, 313190920, 600769443, 140712186,
			218387168, 416515873, 194487510, 149671312, 241556542,
			575727819, 873823206},
			[]int64{984349810, 716998518, 716998518, 550142629, 550142629,
				448775176, 355847868, 285650353, 285650353, 285650353,
				285650353, 144603739, 144603739, 144603739, 144603739,
				140712186, 140712186, 140712186, 140712186, 140712186,
				140712186, 140712186, 140712186, 50621903, 20472621,
				20472621, 20472621, 20472621, 5125280, 5125280,
				5125280, 5125280, 5125280, 5125280, 5125280,
				5125280, 5125280, 5125280, 5125280, 5125280,
				5125280, 5125280, 5125280, 5125280, 5125280,
				5125280, 5125280, 5125280, 5125280, 5125280,
				5125280, 5125280}},
	}

	for _, test := range tests {
		if result := riddle(test.arr); !arrayInt64DeepEqual(result, test.expected) {
			t.Errorf("riddle(%v) = %v, expected = %v", test.arr, result, test.expected)
		}
	}
}

func arrayIntDeepEqual(a, b []int) bool {
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

func arrayInt64DeepEqual(a, b []int64) bool {
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

func TestMinimumMoves(t *testing.T) {
	var tests = []struct {
		grid     []string
		startX   int32
		startY   int32
		goalX    int32
		goalY    int32
		expected int32
	}{
		{[]string{
			"...",
			".X.",
			"...",
		},
			0, 0, 1, 2, 2},
		{[]string{
			".X.",
			".X.",
			"...",
		},
			0, 0, 0, 2, 3},
		{[]string{
			".X..XX...X",
			"X.........",
			".X.......X",
			"..........",
			"........X.",
			".X...XXX..",
			".....X..XX",
			".....X.X..",
			"..........",
			".....X..XX",
		}, 9, 1, 9, 6, 3},
		{[]string{
			"...X......XX.X...........XX....X.XX.....",
			".X..............X...XX..X...X........X.X",
			"......X....X....X.........X...........X.",
			".X.X.X..X........X.....X.X...X.....X..X.",
			"....X.X.X...X..........X..........X.....",
			"..X......X....X....X...X....X.X.X....XX.",
			"...X....X.......X..XXX.......X.X.....X..",
			"...X.X.........X.X....X...X.........X...",
			"XXXX..X......X.XX......X.X......XX.X..XX",
			".X........X....X.X......X..X....XX....X.",
			"...X.X..X.X.....X...X....X..X....X......",
			"....XX.X.....X.XX.X...X.X.....X.X.......",
			".X.X.X..............X.....XX..X.........",
			"..X...............X......X........XX...X",
			".X......X...X.XXXX.....XX...........X..X",
			"...X....XX....X...XX.X..X..X..X.....X..X",
			"...X...X.X.....X.....X.....XXXX.........",
			"X.....XX..X.............X.XX.X.XXX......",
			".....X.X..X..........X.X..X...X.X......X",
			"...X.....X..X.............X......X..X..X",
			"........X.....................X....X.X..",
			"..........X.....XXX...XX.X..............",
			"........X..X..........X.XXXX..X.........",
			"..X..X...X.......XX...X.....X...XXX..X..",
			".X.......X..............X....X...X....X.",
			".................X.....X......X.....X...",
			".......X.X.XX..X.XXX.X.....X..........X.",
			"X..X......X..............X..X.X.......X.",
			"X........X.....X.....X....XX.......XX...",
			"X.....X.................X.....X..X...XXX",
			"XXX..X..X.X.XX..X....X.....XXX..X......X",
			"..........X.....X.....XX................",
			"..X.........X..X.........X...X.....X....",
			".X.X....X...XX....X...............X.....",
			".X....X....XX.XX........X..X............",
			"X...X.X................XX......X..X.....",
			"..X.X.......X.X..X.....XX.........X..X..",
			"........................X..X.XX..X......",
			"........X..X.X.....X.....X......X.......",
			".X..X....X.X......XX....................",
		}, 34, 28, 16, 8, 9},
	}

	for _, test := range tests {
		if result := minimumMoves(test.grid, test.startX, test.startY, test.goalX, test.goalY); result != test.expected {
			t.Errorf("minimumMoves(\n%v, (%d, %d), (%d, %d)) = %d, expected = %d",
				strings.Join(test.grid, "\n"),
				test.startX, test.startY, test.goalX, test.goalY, result, test.expected)
		}
	}
}

func TestPoisonousPlants(t *testing.T) {
	var tests = []struct {
		plants   []int32
		expected int32
	}{
		{[]int32{3, 6, 2, 7, 5}, 2},
		{[]int32{6, 5, 8, 4, 7, 10, 9}, 2},
		{[]int32{3, 2, 5, 4}, 2},
		{[]int32{4, 3, 7, 5, 6, 4, 2}, 3},
		{[]int32{4, 5, 3, 2, 1, 6}, 1},
	}

	for _, test := range tests {
		if result := poisonousPlants(test.plants); result != test.expected {
			t.Errorf("poisonousPlants(%v) = %d, expected = %d", test.plants, result, test.expected)
		}
	}
}
