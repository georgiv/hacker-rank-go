// https://www.hackerrank.com/interview/interview-preparation-kit/trees/challenges
package trees

import (
	"bytes"
	"math"
	"time"
)

type BinaryTreeNode struct {
	data  int32
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func InsertInBST(root *BinaryTreeNode, data int32) *BinaryTreeNode {
	if root == nil {
		return &BinaryTreeNode{data: data}
	}

	var current *BinaryTreeNode

	if data <= root.data {
		current = InsertInBST(root.left, data)
		root.left = current
	} else {
		current = InsertInBST(root.right, data)
		root.right = current
	}

	return root
}

// https://www.hackerrank.com/challenges/tree-height-of-a-binary-tree/problem
func height(root *BinaryTreeNode) int32 {
	if root == nil {
		return -1
	}

	leftHeight := height(root.left)
	rightHeight := height(root.right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

// https://www.hackerrank.com/challenges/binary-search-tree-lowest-common-ancestor/problem
func lca(root *BinaryTreeNode, x, y int32) int32 {
	if x < root.data && y < root.data {
		return lca(root.left, x, y)
	}

	if x > root.data && y > root.data {
		return lca(root.right, x, y)
	}

	return root.data
}

// https://www.hackerrank.com/challenges/ctci-is-binary-search-tree/problem
func checkBST(root *BinaryTreeNode) bool {
	return checkBSTInternal(root, math.MinInt32, math.MaxInt32)
}

func checkBSTInternal(root *BinaryTreeNode, min, max int32) bool {
	if root == nil {
		return true
	}

	if root.data < min || root.data > max {
		return false
	}

	return checkBSTInternal(root.left, min, root.data-1) && checkBSTInternal(root.right, root.data+1, max)
}

// https://www.hackerrank.com/challenges/tree-huffman-decoding/problem
type HuffmanTrieNode struct {
	frequency int32
	data      int32
	left      *HuffmanTrieNode
	right     *HuffmanTrieNode
}

func decode(secret string, root *HuffmanTrieNode) string {
	var result bytes.Buffer

	node := root

	for _, c := range secret {
		if c == '0' {
			node = node.left
		} else {
			node = node.right
		}

		if node.data != 0 {
			result.WriteString(string(node.data))
			node = root
		}
	}

	return result.String()
}

// https://www.hackerrank.com/challenges/balanced-forest/problem
type TreeNode struct {
	data    int32
	sum     int64
	next    []*TreeNode
	enterTS int64
	exitTS  int64
}

func balancedForest(nodes []int32, edges [][]int32) int64 {
	if len(nodes) < 3 {
		return -1
	}

	root := buildTree(nodes, edges)

	sums := make(map[int64][]*TreeNode)
	rootSum := calculateSumsCached(root, nil, sums)

	var result int64 = -1

	q := []*TreeNode{root}

	for {
		if len(q) == 0 {
			break
		}

		n := q[0]
		q = q[1:]

		for _, c := range n.next {
			q = append(q, c)
		}

		if n.sum == rootSum {
			continue
		}

		currentSum := n.sum
		remainderSum := rootSum - currentSum

		if remainderSum == currentSum {
			if result == -1 || currentSum < result {
				result = currentSum
			}
		} else {
			large := currentSum
			small := remainderSum
			largeRoot := n
			excludeRoot := root
			if remainderSum > large {
				large, small = small, large
				largeRoot, excludeRoot = excludeRoot, largeRoot
			}

			if large%2 == 0 && large/2 >= small {
				if ns, ok := sums[large/2]; ok {
					for _, c := range ns {
						if (largeRoot.enterTS < c.enterTS && c.exitTS < largeRoot.exitTS) &&
							!(c.enterTS <= excludeRoot.enterTS && excludeRoot.exitTS <= c.exitTS) &&
							(result == -1 || large/2-small < result) {
							result = large/2 - small
							break
						}
					}
				}

				if ns, ok := sums[large/2+small]; ok {
					for _, c := range ns {
						if (largeRoot.enterTS < c.enterTS && c.exitTS < largeRoot.exitTS) &&
							(c.enterTS <= excludeRoot.enterTS && excludeRoot.exitTS <= c.exitTS) &&
							(result == -1 || large/2-small < result) {
							result = large/2 - small
							break
						}
					}
				}
			}

			if small > large-small {
				if ns, ok := sums[small]; ok {
					for _, c := range ns {
						if (largeRoot.enterTS < c.enterTS && c.exitTS < largeRoot.exitTS) &&
							!(c.enterTS <= excludeRoot.enterTS && excludeRoot.exitTS <= c.exitTS) &&
							(result == -1 || small-(large-small) < result) {
							result = small - (large - small)
							break
						}
					}
				}

				if ns, ok := sums[2*small]; ok {
					for _, c := range ns {
						if (largeRoot.enterTS < c.enterTS && c.exitTS < largeRoot.exitTS) &&
							(c.enterTS <= excludeRoot.enterTS && excludeRoot.exitTS <= c.exitTS) &&
							(result == -1 || small-(large-small) < result) {
							result = small - (large - small)
							break
						}
					}
				}
			}
		}
	}

	return result
}

func buildTree(nodes []int32, edges [][]int32) *TreeNode {
	treeNodes := make([]*TreeNode, len(nodes))

	for i, n := range nodes {
		treeNodes[i] = &TreeNode{data: n}
	}

	for _, e := range edges {
		treeNodes[e[0]-1].next = append(treeNodes[e[0]-1].next, treeNodes[e[1]-1])
		treeNodes[e[1]-1].next = append(treeNodes[e[1]-1].next, treeNodes[e[0]-1])
	}

	root := treeNodes[0]

	establishParentChildRelations(root, nil)

	return root
}

func establishParentChildRelations(node, parent *TreeNode) {
	node.enterTS = time.Now().UnixNano()

	var children []*TreeNode
	for i := 0; i < len(node.next); i++ {
		if node.next[i] == parent {
			continue
		}
		children = append(children, node.next[i])
	}

	node.next = children

	for _, c := range node.next {
		establishParentChildRelations(c, node)
	}

	node.exitTS = time.Now().UnixNano()
}

func calculateSumsCached(root, excludeNode *TreeNode, sums map[int64][]*TreeNode) int64 {
	if root == excludeNode {
		return 0
	}

	result := int64(root.data)

	for _, n := range root.next {
		result += calculateSumsCached(n, excludeNode, sums)
	}

	root.sum = result
	sums[result] = append(sums[result], root)

	return result
}
