// https://www.hackerrank.com/interview/interview-preparation-kit/warmup/challenges
package warmup

// https://www.hackerrank.com/challenges/sock-merchant/problem
func sockMerchant(n int32, arr []int32) int32 {
	var result int32

	frequency := make(map[int32]int)

	for _, e := range arr {
		frequency[e] += 1
	}

	for _, v := range frequency {
		result += int32(v / 2)
	}

	return result
}

// https://www.hackerrank.com/challenges/counting-valleys/problem
func countingValleys(n int32, path string) int32 {
	var result int32

	altitude := 0
	isValley := path[0] == 'D'

	for i, s := range path {
		if s == 'U' {
			altitude++
		} else if s == 'D' {
			altitude--
		}

		if altitude == 0 {
			if isValley {
				result++
			}
			if i < len(path)-1 {
				isValley = path[i+1] == 'D'
			}
		}
	}

	return result
}

// https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem
func jumpingOnClouds(clouds []int32) int32 {
	var result int32

	i := 0

	for {
		if i >= len(clouds)-3 {
			result++
			break
		}

		if clouds[i+2] == 0 {
			i += 2
		} else if clouds[i+1] == 0 {
			i++
		}

		result++
	}

	return result
}

// https://www.hackerrank.com/challenges/repeated-string/problem
func repeatedString(s string, n int64) int64 {
	var occurrencesFull int64
	var occurrencesRem int64

	magnitude := n / int64(len(s))
	remainder := n % int64(len(s))

	for i, c := range s {
		if c == 'a' {
			if int64(i) < remainder {
				occurrencesRem++
			}
			occurrencesFull++
		}
	}

	return occurrencesFull*magnitude + occurrencesRem
}
