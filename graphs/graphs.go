// https://www.hackerrank.com/interview/interview-preparation-kit/graphs/challenges
package graphs

import "sort"

type Vertex struct {
	value    int32
	color    int32
	children []*Vertex
}

type Graph struct {
	vertices map[int32]*Vertex
	directed bool
}

func NewGraph(directed bool) *Graph {
	return &Graph{make(map[int32]*Vertex), directed}
}

func (g *Graph) IsDirected() bool {
	return g.directed
}

func (g *Graph) Connect(from, to int32) {
	fromVertex := g.Vertex(from, true)
	toVertex := g.Vertex(to, true)
	fromVertex.children = append(fromVertex.children, toVertex)

	if !g.IsDirected() {
		toVertex.children = append(toVertex.children, fromVertex)
	}
}

func (g *Graph) Vertex(n int32, create bool) *Vertex {
	v, ok := g.vertices[n]
	if !ok && create {
		v = &Vertex{value: n}
		g.vertices[n] = v
	}

	return v
}

func (g *Graph) Connected(from, to int32) bool {
	fromVertex := g.Vertex(from, false)
	if fromVertex == nil {
		return false
	}

	toVertex := g.Vertex(to, false)
	if toVertex == nil {
		return false
	}

	var result bool

	q := []*Vertex{fromVertex}
	visited := make(map[int32]bool)

	for {
		if len(q) == 0 {
			break
		}

		current := q[0]
		q = q[1:]

		if current.value == to {
			result = true
			break
		}

		visited[current.value] = true

		for _, c := range current.children {
			if !visited[c.value] {
				q = append(q, c)
			}
		}
	}

	return result
}

// https://www.hackerrank.com/challenges/torque-and-development/problem
func roadsAndLibraries(cities, costLib, costRoad int32, roads [][]int32) int64 {
	if costLib <= costRoad {
		return int64(cities) * int64(costLib)
	}

	g := NewGraph(false)
	for _, road := range roads {
		g.Connect(road[0], road[1])
	}

	regions := 0
	visited := make([]bool, cities)

	for i := 1; i <= int(cities); i++ {
		if visited[i-1] {
			continue
		}

		q := []*Vertex{g.Vertex(int32(i), true)}

		for {
			if len(q) == 0 {
				break
			}

			current := q[0]
			q = q[1:]
			visited[current.value-1] = true

			for _, c := range current.children {
				if !visited[c.value-1] {
					q = append(q, c)
				}
			}
		}

		regions++
	}

	return int64(regions)*int64(costLib) + int64(cities-int32(regions))*int64(costRoad)
}

// https://www.hackerrank.com/challenges/find-the-nearest-clone/problem
func findShortest(vertices int32, from, to []int32, colors []int64, color int32) int32 {
	g := NewGraph(false)

	for i := 1; i <= int(vertices); i++ {
		v := g.Vertex(int32(i), true)
		v.color = int32(colors[i-1])
	}

	for i := 0; i < len(from); i++ {
		g.Connect(from[i], to[i])
	}

	var shortestPath int32 = -1

	visited := make([]bool, vertices)

	for i := 1; i <= int(vertices); i++ {
		v := g.Vertex(int32(i), false)
		if v.color == color {
			currentShortestPath := findShortestPathForVertex(v, color, visited, 0)
			if currentShortestPath == -1 {
				continue
			}
			if shortestPath == -1 || currentShortestPath < shortestPath {
				shortestPath = currentShortestPath
			}
		}
	}

	return shortestPath
}

func findShortestPathForVertex(vertex *Vertex, color int32, visited []bool, result int32) int32 {
	if vertex.color == color && result > 0 {
		return result
	}

	visited[vertex.value-1] = true

	var min int32 = -1

	for _, c := range vertex.children {
		if visited[c.value-1] {
			continue
		}

		current := findShortestPathForVertex(c, color, visited, result+1)
		if current == -1 {
			continue
		}
		if min == -1 || current < min {
			min = current
		}
	}

	return min
}

// https://www.hackerrank.com/challenges/ctci-bfs-shortest-reach/problem
func shortestReach(vertices int32, edges [][]int32, start int32) []int32 {
	g := NewGraph(false)

	for i := 1; i <= int(vertices); i++ {
		g.Vertex(int32(i), true)
	}

	for _, e := range edges {
		g.Connect(e[0], e[1])
	}

	distances := make([]int32, vertices)
	for i := 0; i < int(vertices); i++ {
		distances[i] = -1
	}

	q := []*Vertex{g.Vertex(start, false)}

	visited := make([]bool, vertices)

	distance := 0
	currentLevelVertices := 1
	nextLevelVertices := 0

	for {
		if len(q) == 0 {
			break
		}

		current := q[0]
		q = q[1:]
		currentLevelVertices--

		if !visited[current.value-1] {
			visited[current.value-1] = true
			distances[current.value-1] = int32(distance) * 6

			for _, c := range current.children {
				q = append(q, c)
				nextLevelVertices++
			}
		}

		if currentLevelVertices == 0 {
			distance++
			currentLevelVertices = nextLevelVertices
			nextLevelVertices = 0
		}
	}

	return append(distances[:start-1], distances[start:]...)
}

// https://www.hackerrank.com/challenges/ctci-connected-cell-in-a-grid/problem
func maxRegion(grid [][]int32) int32 {
	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}

	var max int32 = 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			localMax := countRegionCells(grid, visited, i, j)
			if localMax > max {
				max = localMax
			}
		}
	}

	return max
}

func countRegionCells(grid [][]int32, visited [][]bool, row, col int) int32 {
	if grid[row][col] == 0 {
		return 0
	}

	var res int32 = 1

	visited[row][col] = true

	if isValidCell(visited, row-1, col-1) {
		res += countRegionCells(grid, visited, row-1, col-1)
	}
	if isValidCell(visited, row-1, col) {
		res += countRegionCells(grid, visited, row-1, col)
	}
	if isValidCell(visited, row-1, col+1) {
		res += countRegionCells(grid, visited, row-1, col+1)
	}
	if isValidCell(visited, row, col+1) {
		res += countRegionCells(grid, visited, row, col+1)
	}
	if isValidCell(visited, row+1, col+1) {
		res += countRegionCells(grid, visited, row+1, col+1)
	}
	if isValidCell(visited, row+1, col) {
		res += countRegionCells(grid, visited, row+1, col)
	}
	if isValidCell(visited, row+1, col-1) {
		res += countRegionCells(grid, visited, row+1, col-1)
	}
	if isValidCell(visited, row, col-1) {
		res += countRegionCells(grid, visited, row, col-1)
	}

	return res
}

func isValidCell(visited [][]bool, row, col int) bool {
	if row < 0 || row >= len(visited) {
		return false
	}

	if col < 0 || col >= len(visited[0]) {
		return false
	}

	if visited[row][col] {
		return false
	}

	return true
}

// https://www.hackerrank.com/challenges/matrix/problem
type Road struct {
	a, b, destructionTime int32
}

type Zion []*Road

func (r Zion) Len() int {
	return len(r)
}

func (r Zion) Less(i, j int) bool {
	return r[i].destructionTime > r[j].destructionTime
}

func (r Zion) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func minTime(roads [][]int32, machines []int32) int32 {
	zion := make([]*Road, len(roads))

	for i, r := range roads {
		zion[i] = &Road{r[0], r[1], r[2]}
	}

	sort.Sort(Zion(zion))

	destroyedCities := make(map[int32]bool)
	for _, m := range machines {
		destroyedCities[m] = true
	}

	g := make([]int32, len(zion)+1)
	for i := 0; i < len(g); i++ {
		g[i] = int32(i)
	}

	destroyedRoots := make(map[int32]int32)
	for _, c := range g {
		if destroyedCities[c] {
			destroyedRoots[c] = c
		} else {
			destroyedRoots[c] = -1
		}
	}

	var result int32

	for _, r := range zion {
		result += buildOrDestroy(g, r, destroyedRoots)
	}

	return result
}

func buildOrDestroy(graph []int32, road *Road, destroyedRoots map[int32]int32) int32 {
	rootA := root(graph, road.a)
	rootB := root(graph, road.b)

	if destroyedRoots[rootA] != -1 && destroyedRoots[rootB] != -1 {
		return road.destructionTime
	}

	if destroyedRoots[rootA] == -1 && destroyedRoots[rootB] == -1 {
		graph[rootA] = rootB
		return 0
	}

	destroyedRoot := rootA
	freeRoot := rootB
	if destroyedRoots[freeRoot] != -1 {
		destroyedRoot, freeRoot = freeRoot, destroyedRoot
	}

	graph[freeRoot] = destroyedRoot
	return 0

}

func root(g []int32, n int32) int32 {
	if g[n] == n {
		return n
	}

	return root(g, g[n])
}
