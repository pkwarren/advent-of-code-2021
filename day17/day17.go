package day17

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

type Velocity struct {
	X, Y int
}

type Probe struct {
	Position Point
	Velocity Velocity
}

type Target struct {
	MinX, MaxX int
	MinY, MaxY int
}

func (t Target) Hit(p Point) bool {
	return (p.X >= t.MinX && p.X <= t.MaxX) && (p.Y >= t.MinY && p.Y <= t.MaxY)
}

func ParseInput(r io.Reader) (*Target, error) {
	pattern := regexp.MustCompile(`^target area: x=(\d+)..(\d+), y=(-?\d+)..(-?\d+)$`)
	s := bufio.NewScanner(r)
	var t Target
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) == 0 {
			continue
		}
		submatches := pattern.FindStringSubmatch(l)
		if submatches == nil || len(submatches) != 5 {
			return nil, fmt.Errorf("invalid input: %s", l)
		}
		vals := make([]int, 0, 4)
		for _, match := range submatches[1:] {
			v, err := strconv.Atoi(match)
			if err != nil {
				return nil, err
			}
			vals = append(vals, v)
		}
		minX, maxX, minY, maxY := vals[0], vals[1], vals[2], vals[3]
		if minX > maxX || minY > maxY {
			return nil, fmt.Errorf("invalid input: %s", l)
		}
		t = Target{MinX: minX, MaxX: maxX, MinY: minY, MaxY: maxY}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return &t, nil
}

func CalculateHighestYPosition(target *Target) int {
	highestY := 0
	for x := 0; x <= target.MaxX; x++ {
		for y := target.MinY; y < abs(target.MinY); y++ {
			p := Probe{Velocity: Velocity{X: x, Y: y}}
			maxY := 0
			for p.Position.X <= target.MaxX && p.Position.Y >= target.MinY {
				if p.Position.Y > maxY {
					maxY = p.Position.Y
				}
				if target.Hit(p.Position) {
					if maxY > highestY {
						highestY = maxY
					}
					break
				}
				if p.Velocity.X == 0 && p.Position.X < target.MinX {
					break
				}
				p.Step()
			}
		}
	}
	return highestY
}

func CalculateDistinctInitialVelocities(target *Target) int {
	count := 0
	// Initial x velocity can't be greater than max target X
	for x := 0; x <= target.MaxX; x++ {
		// Initial y velocity can't be greater than min target Y.
		// If y velocity is positive, when it gets back to zero it's Y velocity will be at -y - 1.
		// This will overshoot the target if initial velocity y == target.MinY.
		for y := target.MinY; y < abs(target.MinY); y++ {
			p := Probe{Velocity: Velocity{X: x, Y: y}}
			for p.Position.X <= target.MaxX && p.Position.Y >= target.MinY {
				if target.Hit(p.Position) {
					count++
					break
				}
				// We won't reach MinX
				if p.Velocity.X == 0 && p.Position.X < target.MinX {
					break
				}
				p.Step()
			}
		}
	}
	return count
}

func (p *Probe) Step() {
	p.Position.X += p.Velocity.X
	p.Position.Y += p.Velocity.Y
	if p.Velocity.X > 0 {
		p.Velocity.X--
	} else if p.Velocity.X < 0 {
		p.Velocity.X++
	}
	p.Velocity.Y--
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
