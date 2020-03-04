// https://www.hackerrank.com/interview/interview-preparation-kit/recursion-backtracking/challenges
package recursion_and_backtracking

import (
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/ctci-fibonacci-numbers/problem
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-2) + fibonacci(n-1)
}

// https://www.hackerrank.com/challenges/ctci-recursive-staircase/problem
func stepPerms(steps int32) int32 {
	return stepPermsCached(steps, make(map[int32]int32))
}

func stepPermsCached(steps int32, cache map[int32]int32) int32 {
	if steps < 0 {
		return 0
	} else if steps == 0 {
		return 1
	}

	if v, ok := cache[steps]; ok {
		return v
	}

	perms := stepPermsCached(steps-1, cache) + stepPermsCached(steps-2, cache) + stepPermsCached(steps-3, cache)

	cache[steps] = perms

	return perms
}

// https://www.hackerrank.com/challenges/recursive-digit-sum/problem
func superDigit(n string, k int32) int32 {
	return superDigitInternal(strconv.Itoa(int(superDigitInternal(n, 0)*k)), 0)
}

func superDigitInternal(n string, sd int32) int32 {
	if len(n) == 0 {
		return sd
	}

	sd += int32(n[0] - 48)
	if sd > 9 {
		sd = sd/10 + sd%10
	}

	return superDigitInternal(n[1:], sd)
}

// https://www.hackerrank.com/challenges/crossword-puzzle/problem
const asciiFilledCell1 = 43
const asciiFilledCell2 = 88
const asciiEmptyCell = 45

func crosswordPuzzle(crossword []string, words string) []string {
	if len(words) == 0 {
		return crossword
	}

	sepIndex := strings.Index(words, ";")
	if sepIndex == -1 {
		sepIndex = len(words)
	}
	word := words[:sepIndex]

	remainder := ""
	if sepIndex < len(words) {
		remainder = words[sepIndex+1:]
	}

	for i := 0; i < len(crossword); i++ {
		for j := 0; j < len(crossword[0]); j++ {
			if isHorizontalWordPossible(crossword, word, i, j) {
				cw := fillHorizontalWord(crossword, word, i, j)
				cw = crosswordPuzzle(cw, remainder)
				if cw != nil {
					return cw
				}
			}
			if isVerticalWordPossible(crossword, word, i, j) {
				cw := fillVerticalWord(crossword, word, i, j)
				cw = crosswordPuzzle(cw, remainder)
				if cw != nil {
					return cw
				}
			}
		}
	}

	return nil
}

func isHorizontalWordPossible(crossword []string, word string, row, column int) bool {
	if column+len(word) > len(crossword[0]) {
		return false
	}

	if column > 0 &&
		crossword[row][column-1] != asciiFilledCell1 && crossword[row][column-1] != asciiFilledCell2 {
		return false
	}

	if column+len(word) < len(crossword[0]) && crossword[row][column+len(word)] != asciiFilledCell1 && crossword[row][column+len(word)] != asciiFilledCell2 {
		return false
	}

	emptyCellFound := false
	for i := 0; i < len(word); i++ {
		cell := crossword[row][column+i]
		if cell == asciiEmptyCell {
			emptyCellFound = true
			continue
		}
		if cell == asciiFilledCell1 || cell == asciiFilledCell2 || cell != word[i] {
			return false
		}
	}

	if !emptyCellFound {
		return false
	}

	return true
}

func isVerticalWordPossible(crossword []string, word string, row, column int) bool {
	if row+len(word) > len(crossword) {
		return false
	}

	if row > 0 && crossword[row-1][column] != asciiFilledCell1 && crossword[row-1][column] != asciiFilledCell2 {
		return false
	}

	if row+len(word) < len(crossword) && crossword[row+len(word)][column] != asciiFilledCell1 && crossword[row+len(word)][column] != asciiFilledCell2 {
		return false
	}

	emptyCellFound := false
	for i := 0; i < len(word); i++ {
		cell := crossword[row+i][column]
		if cell == asciiEmptyCell {
			emptyCellFound = true
			continue
		}
		if cell == asciiFilledCell1 || cell == asciiFilledCell2 || cell != word[i] {
			return false
		}
	}

	if !emptyCellFound {
		return false
	}

	return true
}

func fillHorizontalWord(crossword []string, word string, row int, col int) []string {
	result := make([]string, len(crossword))

	for i := 0; i < len(result); i++ {
		if i == row {
			if col+len(word) < len(crossword[i]) {
				word += crossword[i][col+len(word):]
			}
			if col > 0 {
				word = crossword[i][0:col] + word
			}
			result[i] = word
		} else {
			result[i] = crossword[i]
		}
	}

	return result
}

func fillVerticalWord(crossword []string, word string, row int, col int) []string {
	result := make([]string, len(crossword))

	for i := 0; i < len(result); i++ {
		if i >= row && i < row+len(word) {
			w := ""
			if col == 0 {
				w = word[i-row:i-row+1] + crossword[i][1:]
			} else if col == len(crossword[i])-1 {
				w = crossword[i][0:len(crossword[i])-1] + word[i-row:i-row+1]
			} else {
				w = crossword[i][0:col] + word[i-row:i-row+1] + crossword[i][col+1:len(crossword[i])]
			}
			result[i] = w
		} else {
			result[i] = crossword[i]
		}
	}

	return result
}
