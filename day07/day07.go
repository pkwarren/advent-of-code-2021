package day07

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"
)

func ParseInput(r io.Reader) ([]int, error) {
	s := bufio.NewScanner(r)
	nums := make([]int, 0)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) == 0 {
			continue
		}
		if len(nums) > 0 {
			return nil, fmt.Errorf("invalid input")
		}
		fields := strings.Split(l, ",")
		for _, f := range fields {
			i, err := strconv.Atoi(f)
			if err != nil {
				return nil, err
			}
			nums = append(nums, i)
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return nums, nil
}

func CalculateMinFuelPosition(positions []int) (int, int) {
	if len(positions) <= 1 {
		return 0, 0
	}
	sort.Ints(positions)
	minFuel := -1
	position := -1
	for i := positions[0]; i <= positions[len(positions)-1]; i++ {
		fuel := 0
		for _, pos := range positions {
			fuel += int(math.Abs(float64(pos - i)))
			if minFuel >= 0 && fuel > minFuel {
				// We're already over previous minimum
				break
			}
		}
		if minFuel < 0 || fuel < minFuel {
			minFuel = fuel
			position = i
		}
	}
	return position, minFuel
}

func CalculateMinFuelPositionPart2(positions []int) (int, int) {
	if len(positions) <= 1 {
		return 0, 0
	}
	sort.Ints(positions)
	minFuel := -1
	position := -1
	for i := positions[0]; i <= positions[len(positions)-1]; i++ {
		fuel := 0
		for _, pos := range positions {
			distance := int(math.Abs(float64(pos - i)))
			fuel += (distance * (distance + 1)) / 2
			if minFuel >= 0 && fuel > minFuel {
				// We're already over previous minimum
				break
			}
		}
		if minFuel < 0 || fuel < minFuel {
			minFuel = fuel
			position = i
		}
	}
	return position, minFuel
}
