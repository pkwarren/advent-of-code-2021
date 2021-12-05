package day05

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return "Point[" + strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + "]"
}

type Line struct {
	P0, P1 Point
}

func (l Line) Slope() (int, int) {
	return l.P1.Y - l.P0.Y, l.P1.X - l.P0.X
}

func ParseLines(r io.Reader) ([]Line, error) {
	s := bufio.NewScanner(r)
	lines := make([]Line, 0)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) == 0 {
			continue
		}
		pts := strings.Split(l, "->")
		if len(pts) != 2 {
			return nil, fmt.Errorf("invalid line: %s", l)
		}
		line := Line{}
		for i, p := range pts {
			coords := strings.Split(strings.TrimSpace(p), ",")
			if len(coords) != 2 {
				return nil, fmt.Errorf("invalid coordinate %s on line: %s", p, l)
			}
			x, err := strconv.ParseInt(coords[0], 10, 0)
			if err != nil {
				return nil, err
			}
			y, err := strconv.ParseInt(coords[1], 10, 0)
			if err != nil {
				return nil, err
			}
			if i == 0 {
				line.P0 = Point{int(x), int(y)}
			} else {
				line.P1 = Point{int(x), int(y)}
			}
		}
		lines = append(lines, line)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

type Map struct {
	Marker map[Point]int
}

func (m *Map) Mark(line Line, diagonals bool) {
	// Part 1 - only consider horizontal or vertical
	slopeY, slopeX := line.Slope()
	if slopeY != 0 && slopeX != 0 {
		if !diagonals {
			// part 1 - don't support diagonals
			return
		}
		if math.Abs(float64(slopeY)) != math.Abs(float64(slopeX)) {
			log.Printf("line %v not at a 45 degree angle", line)
			return
		}
	}
	if m.Marker == nil {
		m.Marker = make(map[Point]int)
	}
	stepY, stepX := calculateStep(slopeY), calculateStep(slopeX)
	for p := line.P0; ; p.Y, p.X = p.Y+stepY, p.X+stepX {
		m.Marker[p] = m.Marker[p] + 1
		if p == line.P1 {
			break
		}
	}
}

func (m *Map) OverlappingPoints() int {
	count := 0
	for _, v := range m.Marker {
		if v > 1 {
			count += 1
		}
	}
	return count
}

func calculateStep(n int) int {
	step := 0
	if n > 0 {
		step = 1
	} else if n < 0 {
		step = -1
	}
	return step
}
