// https://www.hackerrank.com/interview/interview-preparation-kit/strings/challenges
package string_manipulation

// https://www.hackerrank.com/challenges/ctci-making-anagrams/problem
func makeAnagram(a string, b string) int32 {
	frequencies := [26]int{}

	for _, c := range a {
		frequencies[c-97]++
	}

	for _, c := range b {
		frequencies[c-97]--
	}

	result := 0
	for _, f := range frequencies {
		if f < 0 {
			f = -f
		}
		result += f
	}

	return int32(result)
}

// https://www.hackerrank.com/challenges/alternating-characters/problem
func alternatingCharacters(s string) int32 {
	var result int32

	var current int32

	for i, c := range s {
		if i == 0 {
			current = c
			continue
		}

		if current == c {
			result++
		} else {
			current = c
		}
	}

	return result
}

// https://www.hackerrank.com/challenges/sherlock-and-valid-string/problem
func isValid(s string) string {
	frequencies := [26]int{}

	for _, c := range s {
		frequencies[c-97]++
	}

	freq1 := 0
	freq2 := 0

	isFreq1Repeating := false
	isFreq2Repeating := false

	for _, f := range frequencies {
		if f == 0 {
			continue
		}

		if f == freq1 {
			if isFreq2Repeating {
				return "NO"
			}
			isFreq1Repeating = true
			continue
		} else if f == freq2 {
			if isFreq1Repeating {
				return "NO"
			}
			isFreq2Repeating = true
			continue
		}

		if freq1 == 0 {
			freq1 = f
		} else if freq2 == 0 {
			freq2 = f
		} else {
			return "NO"
		}
	}

	if freq1 == 0 || freq2 == 0 || (freq1 == 1 && !isFreq1Repeating) || (freq2 == 1 && !isFreq2Repeating) {
		return "YES"
	}

	diff := freq1 - freq2
	if diff == 1 || diff == -1 {
		return "YES"
	} else {
		return "NO"
	}
}

// https://www.hackerrank.com/challenges/special-palindrome-again/problem
func substrCount(n int32, s string) int64 {
	var res int64

	for i := 0; i < len(s); i++ {
		left := i - 1
		right := i + 1
		for {
			if left < 0 || right >= len(s) || s[left] != s[right] || s[left] != s[i+1] || s[left] == s[i] {
				break
			}

			res++
			left--
			right++
		}
	}

	repeatingCharCount := 1
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && s[i] == s[i+1] {
			repeatingCharCount++
		} else {
			res += int64(repeatingCharCount * (repeatingCharCount + 1) / 2)
			repeatingCharCount = 1
		}
	}

	return res
}

// https://www.hackerrank.com/challenges/common-child/problem
func commonChild(s1 string, s2 string) int32 {
	counter := [][]int{}
	for i := 0; i <= len(s1); i++ {
		counter = append(counter, []int{})
		for j := 0; j <= len(s2); j++ {
			counter[i] = append(counter[i], 0)
		}
	}

	for i, c1 := range s1 {
		for j, c2 := range s2 {
			if c1 == c2 {
				counter[i+1][j+1] = counter[i][j] + 1
			} else {
				if counter[i+1][j] > counter[i][j+1] {
					counter[i+1][j+1] = counter[i+1][j]
				} else {
					counter[i+1][j+1] = counter[i][j+1]
				}
			}
		}
	}

	return int32(counter[len(s1)][len(s2)])
}
