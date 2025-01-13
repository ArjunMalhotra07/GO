package graphs

func OrangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	visited := make([][]int, len(grid))
	fresh := 0
	q := &Q{}

	for i := 0; i < len(grid); i++ {
		visited[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				q.enQ([]int{i, j, 0})
				visited[i][j] = 2
			} else if grid[i][j] == 1 {
				visited[i][j] = 0
				fresh += 1
			} else {
				visited[i][j] = 0
			}
		}
	}
	time := 0
	rows := []int{0, 0, 1, -1}
	cols := []int{1, -1, 0, 0}
	for len(q.vertices) != 0 {
		vertex := q.deQ()
		row := vertex[0]
		col := vertex[1]
		time = vertex[2]
		for i := 0; i < 4; i++ {
			newR := row + rows[i]
			newC := col + cols[i]
			if newR >= 0 && newR < len(grid) && newC >= 0 && newC < len(grid[0]) && visited[newR][newC] == 0 && grid[newR][newC] == 1 {
				q.enQ([]int{newR, newC, time + 1})
				visited[newR][newC] = 2
				fresh -= 1
			}

		}
	}
	if fresh != 0 {
		return -1
	}
	return time
}

type Q struct {
	vertices [][]int
}

func (q *Q) enQ(n []int) {
	q.vertices = append(q.vertices, n)
}

func (q *Q) deQ() []int {
	if len(q.vertices) != 0 {
		popped := q.vertices[0]
		if len(q.vertices) == 1 {
			q.vertices = [][]int{}
		} else {
			q.vertices = q.vertices[1:]
		}
		return popped
	}
	return []int{}
}
