// https://www.hackerrank.com/interview/interview-preparation-kit/sorting/challenges
package sorting

import (
	"fmt"
	"sort"
)

// https://www.hackerrank.com/challenges/ctci-bubble-sort/problem
func countSwaps(a []int32) string {
	c := make([]int32, len(a))
	copy(c, a)

	counter := 0

	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c)-1; j++ {
			if c[j] > c[j+1] {
				c[j], c[j+1] = c[j+1], c[j]
				counter++
			}
		}
	}

	result := fmt.Sprintf("Array is sorted in %d swaps.\nFirst Element: %d\nLast Element: %d", counter, c[0], c[len(c)-1])

	fmt.Println(result)

	return result
}

// https://www.hackerrank.com/challenges/mark-and-toys/problem
func maximumToys(prices []int32, money int32) int32 {
	var c []int
	for _, p := range prices {
		c = append(c, int(p))
	}

	sort.Ints(c)

	k := int(money)

	var result int32

	for _, p := range c {
		if p > k {
			break
		}
		k -= p
		result++
	}

	return result
}

// https://www.hackerrank.com/challenges/ctci-comparator-sorting/problem
type Player struct {
	name  string
	score int
}

type PlayersByScore []Player

func (ps PlayersByScore) Len() int {
	return len(ps)
}

func (ps PlayersByScore) Less(i, j int) bool {
	if ps[i].score != ps[j].score {
		return ps[i].score < ps[j].score
	}

	return ps[i].name < ps[j].name
}

func (ps PlayersByScore) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// https://www.hackerrank.com/challenges/fraudulent-activity-notifications/problem
func activityNotifications(expenditure []int32, d int32) int32 {
	if len(expenditure) <= int(d) {
		return 0
	}

	var result int32

	history := [201]int{}
	historyDays := int(d)

	for i := 0; i < historyDays; i++ {
		history[expenditure[i]]++
	}

	for i := int(d); i < len(expenditure); i++ {
		median := -1.0

		if historyDays%2 == 0 {
			medianLeft := -1.0
			medianRight := -1.0
			medianIndexLeft := historyDays/2 - 1
			medianIndexRight := historyDays / 2
			indexCounter := 0
			for i, v := range history {
				if v == 0 {
					continue
				}

				indexCounter += v
				if indexCounter >= medianIndexLeft+1 {
					if medianLeft < 0 {
						medianLeft = float64(i)
					}
				}
				if indexCounter >= medianIndexRight+1 {
					medianRight = float64(i)
					break
				}
			}

			median = (medianLeft + medianRight) / 2
		} else {
			medianIndex := historyDays / 2
			indexCounter := 0
			for i, v := range history {
				if v == 0 {
					continue
				}

				indexCounter += v
				if indexCounter >= medianIndex+1 {
					median = float64(i)
					break
				}
			}
		}

		if float64(expenditure[i]) >= median*2 {
			result++
		}

		history[expenditure[i-historyDays]]--
		history[expenditure[i]]++
	}

	return result
}

// https://www.hackerrank.com/challenges/ctci-merge-sort/problem
func countInversions(arr []int32) int64 {
	c := make([]int32, len(arr))
	copy(c, arr)
	tmp := make([]int32, len(arr))
	return mergesort(c, tmp, 0, len(arr)-1)
}

func mergesort(arr []int32, tmp []int32, left int, right int) int64 {
	if left >= right {
		return 0
	}

	middle := (left + right) / 2
	return mergesort(arr, tmp, left, middle) + mergesort(arr, tmp, middle+1, right) + merge(arr, tmp, left, middle, right)
}

func merge(arr []int32, tmp []int32, leftStart int, middle int, rightEnd int) int64 {
	leftEnd := middle
	rightStart := leftEnd + 1

	index := leftStart
	leftIndex := leftStart
	rightIndex := rightStart

	var swaps int64 = 0

	for {
		if leftIndex > leftEnd || rightIndex > rightEnd {
			break
		}
		if arr[leftIndex] <= arr[rightIndex] {
			tmp[index] = arr[leftIndex]
			leftIndex++
		} else {
			tmp[index] = arr[rightIndex]
			rightIndex++
			swaps += int64(rightStart - leftIndex)
		}
		index++
	}

	i := index
	for {
		if leftIndex > leftEnd {
			break
		}

		tmp[i] = arr[leftIndex]
		i++
		leftIndex++
	}

	i = index
	for {
		if rightIndex > rightEnd {
			break
		}

		tmp[i] = arr[rightIndex]
		i++
		rightIndex++
	}

	i = 0
	for {
		if i >= rightEnd-leftStart+1 {
			break
		}
		arr[leftStart+i] = tmp[leftStart+i]
		i++
	}

	return swaps
}
