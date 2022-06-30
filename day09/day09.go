package day09

import (
	"bufio"
	"io"
	"strings"
)

func ParseInput(r io.Reader, f func(line []int) error) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) == 0 {
			continue
		}
		nums := make([]int, 0)
		for _, c := range l {
			nums = append(nums, int(c-'0'))
		}
		if err := f(nums); err != nil {
			return err
		}
	}
	if err := s.Err(); err != nil {
		return err
	}
	return nil
}

type Point struct {
	X, Y int
}

type LowPoint struct {
	P         Point
	Value     int
	RiskLevel int
}

func ParseLines(r io.Reader) ([][]int, error) {
	lines := make([][]int, 0)
	err := ParseInput(r, func(line []int) error {
		lines = append(lines, line)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return lines, nil
}

func FindLowPoints(lines [][]int) []LowPoint {
	lowPoints := make([]LowPoint, 0)
	for i, l := range lines {
		for j, v := range l {
			up, down, left, right := 10, 10, 10, 10
			if i > 0 {
				up = lines[i-1][j]
			}
			if i < len(lines)-1 {
				down = lines[i+1][j]
			}
			if j > 0 {
				left = lines[i][j-1]
			}
			if j < len(l)-1 {
				right = lines[i][j+1]
			}
			if v < up && v < down && v < left && v < right {
				lowPoints = append(lowPoints, LowPoint{P: Point{X: j, Y: i}, Value: v, RiskLevel: v + 1})
			}
		}
	}
	return lowPoints
}

func SumOfRiskLevels(r io.Reader) (int, error) {
	lines, err := ParseLines(r)
	if err != nil {
		return 0, err
	}
	lowPoints := FindLowPoints(lines)
	sum := 0
	for _, p := range lowPoints {
		sum += p.RiskLevel
	}
	return sum, nil
}

func FindBasinSizes(r io.Reader) ([]int, error) {
	lines, err := ParseLines(r)
	if err != nil {
		return nil, err
	}
	lowPoints := FindLowPoints(lines)
	basinSizes := make([]int, 0, len(lowPoints))
	for _, p := range lowPoints {
		marked := make(map[Point]struct{})
		explore(p.P, lines, marked)
		basinSizes = append(basinSizes, len(marked))
	}
	return basinSizes, nil
}

func explore(p Point, lines [][]int, marked map[Point]struct{}) {
	if _, ok := marked[p]; ok {
		// We've already explored this location
		return
	}
	x, y := p.X, p.Y
	val := lines[y][x]
	if val < 9 {
		marked[p] = struct{}{}
		if x > 0 {
			explore(Point{X: x - 1, Y: y}, lines, marked)
		}
		if x < len(lines[y])-1 {
			explore(Point{X: x + 1, Y: y}, lines, marked)
		}
		if y > 0 {
			explore(Point{X: x, Y: y - 1}, lines, marked)
		}
		if y < len(lines)-1 {
			explore(Point{X: x, Y: y + 1}, lines, marked)
		}
	}
}
