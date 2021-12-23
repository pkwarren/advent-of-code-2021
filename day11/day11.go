package day11

import (
	"bufio"
	"io"
	"strings"
)

func ParseInput(r io.Reader) (*OctopusSimulation, error) {
	var sim OctopusSimulation
	s := bufio.NewScanner(r)
	lineno := 0
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) == 0 {
			continue
		}
		line := [10]int{}
		for i, c := range l {
			n := c - '0'
			line[i] = int(n)
		}
		sim.Grid[lineno] = line
		lineno += 1
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return &sim, nil
}

type OctopusSimulation struct {
	Grid    [10][10]int
	Flashes int
}

type Point struct {
	X, Y int
}

func (o *OctopusSimulation) Step() {
	var reset []Point
	for y, l := range o.Grid {
		for x := range l {
			o.Grid[y][x] += 1
			if o.Grid[y][x] > 9 {
				reset = append(reset, Point{X: x, Y: y})
			}
		}
	}
	for len(reset) > 0 {
		next := make(map[Point]struct{})
		for _, p := range reset {
			for x := max(p.X-1, 0); x < min(p.X+2, len(o.Grid[p.Y])); x++ {
				for y := max(p.Y-1, 0); y < min(p.Y+2, len(o.Grid)); y++ {
					if x == p.X && y == p.Y {
						o.Grid[y][x] = 0
						o.Flashes += 1
					} else if o.Grid[y][x] > 0 {
						o.Grid[y][x] += 1
						if o.Grid[y][x] == 10 {
							next[Point{X: x, Y: y}] = struct{}{}
						}
					}
				}
			}
		}
		reset = make([]Point, 0, len(next))
		for k := range next {
			reset = append(reset, k)
		}
	}
}

func min(n1, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
