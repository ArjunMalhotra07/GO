//! https://youtu.be/U5Mw4eyUmw4?si=awku13MID6_ehM_v
package takeuforward

import (
	"math"
	"reflect"
	"testing"
)

type _q struct {
	nums [][]int
}

func (q *_q) enQ(arr []int) {
	q.nums = append(q.nums, arr)
}
func (q *_q) deQ() []int {
	if len(q.nums) != 0 {
		front := q.nums[0]
		if len(q.nums) == 1 {
			q.nums = [][]int{}
		} else {
			q.nums = q.nums[1:]
		}
		return front
	}
	return []int{}
}

func maze(maze [][]int, src, dest []int) int {
	dist := make([][]int, len(maze))
	//! Make dist array and populate it with infinity
	for i := 0; i < len(dist); i++ {
		dist[i] = make([]int, len(maze[i]))
	}
	for i := 0; i < len(dist); i++ {
		for j := 0; j < len(dist[i]); j++ {
			dist[i][j] = math.MaxInt
		}
	}
	dist[src[0]][src[1]] = 0
	q := &_q{}
	q.enQ([]int{0, src[0], src[1]})
	rows := []int{0, 0, 1, -1}
	cols := []int{1, -1, 0, 0}

	for len(q.nums) != 0 {
		front := q.deQ()
		for i := 0; i < 4; i++ {
			newR := rows[i] + front[1]
			newC := cols[i] + front[2]
			if newR >= 0 && newR < len(dist) && newC >= 0 && newC < len(dist[newR]) && maze[newR][newC] == 1 && 1+front[0] < dist[newR][newC] {
				dist[newR][newC] = 1 + front[0]
				q.enQ([]int{dist[newR][newC], newR, newC})
			}
		}
	}
	return dist[dest[0]][dest[1]]
}

func TestMazeBinaryPath(t *testing.T) {
	tests := []struct {
		arr      [][]int
		expected int
		src      []int
		dest     []int
	}{
		{arr: [][]int{{1, 1, 1, 1}, {1, 1, 0, 1}, {1, 1, 1, 1}, {1, 1, 0, 0}, {1, 0, 0, 0}}, expected: 3, src: []int{0, 1}, dest: []int{2, 2}},
		{arr: [][]int{{1, 1, 1, 1}, {1, 1, 0, 1}, {1, 1, 1, 1}, {1, 1, 0, 0}, {1, 0, 0, 0}}, expected: 5, src: []int{0, 1}, dest: []int{4, 0}},
	}
	for _, testCase := range tests {
		t.Run("", func(t *testing.T) {
			actualResult := maze(testCase.arr, testCase.src, testCase.dest)
			if !reflect.DeepEqual(actualResult, testCase.expected) {
				t.Errorf("Expected %v but got %v", testCase.expected, actualResult)
			}
		})
	}
}
