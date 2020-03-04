// https://www.hackerrank.com/interview/interview-preparation-kit/search/challenges
package search

import (
	"fmt"
	"sort"
)

type Int32Slice []int32

func (s Int32Slice) Len() int {
	return len(s)
}

func (s Int32Slice) Less(i int, j int) bool {
	return s[i] < s[j]
}

func (s Int32Slice) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

type Int64Slice []int64

func (s Int64Slice) Len() int {
	return len(s)
}

func (s Int64Slice) Less(i int, j int) bool {
	return s[i] < s[j]
}

func (s Int64Slice) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

// https://www.hackerrank.com/challenges/ctci-ice-cream-parlor/problem
func whatFlavors(cost []int32, money int32) (int32, int32) {
	c := make([]int32, len(cost))
	copy(c, cost)
	sort.Sort(Int32Slice(c))

	var cost1 int32
	var cost2 int32

	for i, v := range c {
		cs := c[i+1:]
		cost2SortedIndex := sort.Search(len(cs), func(j int) bool { return cs[j] >= money-v })
		if cost2SortedIndex == len(cs) || cs[cost2SortedIndex] != money-v {
			continue
		}

		cost2SortedIndex += i + 1

		cost1 = v
		cost2 = c[cost2SortedIndex]
		break
	}

	var cost1Index int32 = -1
	var cost2Index int32 = -1

	for i, v := range cost {
		if cost1Index > 0 && cost2Index > 0 {
			break
		}
		if v == cost1 && cost1Index < 0 {
			cost1Index = int32(i + 1)
			if cost1 == cost2 {
				continue
			}
		}
		if v == cost2 {
			if cost1 != cost2 {
				cost2Index = int32(i + 1)
			} else {
				if cost1Index > 0 {
					cost2Index = int32(i + 1)
				}
			}
		}
	}

	if cost1Index > cost2Index {
		cost1Index, cost2Index = cost2Index, cost1Index
	}

	fmt.Println(cost1Index, cost2Index)

	return cost1Index, cost2Index
}

// https://www.hackerrank.com/challenges/swap-nodes-algo/problem
type Node struct {
	value  int32
	left   *Node
	right  *Node
	parent *Node
}

func swapNodes(indexes [][]int32, queries []int32) [][]int32 {
	root := constructTree(indexes)

	result := [][]int32{}

	for _, q := range queries {
		node := root
		swapSingleDepth(node, q, 1)
		result = append(result, traverseInOrder(root, []int32{}))
	}

	return result
}

func constructTree(indexes [][]int32) *Node {
	root := &Node{value: 1}

	nodes := []*Node{root}
	for {
		node := nodes[0]
		nodes = nodes[1:]

		index := indexes[0]
		indexes = indexes[1:]

		if index[0] > 0 {
			node.left = &Node{value: index[0]}
			nodes = append(nodes, node.left)
		}

		if index[1] > 0 {
			node.right = &Node{value: index[1]}
			nodes = append(nodes, node.right)
		}

		if len(indexes) == 0 {
			break
		}
	}

	return root
}

func swapSingleDepth(node *Node, swapDepth int32, depth int32) {
	if node == nil {
		return
	}

	swapSingleDepth(node.left, swapDepth, depth+1)
	swapSingleDepth(node.right, swapDepth, depth+1)

	if depth%swapDepth == 0 {
		node.left, node.right = node.right, node.left
	}
}

func traverseInOrder(node *Node, record []int32) []int32 {
	if node == nil {
		return record
	}

	record = traverseInOrder(node.left, record)
	record = append(record, node.value)
	record = traverseInOrder(node.right, record)

	return record
}

// https://www.hackerrank.com/challenges/pairs/problem
func pairs(target int32, ns []int32) int32 {
	nsc := make([]int32, len(ns))
	copy(nsc, ns)
	sort.Sort(Int32Slice(nsc))

	var result int32

	for i, n := range nsc {
		if n <= target {
			continue
		}

		k := sort.Search(i, func(j int) bool { return nsc[j] >= n-target })

		if k < i && nsc[k] == n-target {
			result++
		}
	}

	return result
}

// https://www.hackerrank.com/challenges/triple-sum/problem
func triplets(a, b, c []int32) int64 {
	a = getSortedUniqueElements(a)
	b = getSortedUniqueElements(b)
	c = getSortedUniqueElements(c)

	var result int64

	for _, q := range b {
		p := int64(sort.Search(len(a), func(i int) bool { return a[i] > q }))
		r := int64(sort.Search(len(c), func(i int) bool { return c[i] > q }))

		result += p * r
	}

	return result
}

func getSortedUniqueElements(ns []int32) []int32 {
	m := make(map[int32]interface{})
	keys := make([]int32, 0)
	for _, n := range ns {
		if _, ok := m[n]; !ok {
			m[n] = nil
			keys = append(keys, n)
		}
	}

	sort.Sort(Int32Slice(keys))

	return keys
}

// https://www.hackerrank.com/challenges/minimum-time-required/problem
func minTime(machines []int64, goal int64) int64 {
	c := make([]int64, len(machines))
	copy(c, machines)
	sort.Sort(Int64Slice(c))

	frequencies := make(map[int64]int64)
	for _, machine := range c {
		frequencies[machine]++
	}

	minTime := int64(1)
	maxTime := machines[0] * goal / frequencies[machines[0]]

	for {
		averageTime := minTime + (maxTime-minTime)/2

		production := int64(0)
		for k, v := range frequencies {
			production += averageTime / k * v
			if production >= goal {
				break
			}
		}

		if production >= goal {
			maxTime = averageTime
		} else {
			minTime = averageTime + 1
		}

		if minTime == maxTime {
			break
		}
	}

	return minTime
}

// https://www.hackerrank.com/challenges/making-candies/problem
func minimumPasses(machines int64, workers int64, price int64, goal int64) int64 {
	var passes int64
	var prod int64
	var noMorePurchase bool

	if goal < price {
		noMorePurchase = true
	}

	for {
		if checkIfEnoughResources(machines, workers, goal) {
			return passes + 1
		}

		prodRate := machines * workers

		if noMorePurchase {
			return passes + skipPasses(goal, prod, prodRate)
		}

		wallet := prod / price

		if wallet == 0 {
			passesUntilPurchase := skipPasses(price, prod, prodRate)

			passes += passesUntilPurchase
			prod += passesUntilPurchase * prodRate

			if prod >= goal {
				return passes
			}

			continue
		}

		machinesAfterPurchase, workersAfterPurchase := purchaseResources(machines, workers, wallet)
		if checkIfEnoughResources(machinesAfterPurchase, workersAfterPurchase, goal) {
			return passes + 1
		}

		prodAfterPurchase := prod - wallet*price
		prodRateAfterPurchase := machinesAfterPurchase * workersAfterPurchase

		passesWithoutPurchase := skipPasses(goal, prod, prodRate)
		passesWithPurchase := skipPasses(goal, prodAfterPurchase, prodRateAfterPurchase)

		if passesWithoutPurchase >= passesWithPurchase {
			machines = machinesAfterPurchase
			workers = workersAfterPurchase
			prod = prodAfterPurchase
			prodRate = prodRateAfterPurchase
		} else {
			noMorePurchase = true
		}

		prod += prodRate
		passes++

		if prod >= goal {
			return passes
		}
	}
}

func checkIfEnoughResources(machines int64, workers int64, goal int64) bool {
	if machines >= goal || workers >= goal {
		return true
	}

	prodToGoalRate := int64((float64(machines) / float64(goal)) * float64(workers))

	if prodToGoalRate >= 1 {
		return true
	}

	return false
}

func skipPasses(goal int64, production int64, productionRate int64) int64 {
	passes := (goal - production) / productionRate
	if (goal-production)%productionRate > 0 {
		passes++
	}

	return passes
}

func purchaseResources(machines int64, workers int64, wallet int64) (int64, int64) {
	if machines == workers {
		machines += wallet / 2
		workers += wallet - wallet/2
	} else {
		for {
			if wallet == 0 {
				break
			}
			if machines < workers {
				machines++
			} else {
				workers++
			}
			wallet--
		}
	}

	return machines, workers
}

// https://www.hackerrank.com/challenges/maximum-subarray-sum/problem
func maximumSum(a []int64, m int64) int64 {
	c := make([]int64, len(a))
	copy(c, a)

	var sum int64 = 0

	indices := make(map[int64][]int)

	var maxSingle int64
	for i := 0; i < len(c); i++ {
		c[i] %= m
		if c[i] == m-1 {
			return c[i]
		} else if c[i] > maxSingle {
			maxSingle = c[i]
		}

		sum = (sum + c[i]) % m
		if sum == m-1 {
			return sum
		} else if sum > maxSingle {
			maxSingle = sum
		}

		c[i] = sum

		if _, ok := indices[sum]; !ok {
			indices[sum] = []int{-1, -1}
		}

		if indices[sum][0] < 0 || indices[sum][0] > i {
			indices[sum][0] = i
		}
		if indices[sum][1] < 0 || indices[sum][1] <= i {
			indices[sum][1] = i
		}
	}

	sort.Sort(Int64Slice(c))

	var result int64

	for i := 0; i < len(c)-1; i++ {
		if c[i] == c[i+1] {
			continue
		}

		if indices[c[i+1]][0] < indices[c[i]][1] {
			current := (c[i] - c[i+1] + m) % m
			if current > result {
				result = current
			}
		}
	}

	if maxSingle > result {
		result = maxSingle
	}

	return result
}
