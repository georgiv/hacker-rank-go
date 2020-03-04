// https://www.hackerrank.com/interview/interview-preparation-kit/stacks-queues/challenges
package stacks_and_queues

type IntNode struct {
	value int
	next  *IntNode
}

type IntStack struct {
	top *IntNode
}

func (s *IntStack) IsEmpty() bool {
	return s.top == nil
}

func (s *IntStack) Push(e int) {
	node := &IntNode{e, s.top}
	s.top = node
}

func (s *IntStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	e := s.top.value
	s.top = s.top.next

	return e, true
}

func (s *IntStack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	return s.top.value, true
}

type IntQueue struct {
	head *IntNode
	tail *IntNode
}

func (q *IntQueue) IsEmpty() bool {
	return q.head == nil
}

func (q *IntQueue) Push(e int) {
	node := &IntNode{value: e}
	if q.IsEmpty() {
		q.head = node
	} else {
		q.tail.next = node
	}
	q.tail = node
}

func (q *IntQueue) Poll() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	e := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}

	return e, true
}

func (q *IntQueue) PeekHead() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	return q.head.value, true
}

func (q *IntQueue) PeekTail() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	return q.tail.value, true
}

// https://www.hackerrank.com/challenges/balanced-brackets/problem
func isBalanced(s string) string {
	stack := IntStack{}

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack.Push(int(c))
		} else {
			e, ok := stack.Pop()
			if !ok {
				return "NO"
			}

			if (c == ')' && e != '(') || (c == ']' && e != '[') || (c == '}' && e != '{') {
				return "NO"
			}
		}
	}

	if !stack.IsEmpty() {
		return "NO"
	}

	return "YES"
}

// https://www.hackerrank.com/challenges/ctci-queue-using-two-stacks/problem
type IntQueueStackImpl struct {
	head *IntStack
	tail *IntStack
}

func NewIntQueueStackImpl() IntQueueStackImpl {
	return IntQueueStackImpl{&IntStack{}, &IntStack{}}
}

func (q *IntQueueStackImpl) IsEmpty() bool {
	return q.head.IsEmpty()
}

func (q *IntQueueStackImpl) Enqueue(e int) {
	if q.IsEmpty() {
		q.head.Push(e)
	} else {
		q.tail.Push(e)
	}
}

func (q *IntQueueStackImpl) Dequeue() (int, bool) {
	e, ok := q.head.Pop()

	if q.head.IsEmpty() {
		for {
			if q.tail.IsEmpty() {
				break
			}

			e, _ := q.tail.Pop()
			q.head.Push(e)
		}
	}

	if ok {
		return e, ok
	}

	return q.head.Pop()
}

func (q *IntQueueStackImpl) Peek() (int, bool) {
	return q.head.Peek()
}

// https://www.hackerrank.com/challenges/largest-rectangle/problem
func largestRectangle(heights []int32) int64 {
	var result int64

	s := IntStack{}

	heights = append(heights, 0)

	i := 0
	for {
		if i == len(heights) {
			break
		}

		v, ok := s.Peek()
		if !ok || heights[i] > heights[v] {
			s.Push(i)
			i++
		} else {
			width := 0

			top, _ := s.Pop()

			left, ok := s.Peek()
			if !ok {
				width = i
			} else {
				width = i - left - 1
			}

			square := int64(heights[top]) * int64(width)
			if square > result {
				result = square
			}
		}
	}

	return result
}

// https://www.hackerrank.com/challenges/min-max-riddle/problem
func riddle(arr []int64) []int64 {
	s := IntStack{}
	minimums := make(map[int64]int)

	arr = append(arr, -1)

	i := 0
	for {
		if i == len(arr) {
			break
		}

		v, ok := s.Peek()
		if !ok || arr[i] > arr[v] {
			s.Push(i)
			i++
		} else {
			width := 0

			top, _ := s.Pop()

			left, ok := s.Peek()
			if !ok {
				width = i
			} else {
				width = i - left - 1
			}

			w, ok := minimums[arr[top]]
			if !ok {
				minimums[arr[top]] = width
			} else {
				if width > w {
					minimums[arr[top]] = width
				}
			}
		}
	}

	maximums := make(map[int]int64)
	for k, v := range minimums {
		max, ok := maximums[v]
		if !ok {
			maximums[v] = k
		} else {
			if k > max {
				maximums[v] = k
			}
		}
	}

	arr = arr[:len(arr)-1]

	var result = make([]int64, len(arr))

	var max int64
	for i := len(arr); i >= 1; i-- {
		k, _ := maximums[i]
		if k > max {
			max = k
		}

		result[i-1] = max
	}

	return result
}

// https://www.hackerrank.com/challenges/castle-on-the-grid/problem
type Position struct {
	x int32
	y int32
}

type PositionNode struct {
	value *Position
	next  *PositionNode
}

type PositionQueue struct {
	head *PositionNode
	tail *PositionNode
}

func (s *PositionQueue) IsEmpty() bool {
	return s.head == nil
}

func (s *PositionQueue) Push(e *Position) {
	node := &PositionNode{value: e}
	if s.IsEmpty() {
		s.head = node
	} else {
		s.tail.next = node
	}
	s.tail = node
}

func (s *PositionQueue) Poll() (*Position, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	e := s.head.value
	s.head = s.head.next
	if s.head == nil {
		s.tail = nil
	}

	return e, true
}

func (s *PositionQueue) Peek() (*Position, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	return s.head.value, true
}

func minimumMoves(grid []string, startX int32, startY int32, goalX int32, goalY int32) int32 {
	turnsGrid := make([][]int32, len(grid))
	visitedGrid := make([][]bool, len(grid))

	for i, _ := range grid {
		turnsGrid[i] = make([]int32, len(grid))
		visitedGrid[i] = make([]bool, len(grid))
		for j, v := range grid[i] {
			if v == 'X' {
				turnsGrid[i][j] = -1
				visitedGrid[i][j] = true
			}
		}
	}

	var result int32 = -1

	s := PositionQueue{}
	s.Push(&Position{startX, startY})

	for {
		current, _ := s.Poll()
		if visitedGrid[current.x][current.y] {
			continue
		}

		if current.x == goalX && current.y == goalY {
			result = turnsGrid[goalX][goalY]
			break
		}

		fillTurns(turnsGrid, current.x, current.y)

		if current.x != goalX || current.y != goalY {
			visitedGrid[current.x][current.y] = true
		}

		if validateCell(visitedGrid, current.x, current.y-1) {
			s.Push(&Position{current.x, current.y - 1})
		}
		if validateCell(visitedGrid, current.x, current.y+1) {
			s.Push(&Position{current.x, current.y + 1})
		}
		if validateCell(visitedGrid, current.x-1, current.y) {
			s.Push(&Position{current.x - 1, current.y})
		}
		if validateCell(visitedGrid, current.x+1, current.y) {
			s.Push(&Position{current.x + 1, current.y})
		}
	}

	return result
}

func fillTurns(grid [][]int32, posX int32, posY int32) {
	turns := grid[posX][posY] + 1

	for i := posY - 1; i >= 0; i-- {
		if !fillCellTurns(grid, posX, i, turns) {
			break
		}
	}

	for i := posY + 1; i < int32(len(grid)); i++ {
		if !fillCellTurns(grid, posX, i, turns) {
			break
		}
	}

	for i := posX - 1; i >= 0; i-- {
		if !fillCellTurns(grid, i, posY, turns) {
			break
		}
	}

	for i := posX + 1; i < int32(len(grid)); i++ {
		if !fillCellTurns(grid, i, posY, turns) {
			break
		}
	}
}

func fillCellTurns(grid [][]int32, posX int32, posY int32, turns int32) bool {
	if grid[posX][posY] == -1 {
		return false
	}

	if grid[posX][posY] == 0 {
		grid[posX][posY] = turns
	} else if grid[posX][posY] > 0 {
		if turns < grid[posX][posY] {
			grid[posX][posY] = turns
		}
	}

	return true
}

func validateCell(grid [][]bool, posX int32, posY int32) bool {
	return posX >= 0 && posX < int32(len(grid)) && posY >= 0 && posY < int32(len(grid))
}

func poisonousPlants(plants []int32) int32 {
	var qs []IntQueue

	q := IntQueue{}

	for _, v := range plants {
		if q.IsEmpty() {
			q.Push(int(v))
		} else {
			t, _ := q.PeekTail()
			if int(v) <= t {
				q.Push(int(v))
			} else {
				qs = append(qs, q)
				q = IntQueue{}
				q.Push(int(v))
			}
		}
	}

	qs = append(qs, q)

	var days int32
	var deadPlants bool

	for {
		qsShadow := []IntQueue{qs[0]}

		for i := 1; i < len(qs); i++ {
			t, _ := qsShadow[len(qsShadow)-1].PeekTail()
			h, _ := qs[i].PeekHead()
			if t >= h {
				for {
					if qs[i].IsEmpty() {
						break
					}
					e, _ := qs[i].Poll()
					qsShadow[len(qsShadow)-1].Push(e)
				}
			} else {
				qsShadow = append(qsShadow, qs[i])
			}
		}

		if len(qsShadow) == 1 {
			break
		}

		qs = qsShadow
		qsShadow = []IntQueue{qs[0]}

		for i := 1; i < len(qs); i++ {
			t, _ := qsShadow[len(qsShadow)-1].PeekTail()
			h, _ := qs[i].PeekHead()
			if t < h {
				qs[i].Poll()
				if !qs[i].IsEmpty() {
					qsShadow = append(qsShadow, qs[i])
				}
				deadPlants = true
			}
		}

		qs = qsShadow
		if deadPlants {
			days++
			deadPlants = false
		}
	}

	return days
}
