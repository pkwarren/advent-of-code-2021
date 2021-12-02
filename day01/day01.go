package day01

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

func processFile(r io.Reader, f func(string) error) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if len(line) == 0 {
			continue
		}
		if err := f(line); err != nil {
			return err
		}
	}
	return nil
}

func processFileOfInts(r io.Reader, f func(int64) error) error {
	return processFile(r, func(s string) error {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		return f(i)
	})
}

func NumLargerMeasurements(r io.Reader) (int, error) {
	numLarger := 0
	var prevDepth int64
	if err := processFileOfInts(r, func(depth int64) error {
		if prevDepth > 0 && depth > prevDepth {
			numLarger++
		}
		prevDepth = depth
		return nil
	}); err != nil {
		return 0, err
	}
	return numLarger, nil
}

func NumLargerMeasurementsSlidingWindow(r io.Reader, window int) (int, error) {
	if window <= 0 {
		return 0, errors.New("invalid window")
	}
	numLarger := 0
	nextOffset := 0
	prevSum := int64(0)
	slidingWindow := make([]int64, 0, window)
	if err := processFileOfInts(r, func(depth int64) error {
		if len(slidingWindow) < window {
			slidingWindow = append(slidingWindow, depth)
			prevSum += depth
		} else {
			newSum := prevSum
			newSum = newSum - slidingWindow[nextOffset]
			newSum = newSum + depth
			slidingWindow[nextOffset] = depth
			if newSum > prevSum {
				numLarger = numLarger + 1
			}
			prevSum = newSum
		}
		nextOffset += 1
		if nextOffset >= window {
			nextOffset = 0
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return numLarger, nil
}
