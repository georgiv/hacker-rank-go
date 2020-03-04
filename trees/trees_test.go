package trees

import "testing"

func TestHeight(t *testing.T) {
	var tests = []struct {
		nodes    []int32
		expected int32
	}{
		{[]int32{1}, 0},
		{[]int32{1, 2, 3}, 2},
		{[]int32{4, 2, 6, 1, 3, 5, 7}, 2},
		{[]int32{3, 2, 5, 1, 4, 6, 7}, 3},
	}

	for _, test := range tests {
		var root *BinaryTreeNode
		for _, n := range test.nodes {
			root = InsertInBST(root, n)
		}

		if result := height(root); result != test.expected {
			t.Errorf("height(%v) = %d, expected = %d", test.nodes, result, test.expected)
		}
	}
}

func TestLca(t *testing.T) {
	var tests = []struct {
		nodes    []int32
		x        int32
		y        int32
		expected int32
	}{
		{[]int32{3, 2, 4, 1, 5, 6}, 4, 6, 4},
		{[]int32{4, 2, 7, 1, 3, 6}, 2, 7, 4},
	}

	for _, test := range tests {
		var root *BinaryTreeNode
		for _, n := range test.nodes {
			root = InsertInBST(root, n)
		}

		if result := lca(root, test.x, test.y); result != test.expected {
			t.Errorf("lca(%v, %d, %d) = %v, expected = %d", test.nodes, test.x, test.y, result, test.expected)
		}
	}
}

func TestCheckBST(t *testing.T) {
	var tests = []struct {
		root     *BinaryTreeNode
		expected bool
	}{
		{&BinaryTreeNode{3,
			&BinaryTreeNode{2,
				&BinaryTreeNode{1,
					nil,
					nil},
				nil},
			&BinaryTreeNode{4,
				&BinaryTreeNode{5,
					nil,
					nil},
				&BinaryTreeNode{6,
					nil,
					nil}}}, false},
		{&BinaryTreeNode{3,
			&BinaryTreeNode{2,
				&BinaryTreeNode{1,
					nil,
					nil},
				nil},
			&BinaryTreeNode{5,
				&BinaryTreeNode{6,
					nil,
					nil},
				&BinaryTreeNode{1,
					nil,
					nil}}}, false},
		{&BinaryTreeNode{4,
			&BinaryTreeNode{2,
				&BinaryTreeNode{1,
					nil,
					nil},
				&BinaryTreeNode{3,
					nil,
					nil}},
			&BinaryTreeNode{6,
				&BinaryTreeNode{5,
					nil,
					nil},
				&BinaryTreeNode{7,
					nil,
					nil}}}, true},
	}

	for _, test := range tests {
		if result := checkBST(test.root); result != test.expected {
			t.Errorf("checkBST([TODO: print tree]) = %t, expected = %t", result, test.expected)
		}
	}
}

func TestDecode(t *testing.T) {
	var tests = []struct {
		secret   string
		root     *HuffmanTrieNode
		expected string
	}{
		{"01111001100011010111100",
			&HuffmanTrieNode{11, 0,
				&HuffmanTrieNode{5, 'A',
					nil,
					nil},
				&HuffmanTrieNode{6, 0,
					&HuffmanTrieNode{2, 'R',
						nil,
						nil},
					&HuffmanTrieNode{4, 0,
						&HuffmanTrieNode{2, 0,
							&HuffmanTrieNode{1, 'C',
								nil,
								nil},
							&HuffmanTrieNode{1, 'D',
								nil,
								nil}},
						&HuffmanTrieNode{2, 'B',
							nil,
							nil}}}},
			"ABRACADABRA"},
		{"1001011",
			&HuffmanTrieNode{5, 0,
				&HuffmanTrieNode{2, 0,
					&HuffmanTrieNode{1, 'B',
						nil,
						nil},
					&HuffmanTrieNode{1, 'C',
						nil,
						nil}},
				&HuffmanTrieNode{3, 'A',
					nil,
					nil}},
			"ABACA"},
	}

	for _, test := range tests {
		if result := decode(test.secret, test.root); result != test.expected {
			t.Errorf("decode(%s, [TODO: print tree]) = %s, expected = %s", test.secret, result, test.expected)
		}
	}
}

func TestBalancedForest(t *testing.T) {
	var tests = []struct {
		nodes    []int32
		edges    [][]int32
		expected int64
	}{
		{[]int32{15, 12, 8, 14, 13}, [][]int32{{1, 2}, {1, 3}, {1, 4}, {4, 5}}, 19},
		{[]int32{1, 2, 2, 1, 1}, [][]int32{{1, 2}, {1, 3}, {3, 5}, {1, 4}}, 2},
		{[]int32{1, 3, 5}, [][]int32{{1, 3}, {1, 2}}, -1},
		{[]int32{12, 10, 8, 12, 14, 12}, [][]int32{{1, 2}, {1, 3}, {1, 4}, {2, 5}, {4, 6}}, 4},
		{[]int32{1}, [][]int32{}, -1},
		{[]int32{2, 3, 3, 4}, [][]int32{{1, 2}, {1, 4}, {2, 3}}, 6},
		{[]int32{7, 7, 21, 3, 1, 2}, [][]int32{{1, 2}, {3, 1}, {2, 4}, {5, 2}, {2, 6}}, -1},
		{[]int32{7, 7, 4, 1, 1, 1}, [][]int32{{1, 2}, {3, 1}, {2, 4}, {2, 5}, {2, 6}}, -1},
		{[]int32{1, 1, 1, 18, 10, 11, 5, 6}, [][]int32{{1, 2}, {1, 4}, {2, 3}, {1, 8}, {8, 7}, {7, 6}, {5, 7}}, 10},
		{[]int32{12, 7, 11, 17, 20, 10}, [][]int32{{1, 2}, {2, 3}, {4, 5}, {6, 5}, {1, 4}}, 13},
		{[]int32{10, 4, 1, 5, 6, 4, 5, 5}, [][]int32{{1, 2}, {2, 3}, {1, 4}, {5, 4}, {5, 6}, {7, 8}, {7, 5}}, 5},
		{[]int32{100, 100, 99, 99, 98, 98}, [][]int32{{1, 3}, {3, 5}, {1, 2}, {2, 4}, {4, 6}}, 297},
	}

	for _, test := range tests {
		if result := balancedForest(test.nodes, test.edges); result != test.expected {
			t.Errorf("balancedForest(%v, %v) = %d, expected = %d", test.nodes, test.edges, result, test.expected)
		}
	}
}
