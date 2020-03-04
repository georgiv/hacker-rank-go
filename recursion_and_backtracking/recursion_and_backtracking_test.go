package recursion_and_backtracking

import (
	"strings"
	"testing"
)

func TestFibonacci(t *testing.T) {
	var tests = []struct {
		n        int
		expected int
	}{
		{3, 2},
		{5, 5},
	}

	for _, test := range tests {
		if result := fibonacci(test.n); result != test.expected {
			t.Errorf("fibonacci(%d) = %d, expected = %d", test.n, result, test.expected)
		}
	}
}

func TestStepPerms(t *testing.T) {
	var tests = []struct {
		steps    int32
		expected int32
	}{
		{1, 1},
		{3, 4},
		{7, 44},
		{35, 1132436852},
	}

	for _, test := range tests {
		if result := stepPerms(test.steps); result != test.expected {
			t.Errorf("stepPerms(%d) = %d, expected = %d", test.steps, result, test.expected)
		}
	}
}

func TestSuperDigit(t *testing.T) {
	var tests = []struct {
		n        string
		k        int32
		expected int32
	}{
		{"9875", 4, 8},
		{"148", 3, 3},
		{"123", 3, 9},
	}

	for _, test := range tests {
		if result := superDigit(test.n, test.k); result != test.expected {
			t.Errorf("superDigit(%s, %d) = %d, expected = %d", test.n, test.k, result, test.expected)
		}
	}
}

func TestCrosswordPuzzle(t *testing.T) {
	var tests = []struct {
		crossword []string
		words     string
		expected  []string
	}{
		{[]string{
			"++++++++++",
			"+------+++",
			"+++-++++++",
			"+++-++++++",
			"+++-----++",
			"+++-++-+++",
			"++++++-+++",
			"++++++-+++",
			"++++++-+++",
			"++++++++++",
		},
			"POLAND;LHASA;SPAIN;INDIA",
			[]string{
				"++++++++++",
				"+POLAND+++",
				"+++H++++++",
				"+++A++++++",
				"+++SPAIN++",
				"+++A++N+++",
				"++++++D+++",
				"++++++I+++",
				"++++++A+++",
				"++++++++++",
			}},
		{[]string{
			"+-++++++++",
			"+-++++++++",
			"+-++++++++",
			"+-----++++",
			"+-+++-++++",
			"+-+++-++++",
			"+++++-++++",
			"++------++",
			"+++++-++++",
			"+++++-++++",
		},
			"LONDON;DELHI;ICELAND;ANKARA",
			[]string{
				"+L++++++++",
				"+O++++++++",
				"+N++++++++",
				"+DELHI++++",
				"+O+++C++++",
				"+N+++E++++",
				"+++++L++++",
				"++ANKARA++",
				"+++++N++++",
				"+++++D++++",
			}},
		{[]string{
			"+-++++++++",
			"+-++++++++",
			"+-------++",
			"+-++++++++",
			"+-++++++++",
			"+------+++",
			"+-+++-++++",
			"+++++-++++",
			"+++++-++++",
			"++++++++++",
		},
			"AGRA;NORWAY;ENGLAND;GWALIOR",
			[]string{
				"+E++++++++",
				"+N++++++++",
				"+GWALIOR++",
				"+L++++++++",
				"+A++++++++",
				"+NORWAY+++",
				"+D+++G++++",
				"+++++R++++",
				"+++++A++++",
				"++++++++++",
			}},
		{[]string{
			"XXXXXX-XXX",
			"XX------XX",
			"XXXXXX-XXX",
			"XXXXXX-XXX",
			"XXX------X",
			"XXXXXX-X-X",
			"XXXXXX-X-X",
			"XXXXXXXX-X",
			"XXXXXXXX-X",
			"XXXXXXXX-X",
		},
			"ICELAND;MEXICO;PANAMA;ALMATY",
			[]string{
				"XXXXXXIXXX",
				"XXMEXICOXX",
				"XXXXXXEXXX",
				"XXXXXXLXXX",
				"XXXPANAMAX",
				"XXXXXXNXLX",
				"XXXXXXDXMX",
				"XXXXXXXXAX",
				"XXXXXXXXTX",
				"XXXXXXXXYX",
			}},
	}

	for _, test := range tests {
		if result := crosswordPuzzle(test.crossword, test.words); !arrayDeepEqual(result, test.expected) {
			t.Errorf("crosswordPuzzle(\n%v, \n%s) = \n%v, \nexpected = \n%v",
				strings.Join(test.crossword, "\n"),
				test.words,
				strings.Join(result, "\n"),
				strings.Join(test.expected, "\n"))
		}
	}
}

func arrayDeepEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
