// https://www.hackerrank.com/interview/interview-preparation-kit/arrays/challenges
package arrays

import (
	"fmt"
	"math"
)

// https://www.hackerrank.com/challenges/2d-array/problem
func hourglassSum(arr [][]int32) int32 {
	result := int32(math.MinInt32)

	for i := 1; i < len(arr)-1; i++ {
		for j := 1; j < len(arr[i])-1; j++ {
			sum := arr[i-1][j-1] +
				arr[i-1][j] +
				arr[i-1][j+1] +
				arr[i][j] +
				arr[i+1][j-1] +
				arr[i+1][j] +
				arr[i+1][j+1]
			if sum > result {
				result = sum
			}
		}
	}

	return result
}

// https://www.hackerrank.com/challenges/ctci-array-left-rotation/problem
func rotLeft(arr []int32, offset int32) []int32 {
	var result []int32

	i := offset % int32(len(arr))

	for {
		if len(arr) == len(result) {
			break
		}
		if i == int32(len(arr)) {
			i = 0
		}

		result = append(result, arr[i])

		i++
	}

	return result
}

// https://www.hackerrank.com/challenges/new-year-chaos/problem
func minimumBribes(queue []int32) int32 {
	var result int32

	data := make([]int32, len(queue))
	copy(data, queue)

	for i := len(data) - 1; i >= 2; i-- {
		if data[i] == int32(i+1) {
			continue
		}

		if data[i-1] == int32(i+1) {
			data[i], data[i-1] = data[i-1], data[i]
			result++
		} else if data[i-2] == int32(i+1) {
			data[i-2], data[i-1] = data[i-1], data[i-2]
			data[i-1], data[i] = data[i], data[i-1]
			result += 2
		} else {
			fmt.Println("Too chaotic")

			return -1
		}
	}

	if data[0] != 1 {
		result++
	}

	fmt.Println(result)

	return result
}

// https://www.hackerrank.com/challenges/minimum-swaps-2/problem
func minimumSwaps(arr []int32) int32 {
	var result int32

	data := make([]int32, len(arr))
	copy(data, arr)

	var i int32 = 0

	for {
		if i == int32(len(data)) {
			break
		}

		if data[i] != i+1 {
			data[i], data[data[i]-1] = data[data[i]-1], data[i]
			result++
		} else {
			i++
		}
	}

	return result
}

// https://www.hackerrank.com/challenges/crush/problem
func arrayManipulation(n int32, queries [][]int32) int64 {
	var result int64

	data := make([]int64, n)

	for _, q := range queries {
		data[q[0]-1] += int64(q[2])

		if q[1] < n {
			data[q[1]] -= int64(q[2])
		}
	}

	var sum int64

	for _, e := range data {
		sum += e
		if sum > result {
			result = sum
		}
	}

	return result
}
