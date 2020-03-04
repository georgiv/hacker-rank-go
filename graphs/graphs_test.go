package graphs

import "testing"

func TestRoadsAndLibraries(t *testing.T) {
	var tests = []struct {
		cities   int32
		costLib  int32
		costRoad int32
		roads    [][]int32
		expected int64
	}{
		{3, 2, 1, [][]int32{{1, 2}, {3, 1}, {2, 3}}, 4},
		{6, 2, 5, [][]int32{{1, 3}, {3, 4}, {2, 4}, {1, 2}, {2, 3}, {5, 6}}, 12},
		{5, 6, 1, [][]int32{{1, 2}, {1, 3}, {1, 4}}, 15},
		{6, 2, 3, [][]int32{{1, 2}, {1, 3}, {4, 5}, {4, 6}}, 12},
	}

	for _, test := range tests {
		if result := roadsAndLibraries(test.cities, test.costLib, test.costRoad, test.roads); result != test.expected {
			t.Errorf("roadsAndLibraries(%d, %d, %d, %v) = %d, expected = %d",
				test.cities,
				test.costLib,
				test.costRoad,
				test.roads,
				result,
				test.expected)
		}
	}
}

func TestFindShortest(t *testing.T) {
	var tests = []struct {
		vertices int32
		from     []int32
		to       []int32
		colors   []int64
		color    int32
		expected int32
	}{
		{5, []int32{1, 2, 2, 3}, []int32{2, 3, 4, 5}, []int64{1, 2, 3, 1, 3}, 3, 1},
		{4, []int32{1, 1, 4}, []int32{2, 3, 2}, []int64{1, 2, 1, 1}, 1, 1},
		{4, []int32{1, 1, 4}, []int32{2, 3, 2}, []int64{1, 2, 3, 4}, 2, -1},
		{5, []int32{1, 1, 2, 3}, []int32{2, 3, 4, 5}, []int64{1, 2, 3, 3, 2}, 2, 3},
	}

	for _, test := range tests {
		if result := findShortest(test.vertices, test.from, test.to, test.colors, test.color); result != test.expected {
			t.Errorf("findShortest(%d, %v, %v, %v, %d) = %d, expected = %d",
				test.vertices,
				test.from,
				test.to,
				test.colors,
				test.color,
				result,
				test.expected)
		}
	}
}

func TestShortestReach(t *testing.T) {
	var tests = []struct {
		vertices int32
		edges    [][]int32
		start    int32
		expected []int32
	}{
		{6, [][]int32{{1, 2}, {2, 3}, {3, 4}, {1, 5}}, 1, []int32{6, 12, 18, 6, -1}},
		{4, [][]int32{{1, 2}, {1, 3}}, 1, []int32{6, 6, -1}},
		{3, [][]int32{{2, 3}}, 2, []int32{-1, 6}},
	}

	for _, test := range tests {
		if result := shortestReach(test.vertices, test.edges, test.start); !arrayDeepEqual(result, test.expected) {
			t.Errorf("shortestReach(%d, %v, %d) = %v, expected = %v",
				test.vertices,
				test.edges,
				test.start,
				result,
				test.expected)
		}
	}
}

func TestMaxRegion(t *testing.T) {
	var tests = []struct {
		grid     [][]int32
		expected int32
	}{
		{[][]int32{
			{0, 0, 0, 0, 0, 1},
			{0, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
		},
			1},
		{[][]int32{
			{1, 1, 0, 0, 0, 1},
			{0, 1, 1, 0, 0, 0},
			{0, 0, 0, 1, 0, 0},
		},
			5},
		{[][]int32{
			{1, 1, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 1, 0},
			{1, 0, 0, 0},
		},
			5},
	}

	for _, test := range tests {
		if result := maxRegion(test.grid); result != test.expected {
			t.Errorf("maxRegion(%v) = %d, expected = %d", test.grid, result, test.expected)
		}
	}
}

func TestMinTime(t *testing.T) {
	var tests = []struct {
		roads    [][]int32
		machines []int32
		expected int32
	}{
		{[][]int32{{0, 1, 4}, {0, 4, 2}, {1, 2, 3}, {1, 3, 7}},
			[]int32{2, 3, 4},
			5},
		{[][]int32{{2, 1, 8}, {1, 0, 5}, {2, 4, 5}, {1, 3, 4}},
			[]int32{2, 4, 0},
			10},
		{[][]int32{{0, 3, 3}, {1, 4, 4}, {1, 3, 4}, {0, 2, 5}},
			[]int32{1, 3, 4},
			8},
		{[][]int32{{1, 0, 1}, {2, 0, 2}, {3, 1, 3}, {4, 2, 4}, {5, 3, 5}, {6, 4, 6}},
			[]int32{3, 4},
			1},
		{[][]int32{{9, 78, 35}, {9, 54, 45}, {78, 69, 27}, {9, 55, 9}, {9, 1, 78}, {1, 92, 7}, {55, 42, 57}, {1, 84, 4}, {1, 5, 38}, {92, 8, 75}, {55, 30, 99}, {69, 7, 9}, {1, 81, 45}, {8, 31, 4}, {42, 23, 100}, {78, 95, 3}, {54, 14, 14}, {84, 53, 80}, {92, 32, 8}, {42, 86, 40}, {1, 64, 93}, {78, 60, 65}, {64, 76, 24}, {42, 89, 86}, {7, 28, 48}, {69, 62, 26}, {1, 40, 23}, {78, 38, 29}, {8, 44, 39}, {78, 3, 37}, {54, 26, 17}, {62, 50, 24}, {76, 66, 37}, {30, 51, 75}, {86, 43, 91}, {5, 77, 32}, {64, 91, 11}, {14, 10, 36}, {26, 20, 19}, {9, 52, 50}, {77, 94, 32}, {44, 67, 63}, {64, 15, 61}, {92, 0, 73}, {10, 37, 23}, {89, 2, 37}, {92, 18, 51}, {26, 47, 25}, {30, 87, 15}, {47, 36, 35}, {92, 72, 16}, {28, 75, 93}, {78, 73, 66}, {20, 19, 64}, {73, 57, 1}, {91, 6, 50}, {54, 33, 41}, {78, 11, 38}, {37, 71, 55}, {5, 63, 52}, {10, 46, 22}, {94, 82, 19}, {95, 83, 51}, {57, 90, 10}, {63, 58, 94}, {43, 45, 23}, {72, 68, 62}, {82, 85, 88}, {58, 4, 94}, {82, 41, 62}, {3, 22, 68}, {54, 70, 78}, {31, 74, 27}, {36, 29, 61}, {33, 24, 76}, {40, 35, 61}, {83, 79, 51}, {8, 59, 20}, {45, 34, 26}, {38, 12, 18}, {70, 99, 25}, {40, 80, 81}, {31, 97, 58}, {69, 21, 16}, {83, 13, 22}, {80, 48, 49}, {97, 65, 44}, {74, 17, 1}, {68, 16, 92}, {50, 98, 54}, {
			94, 27, 76}, {81, 61, 67}, {85, 49, 96}, {81, 93, 31}, {22, 25, 67}, {57, 96, 93}, {82, 88, 92}, {86, 56, 80}, {25, 39, 44}},
			[]int32{1, 95, 90, 11, 48, 49, 23, 6, 0, 76, 3, 83, 85, 31, 44, 54, 87, 38, 16, 61, 22, 21, 29},
			610},
	}

	for _, test := range tests {
		if result := minTime(test.roads, test.machines); result != test.expected {
			t.Errorf("minTime(%v, %v) = %d, expected = %d", test.roads, test.machines, result, test.expected)
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
