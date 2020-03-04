// https://www.hackerrank.com/interview/interview-preparation-kit/dictionaries-hashmaps/challenges
package dictionaries_and_hashmaps

import (
	"fmt"
	"sort"
)

// https://www.hackerrank.com/challenges/ctci-ransom-note/problem
func checkMagazine(magazine []string, note []string) string {
	if len(magazine) < len(note) {
		fmt.Println("No")
		return "No"
	}

	magazineMap := make(map[string]int)
	for _, w := range magazine {
		magazineMap[w] += 1
	}

	for _, w := range note {
		occurrence, ok := magazineMap[w]
		if !ok || occurrence == 0 {
			fmt.Println("No")
			return "No"
		}
		magazineMap[w] -= 1
	}

	fmt.Println("Yes")
	return "Yes"
}

// https://www.hackerrank.com/challenges/two-strings/problem
func twoStrings(s1 string, s2 string) string {
	if s1 == "" || s2 == "" {
		return "NO"
	}

	m := make(map[int32]bool)

	for _, c1 := range s1 {
		m[c1] = true
	}

	for _, c2 := range s2 {
		if _, ok := m[c2]; ok {
			return "YES"
		}
	}

	return "NO"
}

// https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem
func sherlockAndAnagrams(s string) int32 {
	m := make(map[string]int)

	for i := 0; i < len(s); i++ {
		var sub []int
		for j := i; j < len(s); j++ {
			sub = append(sub, int(s[j]))
			sort.Ints(sub)
			m[fmt.Sprint(sub)]++
		}
	}

	result := 0
	for _, v := range m {
		if v > 1 {
			result += v * (v - 1) / 2
		}
	}

	return int32(result)
}

// https://www.hackerrank.com/challenges/count-triplets-1/problem
func countTriplets(arr []int64, r int64) int64 {
	indices := make(map[int64][]int)

	for i, n := range arr {
		indices[n] = append(indices[n], i)
	}

	var result int64

	for i, n := range arr {
		if n%r != 0 {
			continue
		}

		firstIndices, ok := indices[n/r]
		if !ok {
			continue
		}

		thirdIndices, ok := indices[n*r]
		if !ok {
			continue
		}

		firstIndicesCount := sort.Search(len(firstIndices), func(j int) bool { return firstIndices[j] >= i })
		thirdIndicesCount := len(thirdIndices) - sort.Search(len(thirdIndices), func(j int) bool { return thirdIndices[j] > i })

		result += int64(firstIndicesCount) * int64(thirdIndicesCount)
	}

	return result
}

// https://www.hackerrank.com/challenges/frequency-queries/problem
func freqQuery(queries [][]int32) []int32 {
	var opAdd int32 = 1
	var opDel int32 = 2
	var opChk int32 = 3

	var result []int32

	operations := make(map[int32]int32)
	frequencies := make(map[int32]int32)

	for _, q := range queries {
		switch q[0] {
		case opAdd:
			freq := operations[q[1]]
			if freq > 0 {
				frequencies[freq]--
				if frequencies[freq] == 0 {
					delete(frequencies, freq)
				}
			}
			operations[q[1]]++
			freq++
			frequencies[freq]++
		case opDel:
			freq := operations[q[1]]
			if freq == 0 {
				break
			}
			frequencies[freq]--
			if frequencies[freq] == 0 {
				delete(frequencies, freq)
			}
			operations[q[1]]--
			if operations[q[1]] == 0 {
				delete(operations, q[1])
			} else {
				freq--
				frequencies[freq]++
			}
		case opChk:
			var found int32 = 0
			if _, ok := frequencies[q[1]]; ok {
				found = 1
			}
			result = append(result, found)
		}
	}

	return result
}
