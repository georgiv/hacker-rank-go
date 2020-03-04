// https://www.hackerrank.com/interview/interview-preparation-kit/greedy-algorithms/challenges
package greedy_algorithms

import (
	"bytes"
	"sort"
)

type Int32Slice []int32

func (s Int32Slice) Len() int {
	return len(s)
}

func (s Int32Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Int32Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// https://www.hackerrank.com/challenges/minimum-absolute-difference-in-an-array/problem
func minimumAbsoluteDifference(ns []int32) int32 {
	nsc := make([]int32, len(ns))
	copy(nsc, ns)
	sort.Sort(Int32Slice(nsc))

	result := nsc[len(nsc)-1] - nsc[0]

	for i := 1; i < len(nsc); i++ {
		diff := nsc[i] - nsc[i-1]
		if diff < result {
			result = diff
		}
	}

	return result
}

// https://www.hackerrank.com/challenges/luck-balance/problem
type Contest struct {
	luck       int32
	importance int32
}

type ContestSlice []Contest

func (s ContestSlice) Len() int {
	return len(s)
}

func (s ContestSlice) Less(i, j int) bool {
	return s[i].luck < s[j].luck
}

func (s ContestSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func luckBalance(k int32, contests [][]int32) int32 {
	var result int32

	var importantContests ContestSlice

	for _, c := range contests {
		if c[1] == 0 {
			result += c[0]
		} else if c[1] == 1 {
			importantContests = append(importantContests, Contest{c[0], c[1]})
		}
	}

	sort.Sort(importantContests)

	for i, c := range importantContests {
		if i <= len(importantContests)-int(k)-1 {
			result -= c.luck
		} else {
			result += c.luck
		}
	}

	return result
}

// https://www.hackerrank.com/challenges/greedy-florist/problem
func getMinimumCost(buyers int32, prices []int32) int32 {
	pricesc := make(Int32Slice, len(prices))
	copy(pricesc, prices)
	sort.Sort(pricesc)

	var result int32

	var multiplier int32 = 1
	freeBuyers := buyers

	for i := len(pricesc) - 1; i >= 0; i-- {
		if freeBuyers == 0 {
			freeBuyers = buyers
			multiplier++
		}

		result += multiplier * pricesc[i]

		freeBuyers--
	}

	return result
}

// https://www.hackerrank.com/challenges/angry-children/problem
func maxMin(k int32, ns []int32) int32 {
	nsc := make(Int32Slice, len(ns))
	copy(nsc, ns)
	sort.Sort(nsc)

	result := nsc[len(nsc)-1] - nsc[0]

	for i := 0; i <= len(nsc)-int(k); i++ {
		if result > nsc[i+int(k)-1]-nsc[i] {
			result = nsc[i+int(k)-1] - nsc[i]
		}
	}

	return result
}

// https://www.hackerrank.com/challenges/reverse-shuffle-merge/problem
func reverseShuffleMerge(s string) string {
	frequencies := [26]int{}
	for _, c := range s {
		frequencies[c-'a']++
	}

	resultFrequencies := [26]int{}
	for i := 0; i < len(resultFrequencies); i++ {
		resultFrequencies[i] = frequencies[i] / 2
	}

	var resultChars []int

	for i := len(s) - 1; i >= 0; i-- {
		frequencies[s[i]-'a']--

		if resultFrequencies[s[i]-'a'] == 0 {
			continue
		}

		resultFrequencies[s[i]-'a']--

		for {
			if len(resultChars) == 0 ||
				resultChars[len(resultChars)-1] <= int(s[i]) ||
				frequencies[resultChars[len(resultChars)-1]-'a'] == resultFrequencies[resultChars[len(resultChars)-1]-'a'] {
				break
			}

			resultFrequencies[resultChars[len(resultChars)-1]-'a']++
			resultChars = resultChars[:len(resultChars)-1]
		}

		resultChars = append(resultChars, int(s[i]))
	}

	var buffer bytes.Buffer
	for _, c := range resultChars {
		buffer.WriteString(string(c))
	}

	return buffer.String()
}
