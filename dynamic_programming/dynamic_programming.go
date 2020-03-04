// https://www.hackerrank.com/interview/interview-preparation-kit/dynamic-programming/challenges
package dynamic_programming

import "fmt"

// https://www.hackerrank.com/challenges/max-array-sum/problem
func maxSubsetSum(ns []int32) int32 {
	if len(ns) == 1 {
		return ns[0]
	}

	result := ns[0]
	if ns[1] > ns[0] {
		result = ns[1]
	}

	ns[1] = result

	for i := 2; i < len(ns); i++ {
		if ns[i] > result {
			result = ns[i]
		}

		if ns[i]+ns[i-2] > result {
			result = ns[i] + ns[i-2]
		}

		ns[i] = result
	}

	return result
}

// https://www.hackerrank.com/challenges/abbr/problem
func abbreviation(a string, b string) string {
	validity := make([][]bool, len(a)+1)
	for i := 0; i < len(validity); i++ {
		validity[i] = make([]bool, len(b)+1)
	}

	validity[0][0] = true

	uppercaseFound := false

	for i := 1; i < len(validity); i++ {
		if uppercaseFound || a[i-1] >= 'A' && a[i-1] <= 'Z' {
			uppercaseFound = true
		} else {
			validity[i][0] = true
		}
	}

	for i := 1; i < len(validity); i++ {
		for j := 1; j < len(validity[i]); j++ {
			if a[i-1] == b[j-1] {
				validity[i][j] = validity[i-1][j-1]
			} else if a[i-1]-32 == b[j-1] {
				validity[i][j] = validity[i-1][j-1] || validity[i-1][j]
			} else if a[i-1] >= 'A' && a[i-1] <= 'Z' {
				validity[i][j] = false
			} else {
				validity[i][j] = validity[i-1][j]
			}
		}
	}

	if validity[len(a)][len(b)] {
		return "YES"
	} else {
		return "NO"
	}
}

// https://www.hackerrank.com/challenges/candies/problem
func candies(students int32, ratings []int32) int64 {
	candies := make([]int64, len(ratings))
	for i := 0; i < len(candies); i++ {
		candies[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	var result int64
	for _, c := range candies {
		result += c
	}

	fmt.Println(fmt.Sprintf("%v", candies))

	return result
}

// https://www.hackerrank.com/challenges/decibinary-numbers/problem
func decibinaryNumbers(x int64) int64 {

	return -1
}
