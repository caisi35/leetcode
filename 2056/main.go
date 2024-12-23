package main

import (
	"fmt"
)

func main() {
	fmt.Println(countCombinations([]string{"bishop"}, [][]int{{4,3}}))
}

var rookDirections = [][2]int{{1,0},{-1,0},{0,1},{0,-1}}
var bishopDirections = [][2]int{{1,1},{1,-1},{-1,1},{-1,-1}}
var queenDirections = [][2]int{{1,0},{-1,0},{0,1},{0,-1},{1,1},{1,-1},{-1,1},{-1,-1}}

func countCombinations(prices []string, positions [][]int) int {
	n, res := len(prices), 0
	stack := []Movement{}
	check := func(u int) bool {
		for v := 0; v < u; v++ {
			if stack[u].cross(&stack[v]) {
				return false
			}
		}
		return true
	}
	var dfs func(u int)
	dfs = func(u int) {
		if u == n {
			res++
			return
		}
		var directions [][2]int
		switch prices[u] {
		case "rook":
			directions = rookDirections
		case "queen":
			directions = queenDirections
		default:
			directions = bishopDirections
		}
		stack = append(stack, Movement{startX: positions[u][0], startY: positions[u][1], endX: positions[u][0], endY: positions[u][1]})
		if check(u) {
			dfs(u+1)
		}
		stack = stack[:len(stack)-1]
		for _, dir := range directions {
			for j := 1; j < 8; j++ {
				x := positions[u][0] + dir[0] * j
				y := positions[u][1] + dir[1] * j
				if x < 1 || x > 8 || y < 1 || y > 8 {
					break
				}
				stack = append(stack, Movement{startX: positions[u][0], startY: positions[u][1], endX: x, endY: y, dx: dir[0], dy: dir[1]})
				if check(u) {
					dfs(u+1)
				}
				stack = stack[:len(stack)-1]
			}
		}
	}
	dfs(0)
	return res
}

type Movement struct {
	startX, startY, endX, endY, dx, dy, curX, curY int
}

func (m *Movement) reset() {
	m.curX = m.startX
	m.curY = m.startY
}

func (m *Movement) stopped() bool {
	return m.curX == m.endX && m.curY == m.endY
}

func (m *Movement) advance() {
	if !m.stopped() {
		m.curX += m.dx
		m.curY += m.dy
	}
}

func (m *Movement) cross(oth *Movement) bool {
	m.reset()
	oth.reset()
	for !m.stopped() || !oth.stopped() {
		m.advance()
		oth.advance()
		if m.curX == oth.curX && m.curY == oth.curY {
			return true
		}
	}
	return false
}