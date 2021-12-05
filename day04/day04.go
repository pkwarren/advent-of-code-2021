package day04

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type BingoGame struct {
	Numbers []int
	Boards  []*Board
}

func ParseGame(r io.Reader) (*BingoGame, error) {
	g := BingoGame{}
	s := bufio.NewScanner(r)
	expectedSize := 0
	gridSize := 0
	var currentBoard []int
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) == 0 {
			continue
		}
		if len(g.Numbers) == 0 {
			fields := strings.Split(l, ",")
			for _, f := range fields {
				i, err := strconv.ParseInt(f, 10, 0)
				if err != nil {
					return nil, err
				}
				g.Numbers = append(g.Numbers, int(i))
			}
		} else {
			fields := strings.Fields(l)
			if len(fields) <= 1 {
				return nil, fmt.Errorf("invalid board size")
			}
			if expectedSize == 0 {
				expectedSize = len(fields)
				gridSize = expectedSize * expectedSize
			} else if len(fields) != expectedSize {
				return nil, fmt.Errorf("inconsistent board size, expected: %d", expectedSize)
			}
			for _, f := range fields {
				i, err := strconv.ParseInt(f, 10, 0)
				if err != nil {
					return nil, err
				}
				currentBoard = append(currentBoard, int(i))
			}
			if len(currentBoard) == gridSize {
				b, err := NewBoard(currentBoard, expectedSize)
				if err != nil {
					return nil, err
				}
				g.Boards = append(g.Boards, b)
				currentBoard = make([]int, 0, len(currentBoard))
			}
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	if len(g.Numbers) == 0 {
		return nil, fmt.Errorf("failed to read bingo numbers")
	}
	if len(g.Boards) == 0 {
		return nil, fmt.Errorf("failed to read bingo boards")
	}
	return &g, nil
}

// Board layout:
//
// given size of 5, slice of 25 elements:
// [N0,  ..., N4,  // row 1
//  N5,  ..., N9,  // row 2
//  N10, ..., N14, // row 3
//  N15, ..., N19, // row 4
//  N20, ..., N24] // row 5
//
// map[int]int - for fast lookup of a number on board
// unmarkedTotal int - for fast calculation of remaining unmarked items
type Board struct {
	values           []int
	marked           map[int]struct{}
	size             int
	numberToPosition map[int]int
	unmarkedTotal    int
	won              bool
}

func NewBoard(values []int, size int) (*Board, error) {
	if size <= 0 {
		return nil, fmt.Errorf("invalid board size: %d", size)
	}
	n := size * size
	if len(values) != n {
		return nil, fmt.Errorf("expected %d numbers, passed %d", n, len(values))
	}
	unmarkedTotal := 0
	posIndex := make(map[int]int, n)
	for i, v := range values {
		if _, ok := posIndex[v]; ok {
			return nil, fmt.Errorf("duplicate value %d found on board", v)
		}
		posIndex[v] = i
		unmarkedTotal += v
	}
	return &Board{
		values:           values,
		marked:           make(map[int]struct{}, n),
		size:             size,
		numberToPosition: posIndex,
		unmarkedTotal:    unmarkedTotal,
	}, nil
}

func (b *Board) Mark(val int) bool {
	idx := 0
	var ok bool
	if idx, ok = b.numberToPosition[val]; !ok {
		return false
	}
	// Only mark a value once - otherwise unmarkedTotal will be inaccurate
	if _, ok := b.marked[idx]; ok {
		return true
	}
	b.marked[idx] = struct{}{}
	b.unmarkedTotal -= val
	// Check row if won
	rowStart := (idx / b.size) * b.size
	won := true
	for i := rowStart; i < rowStart+b.size; i++ {
		if _, ok := b.marked[i]; !ok {
			won = false
			break
		}
	}
	if !won {
		// Check column if won
		columnStart := idx % b.size
		won = true
		for i := columnStart; i < len(b.values); i += 5 {
			if _, ok := b.marked[i]; !ok {
				won = false
				break
			}
		}
	}
	if won {
		b.won = true
	}
	return true
}

func (b *Board) Won() bool {
	return b.won
}

func (b *Board) UnmarkedTotal() int {
	return b.unmarkedTotal
}
