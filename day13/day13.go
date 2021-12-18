package day13

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ParseInput(r io.Reader) (*Manual, error) {
	s := bufio.NewScanner(r)
	var m Manual
	m.Dots = make(map[Point]struct{})
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) == 0 {
			continue
		}
		if strings.HasPrefix(l, "fold along ") {
			l := l[11:]
			fields := strings.Split(l, "=")
			if len(fields) != 2 {
				return nil, fmt.Errorf("expected 2 fields, found: %v", len(fields))
			}
			v, err := strconv.Atoi(fields[1])
			if err != nil {
				return nil, err
			}
			var fold Fold
			switch fields[0] {
			case "x":
				fold.X = v
			case "y":
				fold.Y = v
			default:
				return nil, fmt.Errorf("expected x/y, found: %s", fields[0])
			}
			m.Folds = append(m.Folds, fold)
		} else {
			fields := strings.Split(l, ",")
			if len(fields) != 2 {
				return nil, fmt.Errorf("expected 2 fields, found: %v", len(fields))
			}
			x, err := strconv.Atoi(fields[0])
			if err != nil || x < 0 {
				return nil, fmt.Errorf("invalid X coordinate: %s", fields[0])
			}
			y, err := strconv.Atoi(fields[1])
			if err != nil || y < 0 {
				return nil, fmt.Errorf("invalid Y coordinate: %s", fields[1])
			}
			if x > m.MaxX {
				m.MaxX = x
			}
			if y > m.MaxY {
				m.MaxY = y
			}
			m.Dots[Point{X: x, Y: y}] = struct{}{}
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return &m, nil
}

type Point struct {
	X, Y int
}

type Manual struct {
	Dots        map[Point]struct{}
	Folds       []Fold
	CurrentFold int
	MaxX, MaxY  int
}

func (m *Manual) FoldNext() {
	if m.CurrentFold >= len(m.Folds) {
		// We've performed all the folds
		return
	}
	fold := m.Folds[m.CurrentFold]
	if fold.X > 0 && fold.X <= m.MaxX {
		// fold along X axis
		toDelete := make([]Point, 0)
		for p := range m.Dots {
			if p.X >= fold.X {
				toDelete = append(toDelete, p)
			}
		}
		for _, p := range toDelete {
			delete(m.Dots, p)
			if p.X > fold.X && p.X <= fold.X*2 {
				m.Dots[Point{X: fold.X*2 - p.X, Y: p.Y}] = struct{}{}
			}
		}
		m.MaxX = fold.X - 1
	} else if fold.Y > 0 && fold.Y <= m.MaxY {
		// fold along Y axis
		toDelete := make([]Point, 0)
		for p := range m.Dots {
			if p.Y >= fold.Y {
				toDelete = append(toDelete, p)
			}
		}
		for _, p := range toDelete {
			delete(m.Dots, p)
			if p.Y > fold.Y && p.Y <= fold.Y*2 {
				m.Dots[Point{X: p.X, Y: fold.Y*2 - p.Y}] = struct{}{}
			}
		}
		m.MaxY = fold.Y - 1
	}
	m.CurrentFold += 1
}

func (m *Manual) NumDots() int {
	return len(m.Dots)
}

func (m *Manual) String() string {
	sb := strings.Builder{}
	for y := 0; y <= m.MaxY; y++ {
		for x := 0; x <= m.MaxX; x++ {
			if _, ok := m.Dots[Point{X: x, Y: y}]; ok {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

type Fold struct {
	X, Y int
}
