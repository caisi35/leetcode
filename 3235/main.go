package main

import (
	"fmt"
)

func main() {
	fmt.Println(canReachCorner(3, 3, [][]int{{1,1,2}}))
}

func canReachCorner(xCorner int, yCorner int, circles [][]int) bool {
	visited := make([]bool, len(circles))
	var dfs func(i int) bool
	dfs = func(i int) bool {
		x1, y1, r1 := circles[i][0], circles[i][1], circles[i][2]
		if circleIntersectsBottomRightOfRectangle(x1, y1, r1, xCorner, yCorner) {
			return true
		}
		visited[i] = true
		for j := 0; j < len(circles); j++ {
			if !visited[j] && circlesIntersectInRectangle(x1, y1, r1, circles[j][0], circles[j][1], circles[j][2], xCorner, yCorner) && dfs(j) {
				return true
			}
		}
		return false
	}
	for i := range circles {
		x, y, r := circles[i][0], circles[i][1], circles[i][2]
		if pointInCircle(0, 0, x, y, r) || pointInCircle(xCorner, yCorner, x, y ,r) {
			return false
		}
		if !visited[i] && circleIntersectsTopLeftOfRectangle(x, y, r, xCorner, yCorner) && dfs(i) {
			return false
		}
	}
	return true
}

func pointInCircle(px, py, x, y, r int) bool {
	return (x - px) * (x - px) + (y - py) * (y - py) <= r * r
}

func circleIntersectsTopLeftOfRectangle(x, y, r, xCorner, yCorner int) bool {
	return (abs(x) <= r && 0 <= y && y <= yCorner) ||
	(0 <= x && x <= xCorner && abs(y - yCorner) <= r) ||
	pointInCircle(x, y, 0, yCorner, r)
}

func circleIntersectsBottomRightOfRectangle(x, y, r, xCorner, yCorner int) bool {
	return (abs(y) <= r && 0 <= x && x <= xCorner) ||
	(0 <= y && y <= yCorner && abs(x - xCorner) <= r) ||
	pointInCircle(x, y, xCorner, 0, r)
}

func circlesIntersectInRectangle(x1, y1, r1, x2, y2, r2, xCorner, yCorner int) bool {
	return (x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2) <= (r1 + r2) * (r1 + r2) && 
	x1 * r2 + x2 * r1 < (r1 + r2) * xCorner &&
	y1 * r2 + y2 * r1 < (r1 + r2) * yCorner
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}