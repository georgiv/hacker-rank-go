// https://www.hackerrank.com/interview/interview-preparation-kit/miscellaneous/challenges
package miscellaneous

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

// https://www.hackerrank.com/challenges/flipping-bits/problem
func flippingBits(n int64) int64 {
	return int64(^uint32(n))
}

// https://www.hackerrank.com/challenges/ctci-big-o/problem
func primality(n int32) string {
	if n <= 1 {
		return "Not prime"
	}

	sqrtN := int(math.Sqrt(float64(n))) + 1

	for i := 2; i < sqrtN; i++ {
		if int(n)%i == 0 {
			return "Not prime"
		}
	}

	return "Prime"
}

// https://www.hackerrank.com/challenges/friend-circle-queries/problem
func maxCircle(queries [][]int32) []int32 {
	aliases := make(map[int32]int32)
	var nextAlias int32

	for _, q := range queries {
		if _, ok := aliases[q[0]]; !ok {
			nextAlias++
			aliases[q[0]] = nextAlias
		}
		if _, ok := aliases[q[1]]; !ok {
			nextAlias++
			aliases[q[1]] = nextAlias
		}
	}

	ufArr := make([]int32, nextAlias)

	var result []int32

	circles := make(map[int32]int32)
	var max int32

	for _, q := range queries {
		if !find(ufArr, aliases[q[0]], aliases[q[1]]) {
			r1 := root(ufArr, aliases[q[0]])
			r2 := root(ufArr, aliases[q[1]])

			if circles[r1] == 0 {
				circles[r1] = 1
			}
			if circles[r2] == 0 {
				circles[r2] = 1
			}

			union(ufArr, aliases[q[0]], aliases[q[1]])

			unifiedCircles := circles[r1] + circles[r2]
			circles[r1] = unifiedCircles
			circles[r2] = unifiedCircles

			if unifiedCircles > max {
				max = unifiedCircles
			}
		}

		result = append(result, max)
	}

	return result
}

func union(ufArr []int32, x, y int32) {
	ufArr[root(ufArr, x)-1] = root(ufArr, y)
}

func find(ufArr []int32, x, y int32) bool {
	return root(ufArr, x) == root(ufArr, y)
}

func root(ufArr []int32, x int32) int32 {
	for {
		if ufArr[x-1] == 0 {
			break
		}

		x = ufArr[x-1]
	}

	return x
}

// https://www.hackerrank.com/challenges/maximum-xor/problem
type TrieNode struct {
	next map[int32]*TrieNode
}

func maxXor(ns []int32, queries []int32) []int32 {
	trie := TrieNode{make(map[int32]*TrieNode)}

	for _, n := range ns {
		addToTrie(&trie, n)
	}

	var result []int32

	for _, q := range queries {
		result = append(result, q^findPath(&trie, q))
	}

	return result
}

func addToTrie(trie *TrieNode, n int32) {
	bs := fmt.Sprintf("%b", n)

	for i := 0; i < 32-len(bs); i++ {
		if _, ok := trie.next['0']; !ok {
			trie.next['0'] = &TrieNode{make(map[int32]*TrieNode)}
		}
		trie = trie.next['0']
	}

	for _, c := range bs {
		if _, ok := trie.next[c]; !ok {
			trie.next[c] = &TrieNode{make(map[int32]*TrieNode)}
		}
		trie = trie.next[c]
	}
}

func findPath(trie *TrieNode, n int32) int32 {
	var path bytes.Buffer

	bs := fmt.Sprintf("%b", n)

	for i := 0; i < 32-len(bs); i++ {
		if _, ok := trie.next['1']; ok {
			trie = trie.next['1']
			path.WriteString(string('1'))
		} else {
			trie = trie.next['0']
			path.WriteString(string('0'))
		}
	}

	for _, c := range bs {
		node0, ok0 := trie.next['0']
		node1, ok1 := trie.next['1']

		if c == '1' {
			if ok0 {
				trie = node0
				path.WriteString(string('0'))
			} else {
				trie = node1
				path.WriteString(string('1'))
			}
		} else {
			if ok1 {
				trie = node1
				path.WriteString(string('1'))
			} else {
				trie = node0
				path.WriteString(string('0'))
			}
		}
	}

	result, _ := strconv.ParseInt(path.String(), 2, 32)

	return int32(result)
}
