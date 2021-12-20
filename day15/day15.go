package day15

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"strings"
)

func ParseInput(r io.Reader) ([][]int, error) {
	m := make([][]int, 0)
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) == 0 {
			continue
		}
		row := make([]int, 0, len(l))
		for _, r := range l {
			v := int(r - '0')
			if v >= 0 && v <= 9 {
				row = append(row, v)
			} else {
				return nil, fmt.Errorf("invalid row: %s", l)
			}
		}
		if len(m) > 0 && len(m[0]) != len(row) {
			return nil, fmt.Errorf("expected row with %d values, found: %d", len(m[0]), len(row))
		}
		m = append(m, row)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return m, nil
}

type Point struct {
	X, Y int
}

type PointValue struct {
	X, Y int
	Val  int
}

type PointValueHeap []PointValue

func (p PointValueHeap) Len() int {
	return len(p)
}

func (p PointValueHeap) Less(i, j int) bool {
	return p[i].Val < p[j].Val
}

func (p PointValueHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PointValueHeap) Push(i interface{}) {
	*p = append(*p, i.(PointValue))
}

func (p *PointValueHeap) Pop() interface{} {
	prev := *p
	n := len(prev)
	v := prev[n-1]
	*p = prev[0 : n-1]
	return v
}

type Path []Point

func FindLowestTotalRisk(m [][]int) int {
	h := &PointValueHeap{}
	heap.Push(h, PointValue{X: 0, Y: 0, Val: m[0][0]})
	parent := make(map[Point]Point)
	cost := make(map[Point]int)
	start := Point{X: 0, Y: 0}
	end := Point{X: len(m[0]) - 1, Y: len(m) - 1}
	cost[start] = 0

	for h.Len() > 0 {
		current := heap.Pop(h).(PointValue)
		neighbors := make([]Point, 0, 4)
		// Try neighbor to the left
		if current.X > 0 {
			neighbors = append(neighbors, Point{X: current.X - 1, Y: current.Y})
		}
		// Try neighbor to the right
		if current.X < len(m[0])-1 {
			neighbors = append(neighbors, Point{X: current.X + 1, Y: current.Y})
		}
		// Try neighbor to the top
		if current.Y > 0 {
			neighbors = append(neighbors, Point{X: current.X, Y: current.Y - 1})
		}
		// Try neighbor to the bottom
		if current.Y < len(m)-1 {
			neighbors = append(neighbors, Point{X: current.X, Y: current.Y + 1})
		}
		for _, neighbor := range neighbors {
			newCost := cost[Point{X: current.X, Y: current.Y}] + m[neighbor.Y][neighbor.X]
			if currentCost, ok := cost[neighbor]; ok {
				if currentCost > newCost {
					cost[neighbor] = newCost
					parent[neighbor] = Point{X: current.X, Y: current.Y}
					heap.Push(h, PointValue{X: neighbor.X, Y: neighbor.Y, Val: newCost})
				}
			} else {
				// We don't have a current cost - assume this is shortest
				cost[neighbor] = newCost
				parent[neighbor] = Point{X: current.X, Y: current.Y}
				heap.Push(h, PointValue{X: neighbor.X, Y: neighbor.Y, Val: newCost})
			}
		}
	}
	return cost[end]
}

func FullMap(tile [][]int) [][]int {
	xLen := len(tile[0])
	yLen := len(tile)
	m := make([][]int, 0, len(tile)*5)
	for y := 0; y < yLen*5; y++ {
		m = append(m, make([]int, 0, xLen*5))
		for x := 0; x < xLen*5; x++ {
			if x < xLen && y < yLen {
				m[y] = append(m[y], tile[y][x])
			} else if x < xLen {
				val := m[y-yLen][x] + 1
				if val > 9 {
					val = 1
				}
				m[y] = append(m[y], val)
			} else {
				val := m[y][x-xLen] + 1
				if val > 9 {
					val = 1
				}
				m[y] = append(m[y], val)
			}
		}
	}
	return m
}
